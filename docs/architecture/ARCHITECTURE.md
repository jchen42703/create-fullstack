# Architecture

## User Config

```yaml
# Each of the top level sections (ui, api, infra) are optional unless fullstack exists
# Generates everything
# Infra, ui, backend
fullstack:
  payments:
    stripe: true
    paypal: true
  # One of exists or doesn't exist
  # Uses ory
  # Future: auth0
  # Generating auth code took complex and potentially dangerous for this project
  auth:
    # email verification by default
    username_password: true
    social_sign_in:
      google: true
      facebook: true
      github: true

ui:
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
    tailwind: true
    scss: true
    styled_components: true
    # Should be null if lang != javascript/typescript
    # Prettier and eslint required for all js/ts
    husky:
      commitlint: true
    prettier_eslint: true
    # If custom base template, then behavior can vary.
    dockerfile: true
    # If custom base template, then behavior can vary.
    # I.e. circleci, travisci, jenkins, git_workflows
    ci: git_workflows

api:
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
    sql: postgres
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
