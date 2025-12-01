---
title: "docker plugin enable"
source: "https://docs.docker.com/reference/cli/docker/plugin/enable/"
---

# docker plugin enable

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker plugin enable

DescriptionEnable a pluginUsage`docker plugin enable [OPTIONS] PLUGIN`

## [Description](#description)

Enables a plugin. The plugin must be installed before it can be enabled, see [`docker plugin install`](/reference/cli/docker/plugin/install/).

## [Options](#options)

OptionDefaultDescription`--timeout``30`HTTP client timeout (in seconds)

## [Examples](#examples)

The following example shows that the `sample-volume-plugin` plugin is installed, but disabled:

```console
$ docker plugin ls

ID            NAME                                    DESCRIPTION                ENABLED
69553ca1d123  tiborvass/sample-volume-plugin:latest   A test plugin for Docker   false
```

To enable the plugin, use the following command:

```console
$ docker plugin enable tiborvass/sample-volume-plugin

tiborvass/sample-volume-plugin

$ docker plugin ls

ID            NAME                                    DESCRIPTION                ENABLED
69553ca1d123  tiborvass/sample-volume-plugin:latest   A test plugin for Docker   true
```

Table of contents
