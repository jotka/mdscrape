---
title: "docker container wait"
source: "https://docs.docker.com/reference/cli/docker/container/wait/"
---

# docker container wait

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker container wait

DescriptionBlock until one or more containers stop, then print their exit codesUsage`docker container wait CONTAINER [CONTAINER...]`Aliases

An alias is a short or memorable alternative for a longer command.

`docker wait`

## [Description](#description)

Block until one or more containers stop, then print their exit codes

## [Examples](#examples)

Start a container in the background.

```console
$ docker run -dit --name=my_container ubuntu bash
```

Run `docker wait`, which should block until the container exits.

```console
$ docker wait my_container
```

In another terminal, stop the first container. The `docker wait` command above returns the exit code.

```console
$ docker stop my_container
```

This is the same `docker wait` command from above, but it now exits, returning `0`.

```console
$ docker wait my_container

0
```

Table of contents
