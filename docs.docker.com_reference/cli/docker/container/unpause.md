---
title: "docker container unpause"
source: "https://docs.docker.com/reference/cli/docker/container/unpause/"
---

# docker container unpause

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker container unpause

DescriptionUnpause all processes within one or more containersUsage`docker container unpause CONTAINER [CONTAINER...]`Aliases

An alias is a short or memorable alternative for a longer command.

`docker unpause`

## [Description](#description)

The `docker unpause` command un-suspends all processes in the specified containers. On Linux, it does this using the freezer cgroup.

See the [freezer cgroup documentation](https://www.kernel.org/doc/Documentation/cgroup-v1/freezer-subsystem.txt) for further details.

## [Examples](#examples)

```console
$ docker unpause my_container
my_container
```

Table of contents
