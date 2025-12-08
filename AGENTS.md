# Repository Guidelines

## Project Structure & Module Organization
Stateset Core follows Cosmos SDK layering. Application wiring sits in `app/`, the `statesetd` CLI lives in `cmd/statesetd/`, and feature modules are in `x/<module>/` with `keeper` and `types` packages. Shared helpers stay in `utils/`, `tools/`, and `testutil/`, generated protobuf assets in `proto/`, and smart-contract code in `contracts/` and `cosmwasm/`. Build outputs land in `build/`, while automation and deployment scripts sit in `scripts/` and the top-level shell helpers.

## Build, Test, and Development Commands
- `make build` compiles `statesetd` into `build/statesetd` with the repository’s default build tags and ldflags.
- `make install` installs the binary into `$GOPATH/bin` for downstream tooling.
- `make dev` rebuilds, initializes, and launches a single-node network for rapid manual testing.
- `make test`, `make test-all`, and `make test-cover` wrap `go test` with the ledger tags; `make test-cover` writes `coverage.txt`.
- `go build -o build/statesetd ./cmd/statesetd` and `go test ./...` remain useful for quick, package-scoped iterations.

## Coding Style & Naming Conventions
Run `gofmt` (tabs, grouped imports) before committing. Package names stay lowercase and match their directories, while exported identifiers use CamelCase and shared constants use `ALL_CAPS`. Mirror YAML or JSON field names when introducing configuration keys. Use `fix_imports.sh` when SDK or CometBFT imports move, and regenerate protobuf output so `proto/` stays aligned with generated Go.

## Testing Guidelines
Place unit tests beside their packages with the `_test.go` suffix and `Test<Name>` functions so `go test ./...` can find them. Default Make targets add `ledger` mocks; use `make test` for quick feedback and `make test-cover` when refreshing coverage baselines. Integration suites under `tests/` expect a live node, so run `make dev` in another terminal first.

## Commit & Pull Request Guidelines
History favors short, present-tense summaries (for example, `updating api`); continue writing imperative subject lines such as `Add escrow module invariants`, adding a brief body when migrations or schema updates are involved. Each pull request should state intent, link related issues (`Closes #123`), and include test evidence or coverage notes. Mention configuration or Docker changes so reviewers can reproduce the setup, and attach logs or screenshots when touching operational tooling.

## Security & Configuration Tips
Never commit validator keys or `.statesetd` data—keep secrets in local keyrings and reference `config/` templates. After protobuf or ABCI changes, regenerate code and run `docker-build.sh` or the relevant `docker-compose*.yml` stack to confirm the containerized path still builds.
