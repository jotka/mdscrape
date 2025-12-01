---
title: "Interface: NavigationIntents"
source: "https://docs.docker.com/reference/api/extensions-sdk/NavigationIntents/"
---

# Interface: NavigationIntents

Back

Ask AI

Start typing to search or try Ask AI.

[Contact support](https://hub.docker.com/support/contact)

[Reference](https://docs.docker.com/reference/)

- [Get started](/get-started/)
- [Guides](/guides/)
- [Manuals](/manuals/)

# Interface: NavigationIntents

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

## [Container Methods](#container-methods)

### [viewContainers](#viewcontainers)

▸ **viewContainers**(): `Promise`&lt;`void`&gt;

Navigate to the **Containers** tab in Docker Desktop.

```typescript
ddClient.desktopUI.navigate.viewContainers()
```

#### [Returns](#returns)

`Promise`&lt;`void`&gt;

* * *

### [viewContainer](#viewcontainer)

▸ **viewContainer**(`id`): `Promise`&lt;`void`&gt;

Navigate to the **Container** tab in Docker Desktop.

```typescript
await ddClient.desktopUI.navigate.viewContainer(id)
```

#### [Parameters](#parameters)

NameTypeDescription`id``string`The full container id, e.g. `46b57e400d801762e9e115734bf902a2450d89669d85881058a46136520aca28`. You can use the `--no-trunc` flag as part of the `docker ps` command to display the full container id.

#### [Returns](#returns-1)

`Promise`&lt;`void`&gt;

A promise that fails if the container doesn't exist.

* * *

### [viewContainerLogs](#viewcontainerlogs)

▸ **viewContainerLogs**(`id`): `Promise`&lt;`void`&gt;

Navigate to the **Container logs** tab in Docker Desktop.

```typescript
await ddClient.desktopUI.navigate.viewContainerLogs(id)
```

#### [Parameters](#parameters-1)

NameTypeDescription`id``string`The full container id, e.g. `46b57e400d801762e9e115734bf902a2450d89669d85881058a46136520aca28`. You can use the `--no-trunc` flag as part of the `docker ps` command to display the full container id.

#### [Returns](#returns-2)

`Promise`&lt;`void`&gt;

A promise that fails if the container doesn't exist.

* * *

### [viewContainerInspect](#viewcontainerinspect)

▸ **viewContainerInspect**(`id`): `Promise`&lt;`void`&gt;

Navigate to the **Inspect container** view in Docker Desktop.

```typescript
await ddClient.desktopUI.navigate.viewContainerInspect(id)
```

#### [Parameters](#parameters-2)

NameTypeDescription`id``string`The full container id, e.g. `46b57e400d801762e9e115734bf902a2450d89669d85881058a46136520aca28`. You can use the `--no-trunc` flag as part of the `docker ps` command to display the full container id.

#### [Returns](#returns-3)

`Promise`&lt;`void`&gt;

A promise that fails if the container doesn't exist.

* * *

### [viewContainerTerminal](#viewcontainerterminal)

▸ **viewContainerTerminal**(`id`): `Promise`&lt;`void`&gt;

Navigate to the container terminal window in Docker Desktop.

```typescript
await ddClient.desktopUI.navigate.viewContainerTerminal(id)
```

**`Since`**

0.3.4

#### [Parameters](#parameters-3)

NameTypeDescription`id``string`The full container id, e.g. `46b57e400d801762e9e115734bf902a2450d89669d85881058a46136520aca28`. You can use the `--no-trunc` flag as part of the `docker ps` command to display the full container id.

#### [Returns](#returns-4)

`Promise`&lt;`void`&gt;

A promise that fails if the container doesn't exist.

* * *

### [viewContainerStats](#viewcontainerstats)

▸ **viewContainerStats**(`id`): `Promise`&lt;`void`&gt;

Navigate to the container stats to see the CPU, memory, disk read/write and network I/O usage.

```typescript
await ddClient.desktopUI.navigate.viewContainerStats(id)
```

#### [Parameters](#parameters-4)

NameTypeDescription`id``string`The full container id, e.g. `46b57e400d801762e9e115734bf902a2450d89669d85881058a46136520aca28`. You can use the `--no-trunc` flag as part of the `docker ps` command to display the full container id.

#### [Returns](#returns-5)

`Promise`&lt;`void`&gt;

A promise that fails if the container doesn't exist.

* * *

## [Images Methods](#images-methods)

### [viewImages](#viewimages)

▸ **viewImages**(): `Promise`&lt;`void`&gt;

Navigate to the **Images** tab in Docker Desktop.

```typescript
await ddClient.desktopUI.navigate.viewImages()
```

#### [Returns](#returns-6)

`Promise`&lt;`void`&gt;

* * *

### [viewImage](#viewimage)

▸ **viewImage**(`id`, `tag`): `Promise`&lt;`void`&gt;

Navigate to a specific image referenced by `id` and `tag` in Docker Desktop. In this navigation route you can find the image layers, commands, created time and size.

```typescript
await ddClient.desktopUI.navigate.viewImage(id, tag)
```

#### [Parameters](#parameters-5)

NameTypeDescription`id``string`The full image id (including sha), e.g. `sha256:34ab3ae068572f4e85c448b4035e6be5e19cc41f69606535cd4d768a63432673`.`tag``string`The tag of the image, e.g. `latest`, `0.0.1`, etc.

#### [Returns](#returns-7)

`Promise`&lt;`void`&gt;

A promise that fails if the image doesn't exist.

* * *

## [Volume Methods](#volume-methods)

### [viewVolumes](#viewvolumes)

▸ **viewVolumes**(): `Promise`&lt;`void`&gt;

Navigate to the **Volumes** tab in Docker Desktop.

```typescript
ddClient.desktopUI.navigate.viewVolumes()
```

#### [Returns](#returns-8)

`Promise`&lt;`void`&gt;

* * *

### [viewVolume](#viewvolume)

▸ **viewVolume**(`volume`): `Promise`&lt;`void`&gt;

Navigate to a specific volume in Docker Desktop.

```typescript
await ddClient.desktopUI.navigate.viewVolume(volume)
```

#### [Parameters](#parameters-6)

NameTypeDescription`volume``string`The name of the volume, e.g. `my-volume`.

#### [Returns](#returns-9)

`Promise`&lt;`void`&gt;

[Edit this page](https://github.com/docker/docs/edit/main/content/reference/api/extensions-sdk/NavigationIntents.md)

[Request changes](https://github.com/docker/docs/issues/new?template=doc_issue.yml&location=https%3A%2F%2Fdocs.docker.com%2Freference%2Fapi%2Fextensions-sdk%2FNavigationIntents%2F&labels=status%2Ftriage)

Table of contents
