---
title: "Interface: ExtensionVM"
source: "https://docs.docker.com/reference/api/extensions-sdk/ExtensionVM/"
---

# Interface: ExtensionVM

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: ExtensionVM

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

• `Readonly` **cli**: [`ExtensionCli`](https://docs.docker.com/reference/api/extensions-sdk/ExtensionCli/)

Executes a command in the backend container.

Example: Execute the command `ls -l` inside the backend container:

```typescript
await ddClient.extension.vm.cli.exec(
  "ls",
  ["-l"]
);
```

Streams the output of the command executed in the backend container.

When the extension defines its own `compose.yaml` file with multiple containers, the command is executed on the first container defined. Change the order in which containers are defined to execute commands on another container.

Example: Spawn the command `ls -l` inside the backend container:

```typescript
await ddClient.extension.vm.cli.exec("ls", ["-l"], {
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

**`Param`**

Command to execute.

**`Param`**

Arguments of the command to execute.

**`Param`**

The callback function where to listen from the command output data and errors.

* * *

### [service](#service)

• `Optional` `Readonly` **service**: [`HttpService`](https://docs.docker.com/reference/api/extensions-sdk/HttpService/)

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/ExtensionVM.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FExtensionVM%2F&labels=status%2Ftriage)

Table of contents
