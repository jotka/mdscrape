---
title: "docker scout cache prune"
source: "https://docs.docker.com/reference/cli/docker/scout/cache/prune/"
---

# docker scout cache prune

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker scout cache prune

DescriptionRemove temporary or cached dataUsage`docker scout cache prune`

## [Description](#description)

The `docker scout cache prune` command removes temporary data and SBOM cache.

By default, `docker scout cache prune` only deletes temporary data. To delete temporary data and clear the SBOM cache, use the `--sboms` flag.

## [Options](#options)

OptionDefaultDescription`-f, --force`Do not prompt for confirmation`--sboms`Prune cached SBOMs

## [Examples](#examples)

### [Delete temporary data](#delete-temporary-data)

```console
$ docker scout cache prune
? Are you sure to delete all temporary data? Yes
    ✓ temporary data deleted
```

### [Delete temporary *and* cache data](#delete-temporary-_and_-cache-data)

```console
$ docker scout cache prune --sboms
? Are you sure to delete all temporary data and all cached SBOMs? Yes
    ✓ temporary data deleted
    ✓ cached SBOMs deleted
```

Table of contents
