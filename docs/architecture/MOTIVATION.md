# CLI Design Methodology

The overarching goal is to make standardizing repository creation and standards consistent across organizations quick and easy!

Hence, the CLI needs to be easily generalizable and customizable for different organizations, while providing signficant utility.

**The key requirements are:**

1. Should be able to generate code with a configuration.
   1. Should be able to modify the template that's pulled with each code generation.
   2. Store default templates locally.
2. Should have separate code generators for:
   1. Frontend-Only
      1. Static?
      2. Non-Static
   2. Backend Only
   3. Fullstack

Here's a rundown of the specific options needed:

1. **Fullstack**

   - Open ID Connect?
   - Username/Passwords?
   - Stripe Integration
   - Paypal Integration

2. **Frontend**

   - Next.js with Dockerfile
   - Mid Prio Alternatives
     - Angular.js with Dockerfile
     - Vue with Dockerfile
     - Svelte with Dockerfile
   - Low Prio Alternatives
     - Gatsby
     - Vite
     - Jekyll
   - Options for All
     - JS/TS?
     - SCSS vs CSS?
     - Tailwind?
     - Styled Components?
     - Husky
       - Prettier?
       - Commitlint?
       - ESLint?
     - Use Jest for testing
   - **Enabled for All**
     - ESLint
     - Prettier

3. **Backend**

   - FastAPI with poetry + Docker
   - Express with Dockerfile
   - Go Echo API with Dockerfile
   - Options for all
     - Database?
       - SQL
         - CockroachDB
         - Postgres
         - MySQL
       - No-SQL
         - MongoDB
         - ScyllaDB
         - Redis
         - Cassandra
     - Precommit
       - Commit lint?
       - If JS
         - Prettier?
         - ESLint?
       - If Go
       - If Python
         - Yapf
         - Black
         - Pep8

4. **IAC**

   - Docker compose
   - Example Kubernetes Configs
   - TODO:
     - Example Terraform modules
     - Example Helm Modules
   - User Inputs
     - Gcloud vs AWS
     - Key
     - Gcloud specific info
     - Domain name
     - Frontend/backend image names
   - TODO:
     - Figure out how to deploy containers to separate node pools
     - Make configs for existing deployments

## Standards

1. Enforcing code quality
   1. Use precommit
      1. https://www.conventionalcommits.org/en/v1.0.0/
      2. https://bongnv.com/blog/2021-08-29-pre-commit-hooks-golang-projects/
   2. CI + tests
      1. Block PRs if tests fail
2. Continuous Delivery
   1. Whenever a push is made to main with a new version, start Git Release + prompt for release notes.
