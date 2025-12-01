---
title: "Interface: Docker"
source: "https://docs.docker.com/reference/api/extensions-sdk/Docker/"
---

# Interface: Docker

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: Docker

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

**`Since`**

0.2.0

## [Properties](#properties)

### [cli](#cli)

• `Readonly` **cli**: [`DockerCommand`](https://docs.docker.com/reference/api/extensions-sdk/DockerCommand/)

You can also directly execute the Docker binary.

```typescript
const output = await ddClient.docker.cli.exec("volume", [
  "ls",
  "--filter",
  "dangling=true"
]);
```

Output:

```json
{
  "stderr": "...",
  "stdout": "..."
}
```

For convenience, the command result object also has methods to easily parse it depending on output format. See [ExecResult](https://docs.docker.com/reference/api/extensions-sdk/ExecResult/) instead.

* * *

Streams the output as a result of the execution of a Docker command. It is useful when the output of the command is too long, or you need to get the output as a stream.

```typescript
await ddClient.docker.cli.exec("logs", ["-f", "..."], {
  stream: {
    onOutput(data): void {
        // As we can receive both `stdout` and `stderr`, we wrap them in a JSON object
        JSON.stringify(
          {
            stdout: data.stdout,
            stderr: data.stderr,
          },
          null,
          "  "
        );
    },
    onError(error: any): void {
      console.error(error);
    },
    onClose(exitCode: number): void {
      console.log("onClose with exit code " + exitCode);
    },
  },
});
```

## [Methods](#methods)

### [listContainers](#listcontainers)

▸ **listContainers**(`options?`): `Promise`&lt;`unknown`&gt;

Get the list of running containers (same as `docker ps`).

By default, this will not list stopped containers. You can use the option `{"all": true}` to list all the running and stopped containers.

```typescript
const containers = await ddClient.docker.listContainers();
```

#### [Parameters](#parameters)

NameTypeDescription`options?``any`(Optional). A JSON like `{ "all": true, "limit": 10, "size": true, "filters": JSON.stringify({ status: ["exited"] }), }` For more information about the different properties see [the Docker API endpoint documentation](https://docs.docker.com/engine/api/v1.41/#operation/ContainerList).

#### [Returns](#returns)

`Promise`&lt;`unknown`&gt;

* * *

### [listImages](#listimages)

▸ **listImages**(`options?`): `Promise`&lt;`unknown`&gt;

Get the list of local container images

```typescript
const images = await ddClient.docker.listImages();
```

#### [Parameters](#parameters-1)

NameTypeDescription`options?``any`(Optional). A JSON like `{ "all": true, "filters": JSON.stringify({ dangling: ["true"] }), "digests": true * }` For more information about the different properties see [the Docker API endpoint documentation](https://docs.docker.com/engine/api/v1.41/#tag/Image).

#### [Returns](#returns-1)

`Promise`&lt;`unknown`&gt;

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/Docker.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FDocker%2F&labels=status%2Ftriage)

Table of contents
