# semrel-api

Shared gRPC API module for [SemRel](https://github.com/SemRels/semrel) plugins.

## Overview

This module provides:

- **Proto definitions** (`api/proto/v1/`) — the canonical gRPC contracts for all SemRel plugin types
- **Generated Go stubs** (`api/gen/v1/`) — protobuf message types and gRPC service interfaces
- **go-plugin wiring** (`plugin/`) — [hashicorp/go-plugin](https://github.com/hashicorp/go-plugin) `GRPCPlugin` implementations and shared `HandshakeConfig`

## Plugin Types

| Key         | Service                    | Purpose                                      |
|-------------|----------------------------|----------------------------------------------|
| `provider`  | `ProviderPlugin`           | VCS platform (GitHub, GitLab, Gitea, git)    |
| `condition` | `CIConditionPlugin`        | Verify CI environment before releasing       |
| `analyzer`  | `CommitAnalyzerPlugin`     | Determine semver bump from commits           |
| `generator` | `ChangelogGeneratorPlugin` | Render changelog / release notes             |
| `updater`   | `FilesUpdaterPlugin`       | Write new version into project files         |
| `hooks`     | `HooksPlugin`              | Lifecycle callbacks (success / failure)      |

## Usage

### Host (semrel tool)

```go
import (
    semrelapi "github.com/GoSemantics/go-semrel-api/plugin"
    "github.com/hashicorp/go-plugin"
)

client := plugin.NewClient(&plugin.ClientConfig{
    HandshakeConfig: semrelapi.HandshakeConfig,
    Plugins:         semrelapi.PluginMap,
    Cmd:             exec.Command("./provider-github"),
    AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
})
```

### Plugin binary

```go
import (
    semrelapi "github.com/GoSemantics/go-semrel-api/plugin"
    semrelv1  "github.com/GoSemantics/go-semrel-api/api/gen/v1"
    "github.com/hashicorp/go-plugin"
)

plugin.Serve(&plugin.ServeConfig{
    HandshakeConfig: semrelapi.HandshakeConfig,
    Plugins: map[string]plugin.Plugin{
        "provider": &semrelapi.ProviderGRPCPlugin{Impl: &MyProviderImpl{}},
    },
    GRPCServer: plugin.DefaultGRPCServer,
})
```

## Regenerating Proto Code

```bash
buf generate
go mod tidy
```

## License

Apache-2.0 — see [LICENSE](LICENSE).
