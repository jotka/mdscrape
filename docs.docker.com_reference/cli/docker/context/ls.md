---
title: "docker context ls"
source: "https://docs.docker.com/reference/cli/docker/context/ls/"
---

# docker context ls

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker context ls

DescriptionList contextsUsage`docker context ls [OPTIONS]`Aliases

An alias is a short or memorable alternative for a longer command.

`docker context list`

## [Description](#description)

List contexts

## [Options](#options)

OptionDefaultDescription`--format`Format output using a custom template:  
'table': Print output in table format with column headers (default)  
'table TEMPLATE': Print output in table format using the given Go template  
'json': Print in JSON format  
'TEMPLATE': Print output using the given Go template.  
Refer to [https://docs.docker.com/go/formatting/](https://docs.docker.com/go/formatting/) for more information about formatting output with templates`-q, --quiet`Only show context names

## [Examples](#examples)

Use `docker context ls` to print all contexts. The currently active context is indicated with an `*`:

```console
$ docker context ls

NAME                DESCRIPTION                               DOCKER ENDPOINT                      ORCHESTRATOR
default *           Current DOCKER_HOST based configuration   unix:///var/run/docker.sock          swarm
production                                                    tcp:///prod.corp.example.com:2376
staging                                                       tcp:///stage.corp.example.com:2376
```

Table of contents
