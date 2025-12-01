---
title: "docker sandbox ls"
source: "https://docs.docker.com/reference/cli/docker/sandbox/ls/"
---

# docker sandbox ls

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker sandbox ls

DescriptionList sandboxesUsage`docker sandbox ls`Aliases

An alias is a short or memorable alternative for a longer command.

`docker sandbox list`

## [Description](#description)

List all sandboxes.

This command lists all sandboxes using the Docker API.

## [Options](#options)

OptionDefaultDescription[`--no-trunc`](#no-trunc)Don't truncate output[`-q, --quiet`](#quiet)Only display sandbox IDs

## [Examples](#examples)

### [List all sandboxes](#list-all-sandboxes)

```console
$ docker sandbox ls
SANDBOX ID    NAME         WORKSPACE                    CREATED
abc123def     my-project   /home/user/my-project        2 hours ago
def456ghi     ml-work      /home/user/ml-projects       1 day ago
```

### [Show only sandbox IDs (--quiet)](#quiet)

```text
--quiet
```

Output only sandbox IDs:

```console
$ docker sandbox ls --quiet
abc123def
def456ghi
```

### [Don't truncate output (--no-trunc)](#no-trunc)

```text
--no-trunc
```

By default, long sandbox IDs and workspace paths are truncated for readability. Use `--no-trunc` to display the full values:

```console
$ docker sandbox ls
SANDBOX ID    TEMPLATE  NAME         WORKSPACE                     STATUS   CREATED
abc123def456  ubuntu    my-project   /home/user/.../my-project     running  2 hours ago

$ docker sandbox ls --no-trunc
SANDBOX ID              TEMPLATE  NAME         WORKSPACE                                          STATUS   CREATED
abc123def456ghi789jkl   ubuntu    my-project   /home/user/very/long/path/to/my-project           running  2 hours ago
```

Table of contents
