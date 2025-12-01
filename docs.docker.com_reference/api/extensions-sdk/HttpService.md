---
title: "Interface: HttpService"
source: "https://docs.docker.com/reference/api/extensions-sdk/HttpService/"
---

# Interface: HttpService

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: HttpService

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

## [Methods](#methods)

### [get](#get)

▸ **get**(`url`): `Promise`&lt;`unknown`&gt;

Performs an HTTP GET request to a backend service.

```typescript
ddClient.extension.vm.service
 .get("/some/service")
 .then((value: any) => console.log(value)
```

#### [Parameters](#parameters)

NameTypeDescription`url``string`The URL of the backend service.

#### [Returns](#returns)

`Promise`&lt;`unknown`&gt;

* * *

### [post](#post)

▸ **post**(`url`, `data`): `Promise`&lt;`unknown`&gt;

Performs an HTTP POST request to a backend service.

```typescript
ddClient.extension.vm.service
 .post("/some/service", { ... })
 .then((value: any) => console.log(value));
```

#### [Parameters](#parameters-1)

NameTypeDescription`url``string`The URL of the backend service.`data``any`The body of the request.

#### [Returns](#returns-1)

`Promise`&lt;`unknown`&gt;

* * *

### [put](#put)

▸ **put**(`url`, `data`): `Promise`&lt;`unknown`&gt;

Performs an HTTP PUT request to a backend service.

```typescript
ddClient.extension.vm.service
 .put("/some/service", { ... })
 .then((value: any) => console.log(value));
```

#### [Parameters](#parameters-2)

NameTypeDescription`url``string`The URL of the backend service.`data``any`The body of the request.

#### [Returns](#returns-2)

`Promise`&lt;`unknown`&gt;

* * *

### [patch](#patch)

▸ **patch**(`url`, `data`): `Promise`&lt;`unknown`&gt;

Performs an HTTP PATCH request to a backend service.

```typescript
ddClient.extension.vm.service
 .patch("/some/service", { ... })
 .then((value: any) => console.log(value));
```

#### [Parameters](#parameters-3)

NameTypeDescription`url``string`The URL of the backend service.`data``any`The body of the request.

#### [Returns](#returns-3)

`Promise`&lt;`unknown`&gt;

* * *

### [delete](#delete)

▸ **delete**(`url`): `Promise`&lt;`unknown`&gt;

Performs an HTTP DELETE request to a backend service.

```typescript
ddClient.extension.vm.service
 .delete("/some/service")
 .then((value: any) => console.log(value));
```

#### [Parameters](#parameters-4)

NameTypeDescription`url``string`The URL of the backend service.

#### [Returns](#returns-4)

`Promise`&lt;`unknown`&gt;

* * *

### [head](#head)

▸ **head**(`url`): `Promise`&lt;`unknown`&gt;

Performs an HTTP HEAD request to a backend service.

```typescript
ddClient.extension.vm.service
 .head("/some/service")
 .then((value: any) => console.log(value));
```

#### [Parameters](#parameters-5)

NameTypeDescription`url``string`The URL of the backend service.

#### [Returns](#returns-5)

`Promise`&lt;`unknown`&gt;

* * *

### [request](#request)

▸ **request**(`config`): `Promise`&lt;`unknown`&gt;

Performs an HTTP request to a backend service.

```typescript
ddClient.extension.vm.service
 .request({ url: "/url", method: "GET", headers: { 'header-key': 'header-value' }, data: { ... }})
 .then((value: any) => console.log(value));
```

#### [Parameters](#parameters-6)

NameTypeDescription`config`[`RequestConfig`](https://docs.docker.com/reference/api/extensions-sdk/RequestConfig/)The URL of the backend service.

#### [Returns](#returns-6)

`Promise`&lt;`unknown`&gt;

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/HttpService.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FHttpService%2F&labels=status%2Ftriage)

Table of contents
