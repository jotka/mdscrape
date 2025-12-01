---
title: "docker plugin push"
source: "https://docs.docker.com/reference/cli/docker/plugin/push/"
---

# docker plugin push

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker plugin push

DescriptionPush a plugin to a registryUsage`docker plugin push [OPTIONS] PLUGIN[:TAG]`

## [Description](#description)

After you have created a plugin using `docker plugin create` and the plugin is ready for distribution, use `docker plugin push` to share your images to Docker Hub or a self-hosted registry.

Registry credentials are managed by [docker login](/reference/cli/docker/login/).

## [Options](#options)

OptionDefaultDescription`--disable-content-trust``true`Skip image signing

## [Examples](#examples)

The following example shows how to push a sample `user/plugin`.

```console
$ docker plugin ls

ID             NAME                    DESCRIPTION                  ENABLED
69553ca1d456   user/plugin:latest      A sample plugin for Docker   false

$ docker plugin push user/plugin
```

Table of contents
