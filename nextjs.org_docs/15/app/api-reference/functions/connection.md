---
title: "connection"
source: "https://nextjs.org/docs/15/app/api-reference/functions/connection"
---

# connection

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)connection

You are currently viewing documentation for version 15 of Next.js.

# connection

The `connection()` function allows you to indicate rendering should wait for an incoming user request before continuing.

It's useful when a component doesn’t use [Dynamic APIs](/docs/15/app/getting-started/partial-prerendering#dynamic-rendering), but you want it to be dynamically rendered at runtime and not statically rendered at build time. This usually occurs when you access external information that you intentionally want to change the result of a render, such as `Math.random()` or `new Date()`.

app/page.tsx

TypeScript

JavaScriptTypeScript

```
import { connection } from 'next/server'
 
export default async function Page() {
  await connection()
  // Everything below will be excluded from prerendering
  const rand = Math.random()
  return <span>{rand}</span>
}
```

## [Reference](#reference)

### [Type](#type)

```
function connection(): Promise<void>
```

### [Parameters](#parameters)

- The function does not accept any parameters.

### [Returns](#returns)

- The function returns a `void` Promise. It is not meant to be consumed.

## [Good to know](#good-to-know)

- `connection` replaces [`unstable_noStore`](/docs/15/app/api-reference/functions/unstable_noStore) to better align with the future of Next.js.
- The function is only necessary when dynamic rendering is required and common Dynamic APIs are not used.

### [Version History](#version-history)

VersionChanges`v15.0.0``connection` stabilized.`v15.0.0-RC``connection` introduced.

Was this helpful?

supported.

Send
