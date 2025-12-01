---
title: "SecretsUsedInArgOrEnv"
source: "https://docs.docker.com/reference/build-checks/secrets-used-in-arg-or-env/"
---

# SecretsUsedInArgOrEnv

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# SecretsUsedInArgOrEnv

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

## [Output](#output)

```text
Potentially sensitive data should not be used in the ARG or ENV commands
```

## [Description](#description)

While it is common to pass secrets to running processes through environment variables during local development, setting secrets in a Dockerfile using `ENV` or `ARG` is insecure because they persist in the final image. This rule reports violations where `ENV` and `ARG` keys indicate that they contain sensitive data.

Instead of `ARG` or `ENV`, you should use secret mounts, which expose secrets to your builds in a secure manner, and do not persist in the final image or its metadata. See [Build secrets](https://docs.docker.com/build/building/secrets/).

## [Examples](#examples)

❌ Bad: `AWS_SECRET_ACCESS_KEY` is a secret value.

```dockerfile
FROM scratch
ARG AWS_SECRET_ACCESS_KEY
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fsecrets-used-in-arg-or-env%2F&labels=status%2Ftriage)

Table of contents
