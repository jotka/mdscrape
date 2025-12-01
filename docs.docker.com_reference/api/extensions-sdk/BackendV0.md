---
title: "Interface: BackendV0"
source: "https://docs.docker.com/reference/api/extensions-sdk/BackendV0/"
---

# Interface: BackendV0

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: BackendV0

Page options

Copy page as Markdown for LLMs

View page as plain text

Ask questions with Docs AI

Claude

Open in Claude

Table of contents

* * *

## [Container Methods](#container-methods)

### [execInContainer](#execincontainer)

▸ **execInContainer**(`container`, `cmd`): `Promise`&lt;[`ExecResultV0`](https://docs.docker.com/reference/api/extensions-sdk/ExecResultV0/)&gt;

Executes a command inside a container.

```typescript
const output = await window.ddClient.backend.execInContainer(container, cmd);

console.log(output);
```

> Warning
> 
> It will be removed in a future version.

#### [Parameters](#parameters)

NameTypeDescription`container``string`-`cmd``string`The command to be executed.

#### [Returns](#returns)

`Promise`&lt;[`ExecResultV0`](https://docs.docker.com/reference/api/extensions-sdk/ExecResultV0/)&gt;

* * *

## [HTTP Methods](#http-methods)

### [get](#get)

▸ **get**(`url`): `Promise`&lt;`unknown`&gt;

Performs an HTTP GET request to a backend service.

```typescript
window.ddClient.backend
 .get("/some/service")
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [get](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#get) instead.

#### [Parameters](#parameters-1)

NameTypeDescription`url``string`The URL of the backend service.

#### [Returns](#returns-1)

`Promise`&lt;`unknown`&gt;

* * *

### [post](#post)

▸ **post**(`url`, `data`): `Promise`&lt;`unknown`&gt;

Performs an HTTP POST request to a backend service.

```typescript
window.ddClient.backend
 .post("/some/service", { ... })
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [post](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#post) instead.

#### [Parameters](#parameters-2)

NameTypeDescription`url``string`The URL of the backend service.`data``any`The body of the request.

#### [Returns](#returns-2)

`Promise`&lt;`unknown`&gt;

* * *

### [put](#put)

▸ **put**(`url`, `data`): `Promise`&lt;`unknown`&gt;

Performs an HTTP PUT request to a backend service.

```typescript
window.ddClient.backend
 .put("/some/service", { ... })
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [put](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#put) instead.

#### [Parameters](#parameters-3)

NameTypeDescription`url``string`The URL of the backend service.`data``any`The body of the request.

#### [Returns](#returns-3)

`Promise`&lt;`unknown`&gt;

* * *

### [patch](#patch)

▸ **patch**(`url`, `data`): `Promise`&lt;`unknown`&gt;

Performs an HTTP PATCH request to a backend service.

```typescript
window.ddClient.backend
 .patch("/some/service", { ... })
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [patch](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#patch) instead.

#### [Parameters](#parameters-4)

NameTypeDescription`url``string`The URL of the backend service.`data``any`The body of the request.

#### [Returns](#returns-4)

`Promise`&lt;`unknown`&gt;

* * *

### [delete](#delete)

▸ **delete**(`url`): `Promise`&lt;`unknown`&gt;

Performs an HTTP DELETE request to a backend service.

```typescript
window.ddClient.backend
 .delete("/some/service")
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [delete](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#delete) instead.

#### [Parameters](#parameters-5)

NameTypeDescription`url``string`The URL of the backend service.

#### [Returns](#returns-5)

`Promise`&lt;`unknown`&gt;

* * *

### [head](#head)

▸ **head**(`url`): `Promise`&lt;`unknown`&gt;

Performs an HTTP HEAD request to a backend service.

```typescript
window.ddClient.backend
 .head("/some/service")
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [head](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#head) instead.

#### [Parameters](#parameters-6)

NameTypeDescription`url``string`The URL of the backend service.

#### [Returns](#returns-6)

`Promise`&lt;`unknown`&gt;

* * *

### [request](#request)

▸ **request**(`config`): `Promise`&lt;`unknown`&gt;

Performs an HTTP request to a backend service.

```typescript
window.ddClient.backend
 .request({ url: "/url", method: "GET", headers: { 'header-key': 'header-value' }, data: { ... }})
 .then((value: any) => console.log(value));
```

> Warning
> 
> It will be removed in a future version. Use [request](https://docs.docker.com/reference/api/extensions-sdk/HttpService/#request) instead.

#### [Parameters](#parameters-7)

NameTypeDescription`config`[`RequestConfigV0`](https://docs.docker.com/reference/api/extensions-sdk/RequestConfigV0/)The URL of the backend service.

#### [Returns](#returns-7)

`Promise`&lt;`unknown`&gt;

* * *

## [VM Methods](#vm-methods)

### [execInVMExtension](#execinvmextension)

▸ **execInVMExtension**(`cmd`): `Promise`&lt;[`ExecResultV0`](https://docs.docker.com/reference/api/extensions-sdk/ExecResultV0/)&gt;

Executes a command inside the backend container. If your extensions ships with additional binaries that should be run inside the backend container you can use the `execInVMExtension` function.

```typescript
const output = await window.ddClient.backend.execInVMExtension(
  `cliShippedInTheVm xxx`
);

console.log(output);
```

> Warning
> 
> It will be removed in a future version. Use [exec](https://docs.docker.com/reference/api/extensions-sdk/ExtensionCli/#exec) instead.

#### [Parameters](#parameters-8)

NameTypeDescription`cmd``string`The command to be executed.

#### [Returns](#returns-8)

`Promise`&lt;[`ExecResultV0`](https://docs.docker.com/reference/api/extensions-sdk/ExecResultV0/)&gt;

* * *

### [spawnInVMExtension](#spawninvmextension)

▸ **spawnInVMExtension**(`cmd`, `args`, `callback`): `void`

Returns a stream from the command executed in the backend container.

```typescript
window.ddClient.spawnInVMExtension(
  `cmd`,
  [`arg1`, `arg2`],
  (data: any, err: any) => {
    console.log(data.stdout, data.stderr);
    // Once the command exits we get the status code
    if (data.code) {
      console.log(data.code);
    }
  }
);
```

> Warning
> 
> It will be removed in a future version.

#### [Parameters](#parameters-9)

NameTypeDescription`cmd``string`The command to be executed.`args``string`\[]The arguments of the command to execute.`callback`(`data`: `any`, `error`: `any`) =&gt; `void`The callback function where to listen from the command output data and errors.

#### [Returns](#returns-9)

`void`

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/BackendV0.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FBackendV0%2F&labels=status%2Ftriage)

Table of contents
