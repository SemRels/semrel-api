# semrel-api

Shared proto and gRPC API definitions for [semrel](https://github.com/SemRels/semrel) plugins.

> **Note**: The current semrel plugin architecture communicates via **environment variables and subprocess execution** — not gRPC. This repository contains the proto definitions for reference and potential future use.

## Contents

- `api/proto/v1/` — Protobuf service definitions for all semrel plugin types
- `api/gen/v1/` — Generated Go stubs (protobuf messages and gRPC service interfaces)
- `plugin/` — [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin) wiring (legacy)

## Plugin types

| Key | Service | Purpose |
|---|---|---|
| `provider` | `ProviderPlugin` | VCS platform (GitHub, GitLab, Gitea, git) |
| `condition` | `CIConditionPlugin` | Verify CI environment before releasing |
| `analyzer` | `CommitAnalyzerPlugin` | Determine semver bump from commits |
| `generator` | `ChangelogGeneratorPlugin` | Render changelog / release notes |
| `updater` | `FilesUpdaterPlugin` | Write new version into project files |
| `hooks` | `HooksPlugin` | Lifecycle callbacks (success / failure) |

## Current plugin contract

semrel plugins today are standalone executables launched as subprocesses. They communicate via:

- **Environment variables** → plugin: `SEMREL_*` (release context) and `SEMREL_PLUGIN_*` (config from `args:`)
- **stdout**: JSON result (analyzers only)
- **stderr**: logs and `plugin_schema_version=N` announcement
- **Exit code**: 0 = success, non-zero = abort release

See the [plugin development guide](https://semrel.io/guide/plugin-development) for a full walkthrough.

## Module

```
github.com/SemRels/semrel-api
```

## Regenerating proto code

```bash
buf generate
go mod tidy
```

## License

Apache-2.0 — see [LICENSE](LICENSE).
