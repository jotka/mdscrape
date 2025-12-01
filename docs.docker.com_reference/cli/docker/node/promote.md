---
title: "docker node promote"
source: "https://docs.docker.com/reference/cli/docker/node/promote/"
---

# docker node promote

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker node promote

DescriptionPromote one or more nodes to manager in the swarmUsage`docker node promote NODE [NODE...]`

Swarm This command works with the Swarm orchestrator.

## [Description](#description)

Promotes a node to manager. This command can only be executed on a manager node.

> Note
> 
> This is a cluster management command, and must be executed on a swarm manager node. To learn about managers and workers, refer to the [Swarm mode section](/engine/swarm/) in the documentation.

## [Examples](#examples)

```console
$ docker node promote <node name>
```

Table of contents
