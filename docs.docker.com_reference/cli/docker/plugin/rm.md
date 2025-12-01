---
title: "docker plugin rm"
source: "https://docs.docker.com/reference/cli/docker/plugin/rm/"
---

# docker plugin rm

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker plugin rm

DescriptionRemove one or more pluginsUsage`docker plugin rm [OPTIONS] PLUGIN [PLUGIN...]`Aliases

An alias is a short or memorable alternative for a longer command.

`docker plugin remove`

## [Description](#description)

Removes a plugin. You cannot remove a plugin if it is enabled, you must disable a plugin using the [`docker plugin disable`](/reference/cli/docker/plugin/disable/) before removing it, or use `--force`. Use of `--force` is not recommended, since it can affect functioning of running containers using the plugin.

## [Options](#options)

OptionDefaultDescription`-f, --force`Force the removal of an active plugin

## [Examples](#examples)

The following example disables and removes the `sample-volume-plugin:latest` plugin:

```console
$ docker plugin disable tiborvass/sample-volume-plugin

tiborvass/sample-volume-plugin

$ docker plugin rm tiborvass/sample-volume-plugin:latest

tiborvass/sample-volume-plugin
```

Table of contents
