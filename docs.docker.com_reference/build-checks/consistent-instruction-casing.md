---
title: "ConsistentInstructionCasing"
source: "https://docs.docker.com/reference/build-checks/consistent-instruction-casing/"
---

# ConsistentInstructionCasing

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# ConsistentInstructionCasing

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
Command 'EntryPoint' should be consistently cased
```

## [Description](#description)

Instruction keywords should use consistent casing (all lowercase or all uppercase). Using a case that mixes uppercase and lowercase, such as `PascalCase` or `snakeCase`, letters result in poor readability.

## [Examples](#examples)

❌ Bad: don't mix uppercase and lowercase.

```dockerfile
From alpine
Run echo hello > /greeting.txt
EntRYpOiNT ["cat", "/greeting.txt"]
```

✅ Good: all uppercase.

```dockerfile
FROM alpine
RUN echo hello > /greeting.txt
ENTRYPOINT ["cat", "/greeting.txt"]
```

✅ Good: all lowercase.

```dockerfile
from alpine
run echo hello > /greeting.txt
entrypoint ["cat", "/greeting.txt"]
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fconsistent-instruction-casing%2F&labels=status%2Ftriage)

Table of contents
