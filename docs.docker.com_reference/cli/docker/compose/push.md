---
title: "docker compose push"
source: "https://docs.docker.com/reference/cli/docker/compose/push/"
---

# docker compose push

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker compose push

DescriptionPush service imagesUsage`docker compose push [OPTIONS] [SERVICE...]`

## [Description](#description)

Pushes images for services to their respective registry/repository.

The following assumptions are made:

- You are pushing an image you have built locally
- You have access to the build key

Examples

```yaml
services:
  service1:
    build: .
    image: localhost:5000/yourimage  ## goes to local registry

  service2:
    build: .
    image: your-dockerid/yourimage  ## goes to your repository on Docker Hub
```

## [Options](#options)

OptionDefaultDescription`--ignore-push-failures`Push what it can and ignores images with push failures`--include-deps`Also push images of services declared as dependencies`-q, --quiet`Push without printing progress information

Table of contents
