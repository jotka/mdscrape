---
title: "docker swarm update"
source: "https://docs.docker.com/reference/cli/docker/swarm/update/"
---

# docker swarm update

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker swarm update

DescriptionUpdate the swarmUsage`docker swarm update [OPTIONS]`

Swarm This command works with the Swarm orchestrator.

## [Description](#description)

Updates a swarm with new parameter values.

> Note
> 
> This is a cluster management command, and must be executed on a swarm manager node. To learn about managers and workers, refer to the [Swarm mode section](/engine/swarm/) in the documentation.

## [Options](#options)

OptionDefaultDescription`--autolock`Change manager autolocking setting (true|false)`--cert-expiry``2160h0m0s`Validity period for node certificates (ns|us|ms|s|m|h)`--dispatcher-heartbeat``5s`Dispatcher heartbeat period (ns|us|ms|s|m|h)`--external-ca`Specifications of one or more certificate signing endpoints`--max-snapshots`API 1.25+ Number of additional Raft snapshots to retain`--snapshot-interval``10000`API 1.25+ Number of log entries between Raft snapshots`--task-history-limit``5`Task history retention limit

## [Examples](#examples)

```console
$ docker swarm update --cert-expiry 720h
```

Table of contents
