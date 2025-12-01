---
title: "docker sandbox rm"
source: "https://docs.docker.com/reference/cli/docker/sandbox/rm/"
---

# docker sandbox rm

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker sandbox rm

DescriptionRemove one or more sandboxesUsage`docker sandbox rm [OPTIONS] SANDBOX [SANDBOX...]`

## [Description](#description)

Remove one or more sandboxes by their IDs or names.

This command removes the specified sandboxes. Each sandbox is identified by its unique ID or name.

## [Examples](#examples)

### [Remove a sandbox](#remove-a-sandbox)

```console
$ docker sandbox rm abc123def
abc123def
```

### [Remove multiple sandboxes](#remove-multiple-sandboxes)

```console
$ docker sandbox rm abc123def def456ghi
abc123def
def456ghi
```

### [Remove all sandboxes](#remove-all-sandboxes)

```console
$ docker sandbox rm $(docker sandbox ls -q)
```

Table of contents
