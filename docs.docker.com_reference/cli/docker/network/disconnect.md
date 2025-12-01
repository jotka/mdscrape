---
title: "docker network disconnect"
source: "https://docs.docker.com/reference/cli/docker/network/disconnect/"
---

# docker network disconnect

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker network disconnect

DescriptionDisconnect a container from a networkUsage`docker network disconnect [OPTIONS] NETWORK CONTAINER`

## [Description](#description)

Disconnects a container from a network. The container must be running to disconnect it from the network.

## [Options](#options)

OptionDefaultDescription`-f, --force`Force the container to disconnect from a network

## [Examples](#examples)

```console
$ docker network disconnect multi-host-network container1
```

Table of contents
