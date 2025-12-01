---
title: "docker mcp policy set"
source: "https://docs.docker.com/reference/cli/docker/mcp/policy/policy_set/"
---

# docker mcp policy set

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# docker mcp policy set

DescriptionSet a policy for secret management in Docker DesktopUsage`docker mcp policy set <content>`

## [Description](#description)

Set a policy for secret management in Docker Desktop

## [Examples](#examples)

### [Backup the current policy to a file](#backup-the-current-policy-to-a-file)

docker mcp policy dump &gt; policy.conf

### [Set a new policy](#set-a-new-policy)

docker mcp policy set "my-secret allows postgres"

### [Restore the previous policy](#restore-the-previous-policy)

cat policy.conf | docker mcp policy set

Table of contents
