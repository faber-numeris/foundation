# AGENTS.md — foundation/authn

**Luciole Authentication Service API** — shared SDK published for all faber-numeris services (backend Go, frontend TS).

## Repo structure

```
authn/
  openapi.yaml                      # Single source of truth (OpenAPI 3.1)
  justfile                          # All commands run through `just`
  .oapi-codegen.{models,server,spec}.yaml  # Go codegen configs (oapi-codegen v2.7.1)
  orval.config.mjs                  # TypeScript codegen config (orval, fetch client + Zod)
  api/                              # Generated Go stubs (package `api`, chi-server interface)
  ui/                               # npm package @faber-numeris/authn-api (published)
  tests/                            # Bruno API test collections
```

## Key facts

- **`openapi.yaml` is the API contract** shared across all faber-numeris projects (backend + frontend).
- **Spec-first workflow.** Edit `openapi.yaml`, then regenerate stubs — both Go and TS outputs are consumed by sibling repos.
- **No `go.mod` in this repo.** Generated Go code is imported via the `github.com/faber-numeris/foundation` module path in backend repos.
- **TypeScript client** uses native `fetch` (not axios), generates Zod schemas for validation.
- **npm package** `@faber-numeris/authn-api` is published publicly (unscoped), imported by front-end projects.

## Commands (all via `just`)

| Command | What it does |
|---|---|
| `just generate-openapi-stubs` | Regenerate Go + TS stubs from `openapi.yaml`, then build TS package |
| `just publish-ui` | Publish `ui/` to npm |
| `just sync-api` | Sync Bruno test collection from `openapi.yaml` (interactive diff) |

Partial regeneration (for speed):
- Go only: `go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.7.1 --config .oapi-codegen.$CONFIG.yaml openapi.yaml` (where `$CONFIG` is `models`, `server`, or `spec`)
- TS only: `npx orval --config orval.config.mjs`
- TS build only: `npm --prefix ui run build` (runs `tsc` via `npm run clean && tsc`)

## Generated code — never edit by hand

Hands-off directories/files:
- `authn/api/*.go` — oapi-codegen output
- `authn/ui/src/api/**` — orval output (tags-split mode)
- `authn/ui/src/models/**` — orval output
- `authn/ui/dist/**` — TypeScript build artifacts

All regenerated on every `just generate-openapi-stubs` run.

## Bruno tests

Test collections in `authn/tests/` are run with the [Bruno](https://docs.usebruno.com) API client.
Three environments available: Local (`http://localhost:8080/v1`), Staging, Production.
Some suites (e.g. register flow) use scripts to chain requests (generate random user → register → confirm → login → logout).
The register suite integrates with Mailpit to fetch confirmation codes from emails.


