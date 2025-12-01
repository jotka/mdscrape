---
title: "docker scout repo enable"
source: "https://docs.docker.com/reference/cli/docker/scout/repo/enable/"
---

# docker scout repo enable

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker scout repo enable

DescriptionEnable Docker ScoutUsage`docker scout repo enable [REPOSITORY]`

## [Description](#description)

The docker scout repo enable command enables Docker Scout on repositories.

## [Options](#options)

OptionDefaultDescription`--all`Enable all repositories of the organization. Can not be used with --filter.  
`--filter`Regular expression to filter repositories by name`--integration`Name of the integration to use for enabling an image`--org`Namespace of the Docker organization`--registry`Container Registry

## [Examples](#examples)

### [Enable a specific repository](#enable-a-specific-repository)

```console
$ docker scout repo enable my/repository
```

### [Enable all repositories of the organization](#enable-all-repositories-of-the-organization)

```console
$ docker scout repo enable --all
```

### [Enable some repositories based on a filter](#enable-some-repositories-based-on-a-filter)

```console
$ docker scout repo enable --filter namespace/backend
```

### [Enable a repository from a specific registry](#enable-a-repository-from-a-specific-registry)

```console
$ docker scout repo enable my/repository --registry 123456.dkr.ecr.us-east-1.amazonaws.com
```

Table of contents
