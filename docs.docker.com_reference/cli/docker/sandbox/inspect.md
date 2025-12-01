---
title: "docker sandbox inspect"
source: "https://docs.docker.com/reference/cli/docker/sandbox/inspect/"
---

# docker sandbox inspect

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker sandbox inspect

DescriptionDisplay detailed information on one or more sandboxesUsage`docker sandbox inspect [OPTIONS] SANDBOX [SANDBOX...]`

## [Description](#description)

Display detailed information on one or more sandboxes.

This command retrieves and displays detailed information about the specified sandboxes using the Docker API. Each sandbox is identified by its unique ID or name.

## [Examples](#examples)

### [Inspect a sandbox](#inspect-a-sandbox)

```console
$ docker sandbox inspect abc123def
[
  {
    "id": "abc123def69b16c5c0dab4cf699e26f8d01e1ace3aeee06254e0999492e11647",
    "name": "claude-sandbox-2025-11-04-170333",
    "created_at": "2025-11-04T16:03:33.910642347Z",
    "status": "running",
    "template": "docker/sandbox-templates:claude-code",
    "labels": {
      "com.docker.sandbox.agent": "claude",
      "com.docker.sandbox.workingDirectory": "/Users/moby/code/docker/sandboxes",
      "com.docker.sandbox.workingDirectoryInode": "3041007",
      "com.docker.sandboxes": "templates",
      "com.docker.sandboxes.base": "ubuntu:questing",
      "com.docker.sandboxes.flavor": "claude-code",
      "com.docker.sdk": "true",
      "com.docker.sdk.client": "0.1.0-alpha011",
      "com.docker.sdk.container": "0.1.0-alpha012",
      "com.docker.sdk.lang": "go",
      "docker/sandbox": "true",
      "org.opencontainers.image.ref.name": "ubuntu",
      "org.opencontainers.image.version": "25.10"
    }
  }
]
```

Table of contents
