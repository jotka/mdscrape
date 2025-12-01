---
title: "docker compose create"
source: "https://docs.docker.com/reference/cli/docker/compose/create/"
---

# docker compose create

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker compose create

DescriptionCreates containers for a serviceUsage`docker compose create [OPTIONS] [SERVICE...]`

## [Description](#description)

Creates containers for a service

## [Options](#options)

OptionDefaultDescription`--build`Build images before starting containers`--force-recreate`Recreate containers even if their configuration and image haven't changed  
`--no-build`Don't build an image, even if it's policy`--no-recreate`If containers already exist, don't recreate them. Incompatible with --force-recreate.  
`--pull``policy`Pull image before running ("always"|"missing"|"never"|"build")`--quiet-pull`Pull without printing progress information`--remove-orphans`Remove containers for services not defined in the Compose file`--scale`Scale SERVICE to NUM instances. Overrides the `scale` setting in the Compose file if present.  
`-y, --yes`Assume "yes" as answer to all prompts and run non-interactively

Table of contents
