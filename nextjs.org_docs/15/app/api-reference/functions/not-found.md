---
title: "notFound"
source: "https://nextjs.org/docs/15/app/api-reference/functions/not-found"
---

# notFound

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)notFound

You are currently viewing documentation for version 15 of Next.js.

# notFound

The `notFound` function allows you to render the [`not-found file`](/docs/15/app/api-reference/file-conventions/not-found) within a route segment as well as inject a `<meta name="robots" content="noindex" />` tag.

## [`notFound()`](#notfound)

Invoking the `notFound()` function throws a `NEXT_HTTP_ERROR_FALLBACK;404` error and terminates rendering of the route segment in which it was thrown. Specifying a [**not-found** file](/docs/15/app/api-reference/file-conventions/not-found) allows you to gracefully handle such errors by rendering a Not Found UI within the segment.

app/user/\[id]/page.js

```
import { notFound } from 'next/navigation'
 
async function fetchUser(id) {
  const res = await fetch('https://...')
  if (!res.ok) return undefined
  return res.json()
}
 
export default async function Profile({ params }) {
  const { id } = await params
  const user = await fetchUser(id)
 
  if (!user) {
    notFound()
  }
 
  // ...
}
```

> **Good to know**: `notFound()` does not require you to use `return notFound()` due to using the TypeScript [`never`](https://www.typescriptlang.org/docs/handbook/2/functions.html#never) type.

## [Version History](#version-history)

VersionChanges`v13.0.0``notFound` introduced.

Was this helpful?

supported.

Send
