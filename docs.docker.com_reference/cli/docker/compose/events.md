---
title: "docker compose events"
source: "https://docs.docker.com/reference/cli/docker/compose/events/"
---

# docker compose events

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker compose events

DescriptionReceive real time events from containersUsage`docker compose events [OPTIONS] [SERVICE...]`

## [Description](#description)

Stream container events for every container in the project.

With the `--json` flag, a json object is printed one per line with the format:

```json
{
    "time": "2015-11-20T18:01:03.615550",
    "type": "container",
    "action": "create",
    "id": "213cf7...5fc39a",
    "service": "web",
    "attributes": {
      "name": "application_web_1",
      "image": "alpine:edge"
    }
}
```

The events that can be received using this can be seen [here](/reference/cli/docker/system/events/#object-types).

## [Options](#options)

OptionDefaultDescription`--json`Output events as a stream of json objects`--since`Show all events created since timestamp`--until`Stream events until this timestamp

Table of contents
