---
title: "docker model pull"
source: "https://docs.docker.com/reference/cli/docker/model/pull/"
---

# docker model pull

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker model pull

DescriptionPull a model from Docker Hub or HuggingFace to your local environmentUsage`docker model pull MODEL`

## [Description](#description)

Pull a model to your local environment. Downloaded models also appear in the Docker Desktop Dashboard.

## [Options](#options)

OptionDefaultDescription`--ignore-runtime-memory-check`Do not block pull if estimated runtime memory for model exceeds system resources.

## [Examples](#examples)

### [Pulling a model from Docker Hub](#pulling-a-model-from-docker-hub)

```console
docker model pull ai/smollm2
```

### [Pulling from HuggingFace](#pulling-from-huggingface)

You can pull GGUF models directly from [Hugging Face](https://huggingface.co/models?library=gguf).

**Note about quantization:** If no tag is specified, the command tries to pull the `Q4_K_M` version of the model. If `Q4_K_M` doesn't exist, the command pulls the first GGUF found in the **Files** view of the model on HuggingFace. To specify the quantization, provide it as a tag, for example: `docker model pull hf.co/bartowski/Llama-3.2-1B-Instruct-GGUF:Q4_K_S`

```console
docker model pull hf.co/bartowski/Llama-3.2-1B-Instruct-GGUF
```

Table of contents
