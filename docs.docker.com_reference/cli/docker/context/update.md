---
title: "docker context update"
source: "https://docs.docker.com/reference/cli/docker/context/update/"
---

# docker context update

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker context update

DescriptionUpdate a contextUsage`docker context update [OPTIONS] CONTEXT`

## [Description](#description)

Updates an existing `context`. See [context create](/reference/cli/docker/context/create/).

## [Options](#options)

OptionDefaultDescription`--description`Description of the context`--docker`set the docker endpoint

## [Examples](#examples)

### [Update an existing context](#update-an-existing-context)

```console
$ docker context update \
    --description "some description" \
    --docker "host=tcp://myserver:2376,ca=~/ca-file,cert=~/cert-file,key=~/key-file" \
    my-context
```

Table of contents
