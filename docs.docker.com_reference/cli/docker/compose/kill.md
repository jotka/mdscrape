---
title: "docker compose kill"
source: "https://docs.docker.com/reference/cli/docker/compose/kill/"
---

# docker compose kill

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker compose kill

DescriptionForce stop service containersUsage`docker compose kill [OPTIONS] [SERVICE...]`

## [Description](#description)

Forces running containers to stop by sending a `SIGKILL` signal. Optionally the signal can be passed, for example:

```console
$ docker compose kill -s SIGINT
```

## [Options](#options)

OptionDefaultDescription`--remove-orphans`Remove containers for services not defined in the Compose file`-s, --signal``SIGKILL`SIGNAL to send to the container

Table of contents
