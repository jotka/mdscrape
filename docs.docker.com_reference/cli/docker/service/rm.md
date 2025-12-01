---
title: "docker service rm"
source: "https://docs.docker.com/reference/cli/docker/service/rm/"
---

# docker service rm

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker service rm

DescriptionRemove one or more servicesUsage`docker service rm SERVICE [SERVICE...]`Aliases

An alias is a short or memorable alternative for a longer command.

`docker service remove`

Swarm This command works with the Swarm orchestrator.

## [Description](#description)

Removes the specified services from the swarm.

> Note
> 
> This is a cluster management command, and must be executed on a swarm manager node. To learn about managers and workers, refer to the [Swarm mode section](/engine/swarm/) in the documentation.

## [Examples](#examples)

Remove the `redis` service:

```console
$ docker service rm redis

redis

$ docker service ls

ID  NAME  MODE  REPLICAS  IMAGE
```

> Warning
> 
> Unlike `docker rm`, this command does not ask for confirmation before removing a running service.

Table of contents
