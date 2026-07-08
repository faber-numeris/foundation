---
name: skill-authn
description: >-
  Shared SDK for the Luciole Authentication Service API — spec-first
  codegen that produces Go stubs (oapi-codegen, chi-server) and a
  TypeScript fetch client (orval, Zod). Load this skill when editing
  openapi.yaml, regenerating stubs, publishing the npm package, or
  working with Bruno API test collections.
type: core
---

# authn — Luciole Authentication Service SDK

> **CRITICAL:** `openapi.yaml` is the single source of truth. Never edit generated files in `api/`, `ui/src/api/`, or `ui/src/models/` by hand. They are overwritten on every `just generate-openapi-stubs` run.

## Workflow: modify the API

1. Edit `openapi.yaml`
2. Run `just generate-openapi-stubs` to regenerate everything (Go + TS + build)
3. For faster iteration, use partial regeneration:
   - Go models only: `go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.1 --config .oapi-codegen.models.yaml openapi.yaml`
   - Go server only: `go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.1 --config .oapi-codegen.server.yaml openapi.yaml`
   - Go spec only: `go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.1 --config .oapi-codegen.spec.yaml openapi.yaml`
   - TS only: `npx orval --config orval.config.mjs`
   - TS build only: `npm --prefix ui run build`

## Workflow: publish the npm package

```bash
just publish-ui
```

This runs `npm publish` from `ui/`. The `prepublishOnly` script runs the build first.
Package name: `@faber-numeris/authn-api` (unscoped, public access).

## Workflow: sync Bruno test collections

```bash
just sync-api
```

Runs `npx @sayedameer/bruno-openapi-sync` which interactively diffs the OpenAPI spec against the Bruno collection in `tests/`. Only needed when endpoints change.

## Workflow: running Bruno tests

1. Open the `tests/Luciole Authentication Service API` collection in the [Bruno](https://docs.usebruno.com) desktop app or CLI
2. Select an environment: Local (`http://localhost:8080/v1`), Staging, or Production
3. Run individual requests or suites

The register suite chains multiple Brunoscripts (generate random user via Faker → register → confirm via Mailpit → login → logout). It requires:
- Mailpit running at `http://localhost:8025` (default) to fetch email confirmation codes
- The `@faker-js/faker` library available for random user generation

## Architecture notes

| Output | Tool | Config | Consumer |
|---|---|---|---|
| Go models (`api/api.models.go`) | oapi-codegen v2.7.1 | `.oapi-codegen.models.yaml` | Backend repos via `github.com/faber-numeris/foundation` |
| Go chi-server interface (`api/api.server.go`) | oapi-codegen v2.7.1 | `.oapi-codegen.server.yaml` | Backend repos |
| Go embedded spec (`api/api.spec.go`) | oapi-codegen v2.7.1 | `.oapi-codegen.spec.yaml` | Backend repos |
| TS types + fetch client (`ui/src/`) | orval v8.18.0 | `orval.config.mjs` (tags-split mode) | Frontend repos via `@faber-numeris/authn-api` |
| TS Zod schemas (`ui/src/`) | orval v8.18.0 | `orval.config.mjs` (zod client) | Frontend repos |

### Go specifics

- Package name is always `api`
- Server interface follows the `go-chi/chi/v5` router pattern
- Orval generates: z.http client for fetch, types + schemas exported from the barrel `index.ts`

## Sub-skills

| Task | See |
|---|---|
| Adding a new endpoint | Edit `openapi.yaml`, then run full `just generate-openapi-stubs` |
| Adding a new model/schema | Edit the `components/schemas` section in `openapi.yaml`, then regenerate |

## Cross-references

- `AGENTS.md` at the repo root for a compact overview of structure and commands
