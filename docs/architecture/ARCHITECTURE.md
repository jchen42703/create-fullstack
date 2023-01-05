# Architecture

## Requirements

### 1.1 Generate Template with CLI UI

`create_fullstack gen`

- Brings up CLI UI
- The CLI UI should build the YAML config and pass it to the `YAMLPipeline`

### 1.2 Generate Template with Yaml Config

`create_fullstack gen -f cfg.yml`

### 1.3 Augment current repository

`create_fullstack augment [ui/api/infra] <TYPE>`

`<TYPE>` should be one of the various augmentation types for the corresponding category.

## Pipeline

```
YamlPipeline
  Reads yaml config from filepath and converts to struct `GeneralConfig`
  Validates the `GeneralConfig`
    if err -> CLI fatals
  Calls GenerateFullstackTemplate(GeneralConfig)
    Calls FullstackGenerator
      Calls UiGenerator, ApiGenerator
```

## User Config

```yaml
# Each of the top level sections (ui, api, infra) are optional unless fullstack exists
# Generates everything
# Infra, ui, backend
fullstack:
  output_dir: "./fullstack_app_name"
  payments:
    stripe: true
    paypal: true
  # One of exists or doesn't exist
  # Uses ory
  # Future: auth0
  # Generating auth code took complex and potentially dangerous for this project
  auth:
    username_password:
      # email verification by default
      email_verification: true
      username_is_email: true
      password_validation: true
    social_sign_in:
      google_callback_url: "http://localhost:3000/dashboard"
      facebook_callback_url: "http://localhost:3000/dashboard"
      github_callback_url: "http://localhost:3000/dashboard"

ui:
  output_dir: "./fullstack_app_name/ui"
  # Base template to augment
  # one of:
  # nextjs, nextjs_ory, create_react_app
  # - can support others like angular, vue, svelte gatsby, vite, etc.
  # If not one of those:
  # Check if it's a git url: and git clone
  # If not: throw error
  base: nextjs
  # go, python, javascript, typescript
  lang: typescript
  augment:
    tailwind:
      # one of default, latest or the version number
      # [package-name]@[version-number]
      # Whenever, we run commands, we fix the version numbers as part of the default behavior.
      # But, if you're making a plugin, you should specify the version numebr that works for you.
      version: default
    scss:
      version: default
    styled_components:
      version: default
    # Should be null if lang != javascript/typescript
    # Prettier and eslint required for all js/ts
    husky:
      commitlint:
        version: default
      format: true
      lint: true

    # If custom base template, then behavior can vary.
    dockerfile: true
    # If custom base template, then behavior can vary.
    # I.e. circleci, travisci, jenkins, git_workflows
    ci: git_workflows

api:
  output_dir: "./fullstack_app_name/api"
  # Base template to augment
  # one of:
  # default, echo, echo_ory, express, fastify, fastapi
  # - can support others like gin, chi
  # If not one of those:
  # Check if it's a git url: and git clone
  # If not: throw error
  base: echo
  lang: go
  # Not too sure about this option
  db:
    # one of:
    # postgres, mysql, cockroachdb
    sql:
      db_type: postgres
      startup_script: ""
    no_sql:
      mongodb: true
      redis: true
      cassandra: true
      scylladb: true
  augment:
    # js-only
    # husky:
    #   commitlint: true
    # If not javascript/typescript, husky is null
    husky: null
    # anything besides js
    pre_commit:
      lint: true
      # null if you don't want a formatter
      # format: null
      format:
        # Go just has go_fmt
        # Python formatters:
        # yapf, black, pep8
        formatter: black

    # if custom base template, then behavior can vary.
    dockerfile: true
    git:
      issue_templates: true
      pr_templates: true
    # If custom base template, then behavior can vary.
    # I.e. circleci, travisci, jenkins, git_workflows
    ci: git_workflows

# Tentative as of now
infra:
  output_dir: "./fullstack_app_name/infra"
  docker_compose: true
  k8s: true
  nginx: true

  # mysql, postgres
  db_image:
  ui_image:
  api_image:
  domain_name:
  # gcloud/aws/azure
  cloud_provider: gcloud

# Low prio
cli:
api_client:
  # api client sdk
```

## Problems

Lots of permutations of potential templates.

Hence, everything should be done dynamically if possible.

### Code Generation

However, it's pretty difficult to generate functional code for custom base templates.

- The way OpenAPI generates code is very rigid.
- Hence, I think it's best to not generate code for DBs/payments for custom base templates
  - They need to make their own plugins
- Should generate infra tho

Generating the CI templates should be pretty easy because we'll just need to generate a very basic template that:

1. Builds
2. Runs tests

The rest can be left to the user (i.e. automating releases, uploading to s3/a Digital Ocean space, etc.)

## Dependency Management

Since we're doing everything dynamically, it can be very easy for templates to break if we do not keep dependency versions fixed.

For example, if we add tailwind dynamically, but if the newest version of tailwind breaks the template, then there is not much we can do about it except by periodocially running tests to confirm that the template still works.

**Main dependencies:**

- UI
  - JS/TS
    - Next.js
    - Create React App
    - Tailwind
    - SCSS
    - Husky
    - Stripe SDK
- API
  - Go
    - Stripe SDK
- Github Actions

We probably just need to fix the versions whenever we `yarn add` or install something.

Then, if people want to use a more modern version, we can simply create plugin for it.

Then, we can merge it to `main` if it is compatible with the base templates and then bump the version of the CLI.
