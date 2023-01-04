package augment

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/jchen42703/create-fullstack/internal/directory"
)

// Adds Tailwind to the project following the official docs:
// https://tailwindcss.com/docs/guides/nextjs
// This is better than the regular with-tailwind Next.js template because
// this approach works with Typescript too.
func AddTailwind() error {
	log.Print("Adding Tailwind and peer dependencies...\n")
	commands := []*exec.Cmd{
		exec.Command("yarn", "add", "-D", "tailwindcss", "postcss", "autoprefixer"),
		exec.Command("npx", "tailwindcss", "init", "-p"),
	}

	for _, cmd := range commands {
		err := RunCommand(cmd, log.Writer()) //blocks until sub process is complete
		if err != nil {
			return fmt.Errorf("AddTailwind: %s", err.Error())
		}
	}

	tailwindConfig := `/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./pages/**/*.{js,ts,jsx,tsx}",
		"./components/**/*.{js,ts,jsx,tsx}",
	],
	theme: {
		extend: {},
	},
	plugins: [],
};
`

	err := os.WriteFile("./tailwind.config.js", []byte(tailwindConfig), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("AddTailwind: writing config failed: %s", err.Error())
	}

	// Assume globals.scss/css is in styles/
	possibleStylesPaths := []string{
		"./styles/globals.css",
		"./styles/globals.scss",
	}

	// Checks if any of the path exists and tries to match
	globalStylesPath := ""
	for _, stylesPath := range possibleStylesPaths {
		var exists bool
		if exists, err = directory.Exists(stylesPath); exists {
			globalStylesPath = stylesPath
			break
		}
	}

	// DNE
	if globalStylesPath == "" && err == nil {
		return fmt.Errorf("AddTailwind: template must have a globals css or scss file for path '%s'", globalStylesPath)
	} else if err != nil {
		return fmt.Errorf("AddTailwind: validating styles path failed: %s", err.Error())
	}

	// Exists, so add tailwind styles to global styles
	tailwindHeader := `@tailwind base;
@tailwind components;
@tailwind utilities;

`

	readBytes, err := os.ReadFile(globalStylesPath)
	if err != nil {
		return fmt.Errorf("AddTailwind: failed to read global styles: %s", err.Error())
	}

	newGlobalStyles := append([]byte(tailwindHeader), readBytes...)
	err = os.WriteFile(globalStylesPath, []byte(newGlobalStyles), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("AddTailwind: writing global styles failed: %s", err.Error())
	}

	return nil
}

func InitializeNextDocker(portNum int) error {
	newNextConfig := `/** @type {import('next').NextConfig} */
const nextConfig = {
	reactStrictMode: true,
	swcMinify: true,
	eslint: {
	ignoreDuringBuilds: true,
	},
	output: "standalone",
};

module.exports = nextConfig;
`
	err := os.WriteFile("./next.config.js", []byte(newNextConfig), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("InitializeNextDocker: failed to write next.config.js: %s", err.Error())
	}

	newDockerIgnore := `Dockerfile
.dockerignore
node_modules
npm-debug.log
README.md
.next
.git
`

	err = os.WriteFile(".dockerignore", []byte(newDockerIgnore), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("InitializeNextDocker: failed to write .dockerignore: %s", err.Error())
	}

	newDockerFile := fmt.Sprintf(`# Install dependencies only when needed
FROM node:16-alpine AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /app

# Install dependencies based on the preferred package manager
COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* ./
RUN \
	if [ -f yarn.lock ]; then yarn --frozen-lockfile; \
	elif [ -f package-lock.json ]; then npm ci; \
	elif [ -f pnpm-lock.yaml ]; then yarn global add pnpm && pnpm i --frozen-lockfile; \
	else echo "Lockfile not found." && exit 1; \
	fi


# Rebuild the source code only when needed
FROM node:16-alpine AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY . .

# Next.js collects completely anonymous telemetry data about general usage.
# Learn more here: https://nextjs.org/telemetry
# Uncomment the following line in case you want to disable telemetry during the build.
ENV NEXT_TELEMETRY_DISABLED 1

RUN yarn build

# If using npm comment out above and use below instead
# RUN npm run build

# Production image, copy all the files and run next
FROM node:16-alpine AS runner
WORKDIR /app

ENV NODE_ENV production

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nextjs

COPY --from=builder /app/public ./public

# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=builder --chown=nextjs:nodejs /app/.next/standalone ./
COPY --from=builder --chown=nextjs:nodejs /app/.next/static ./.next/static

USER nextjs

EXPOSE %d

ENV PORT %d

CMD ["node", "server.js"]
`, portNum, portNum)
	err = os.WriteFile("Dockerfile", []byte(newDockerFile), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("InitializeNextDocker: failed to write Dockerfile: %s", err.Error())
	}

	return nil
}
