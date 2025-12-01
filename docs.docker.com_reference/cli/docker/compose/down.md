---
title: "docker compose down"
source: "https://docs.docker.com/reference/cli/docker/compose/down/"
---

# docker compose down

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker compose down

DescriptionStop and remove containers, networksUsage`docker compose down [OPTIONS] [SERVICES]`

## [Description](#description)

Stops containers and removes containers, networks, volumes, and images created by `up`.

By default, the only things removed are:

- Containers for services defined in the Compose file.
- Networks defined in the networks section of the Compose file.
- The default network, if one is used.

Networks and volumes defined as external are never removed.

Anonymous volumes are not removed by default. However, as they don’t have a stable name, they are not automatically mounted by a subsequent `up`. For data that needs to persist between updates, use explicit paths as bind mounts or named volumes.

## [Options](#options)

OptionDefaultDescription`--remove-orphans`Remove containers for services not defined in the Compose file`--rmi`Remove images used by services. "local" remove only images that don't have a custom tag ("local"|"all")  
`-t, --timeout`Specify a shutdown timeout in seconds`-v, --volumes`Remove named volumes declared in the "volumes" section of the Compose file and anonymous volumes attached to containers

Table of contents
