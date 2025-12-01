---
title: "docker model package"
source: "https://docs.docker.com/reference/cli/docker/model/package/"
---

# docker model package

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker model package

DescriptionPackage a GGUF file into a Docker model OCI artifact, with optional licenses.Usage`docker model package --gguf <path> [--license <path>...] [--context-size <tokens>] [--push] MODEL`

## [Description](#description)

Package a GGUF file into a Docker model OCI artifact, with optional licenses. The package is sent to the model-runner, unless --push is specified. When packaging a sharded model --gguf should point to the first shard. All shard files should be siblings and should include the index in the file name (e.g. model-00001-of-00015.gguf).

## [Options](#options)

OptionDefaultDescription`--chat-template`absolute path to chat template file (must be Jinja format)`--context-size`context size in tokens`--gguf`absolute path to gguf file (required)`-l, --license`absolute path to a license file`--push`push to registry (if not set, the model is loaded into the Model Runner content store)

Table of contents
