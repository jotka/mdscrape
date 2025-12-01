---
title: "docker model run"
source: "https://docs.docker.com/reference/cli/docker/model/run/"
---

# docker model run

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker model run

DescriptionRun a model and interact with it using a submitted prompt or chat modeUsage`docker model run MODEL [PROMPT]`

## [Description](#description)

When you run a model, Docker calls an inference server API endpoint hosted by the Model Runner through Docker Desktop. The model stays in memory until another model is requested, or until a pre-defined inactivity timeout is reached (currently 5 minutes).

You do not have to use Docker model run before interacting with a specific model from a host process or from within a container. Model Runner transparently loads the requested model on-demand, assuming it has been pulled and is locally available.

You can also use chat mode in the Docker Desktop Dashboard when you select the model in the **Models** tab.

## [Options](#options)

OptionDefaultDescription`--color``auto`Use colored output (auto|yes|no)`--debug`Enable debug logging`--ignore-runtime-memory-check`Do not block pull if estimated runtime memory for model exceeds system resources.

## [Examples](#examples)

### [One-time prompt](#one-time-prompt)

```console
docker model run ai/smollm2 "Hi"
```

Output:

```console
Hello! How can I assist you today?
```

### [Interactive chat](#interactive-chat)

```console
docker model run ai/smollm2
```

Output:

```console
Interactive chat mode started. Type '/bye' to exit.
> Hi
Hi there! It's SmolLM, AI assistant. How can I help you today?
> /bye
Chat session ended.
```

Table of contents
