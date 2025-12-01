---
title: "docker plugin disable"
source: "https://docs.docker.com/reference/cli/docker/plugin/disable/"
---

# docker plugin disable

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker plugin disable

DescriptionDisable a pluginUsage`docker plugin disable [OPTIONS] PLUGIN`

## [Description](#description)

Disables a plugin. The plugin must be installed before it can be disabled, see [`docker plugin install`](/reference/cli/docker/plugin/install/). Without the `-f` option, a plugin that has references (e.g., volumes, networks) cannot be disabled.

## [Options](#options)

OptionDefaultDescription`-f, --force`Force the disable of an active plugin

## [Examples](#examples)

The following example shows that the `sample-volume-plugin` plugin is installed and enabled:

```console
$ docker plugin ls

ID            NAME                                    DESCRIPTION                ENABLED
69553ca1d123  tiborvass/sample-volume-plugin:latest   A test plugin for Docker   true
```

To disable the plugin, use the following command:

```console
$ docker plugin disable tiborvass/sample-volume-plugin

tiborvass/sample-volume-plugin

$ docker plugin ls

ID            NAME                                    DESCRIPTION                ENABLED
69553ca1d123  tiborvass/sample-volume-plugin:latest   A test plugin for Docker   false
```

Table of contents
