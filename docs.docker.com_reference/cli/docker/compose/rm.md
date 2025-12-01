---
title: "docker compose rm"
source: "https://docs.docker.com/reference/cli/docker/compose/rm/"
---

# docker compose rm

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker compose rm

DescriptionRemoves stopped service containersUsage`docker compose rm [OPTIONS] [SERVICE...]`

## [Description](#description)

Removes stopped service containers.

By default, anonymous volumes attached to containers are not removed. You can override this with `-v`. To list all volumes, use `docker volume ls`.

Any data which is not in a volume is lost.

Running the command with no options also removes one-off containers created by `docker compose run`:

```console
$ docker compose rm
Going to remove djangoquickstart_web_run_1
Are you sure? [yN] y
Removing djangoquickstart_web_run_1 ... done
```

## [Options](#options)

OptionDefaultDescription`-f, --force`Don't ask to confirm removal`-s, --stop`Stop the containers, if required, before removing`-v, --volumes`Remove any anonymous volumes attached to containers

Table of contents
