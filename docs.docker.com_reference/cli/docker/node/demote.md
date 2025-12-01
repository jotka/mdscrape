---
title: "docker node demote"
source: "https://docs.docker.com/reference/cli/docker/node/demote/"
---

# docker node demote

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker node demote

DescriptionDemote one or more nodes from manager in the swarmUsage`docker node demote NODE [NODE...]`

Swarm This command works with the Swarm orchestrator.

## [Description](#description)

Demotes an existing manager so that it is no longer a manager.

> Note
> 
> This is a cluster management command, and must be executed on a swarm manager node. To learn about managers and workers, refer to the [Swarm mode section](/engine/swarm/) in the documentation.

## [Examples](#examples)

```console
$ docker node demote <node name>
```

Table of contents
