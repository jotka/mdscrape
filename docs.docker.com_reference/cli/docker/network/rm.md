---
title: "docker network rm"
source: "https://docs.docker.com/reference/cli/docker/network/rm/"
---

# docker network rm

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker network rm

DescriptionRemove one or more networksUsage`docker network rm NETWORK [NETWORK...]`Aliases

An alias is a short or memorable alternative for a longer command.

`docker network remove`

## [Description](#description)

Removes one or more networks by name or identifier. To remove a network, you must first disconnect any containers connected to it.

## [Options](#options)

OptionDefaultDescription`-f, --force`Do not error if the network does not exist

## [Examples](#examples)

### [Remove a network](#remove-a-network)

To remove the network named 'my-network':

```console
$ docker network rm my-network
```

### [Remove multiple networks](#remove-multiple-networks)

To delete multiple networks in a single `docker network rm` command, provide multiple network names or ids. The following example deletes a network with id `3695c422697f` and a network named `my-network`:

```console
$ docker network rm 3695c422697f my-network
```

When you specify multiple networks, the command attempts to delete each in turn. If the deletion of one network fails, the command continues to the next on the list and tries to delete that. The command reports success or failure for each deletion.

Table of contents
