package ui

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jchen42703/create-fullstack/internal/directory"
)

// Initializes a Next.js app be buildable by docker.
func InitializeNextDocker(workingDir string, portNum int, expose bool) error {
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
	nextCfgPath := filepath.Join(workingDir, "next.config.js")
	err := os.WriteFile(nextCfgPath, []byte(newNextConfig), directory.READ_WRITE_PERM)
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

	dockerIgnorePath := filepath.Join(workingDir, ".dockerignore")
	err = os.WriteFile(dockerIgnorePath, []byte(newDockerIgnore), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("InitializeNextDocker: failed to write .dockerignore: %s", err.Error())
	}

	newDockerFile := `# Install dependencies only when needed
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

`
	if expose {
		newDockerFile += fmt.Sprintf("EXPOSE %d\n", portNum)
	}

	newDockerFile += fmt.Sprintf(`
ENV PORT %d

CMD ["node", "server.js"]
`, portNum)

	dockerfilePath := filepath.Join(workingDir, "Dockerfile")
	err = os.WriteFile(dockerfilePath, []byte(newDockerFile), directory.READ_WRITE_PERM)
	if err != nil {
		return fmt.Errorf("InitializeNextDocker: failed to write Dockerfile: %s", err.Error())
	}

	return nil
}
