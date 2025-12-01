---
title: "CopyIgnoredFile"
source: "https://docs.docker.com/reference/build-checks/copy-ignored-file/"
---

# CopyIgnoredFile

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# CopyIgnoredFile

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

> Note
> 
> This check is experimental and is not enabled by default. To enable it, see [Experimental checks](https://docs.docker.com/go/build-checks-experimental/).

## [Output](#output)

```text
Attempting to Copy file "./tmp/Dockerfile" that is excluded by .dockerignore
```

## [Description](#description)

When you use the Add or Copy instructions from within a Dockerfile, you should ensure that the files to be copied into the image do not match a pattern present in `.dockerignore`.

Files which match the patterns in a `.dockerignore` file are not present in the context of the image when it is built. Trying to copy or add a file which is missing from the context will result in a build error.

## [Examples](#examples)

With the given `.dockerignore` file:

```text
*/tmp/*
```

❌ Bad: Attempting to Copy file "./tmp/Dockerfile" that is excluded by .dockerignore

```dockerfile
FROM scratch
COPY ./tmp/helloworld.txt /helloworld.txt
```

✅ Good: Copying a file which is not excluded by .dockerignore

```dockerfile
FROM scratch
COPY ./forever/helloworld.txt /helloworld.txt
```

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fbuild-checks%2Fcopy-ignored-file%2F&labels=status%2Ftriage)

Table of contents
