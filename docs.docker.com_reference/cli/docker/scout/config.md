---
title: "docker scout config"
source: "https://docs.docker.com/reference/cli/docker/scout/config/"
---

# docker scout config

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker scout config

DescriptionManage Docker Scout configurationUsage`docker scout config [KEY] [VALUE]`

## [Description](#description)

`docker scout config` allows you to list, get and set Docker Scout configuration.

Available configuration key:

- `organization`: Namespace of the Docker organization to be used by default.

## [Examples](#examples)

### [List existing configuration](#list-existing-configuration)

```console
$ docker scout config
organization=my-org-namespace
```

### [Print configuration value](#print-configuration-value)

```console
$ docker scout config organization
my-org-namespace
```

### [Set configuration value](#set-configuration-value)

```console
$ docker scout config organization my-org-namespace
    ✓ Successfully set organization to my-org-namespace
```

Table of contents
