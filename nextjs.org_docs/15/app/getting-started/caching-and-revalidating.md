---
title: "Caching and Revalidating"
source: "https://nextjs.org/docs/15/app/getting-started/caching-and-revalidating"
---

# Caching and Revalidating

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[App Router](/docs/15/app)[Getting Started](/docs/15/app/getting-started)Caching and Revalidating

You are currently viewing documentation for version 15 of Next.js.

# Caching and Revalidating

Caching is a technique for storing the result of data fetching and other computations so that future requests for the same data can be served faster, without doing the work again. While revalidation allows you to update cache entries without having to rebuild your entire application.

Next.js provides a few APIs to handle caching and revalidation. This guide will walk you through when and how to use them.

- [`fetch`](#fetch)
- [`unstable_cache`](#unstable_cache)
- [`revalidatePath`](#revalidatepath)
- [`revalidateTag`](#revalidatetag)

## [`fetch`](#fetch)

By default, [`fetch`](/docs/15/app/api-reference/functions/fetch) requests are not cached. You can cache individual requests by setting the `cache` option to `'force-cache'`.

app/page.tsx

TypeScript

JavaScriptTypeScript

```
export default async function Page() {
  const data = await fetch('https://...', { cache: 'force-cache' })
}
```

> **Good to know**: Although `fetch` requests are not cached by default, Next.js will [prerender](/docs/15/app/getting-started/partial-prerendering#static-rendering) routes that have `fetch` requests and cache the HTML. If you want to guarantee a route is [dynamic](/docs/15/app/getting-started/partial-prerendering#dynamic-rendering), use the [`connection` API](/docs/15/app/api-reference/functions/connection).

To revalidate the data returned by a `fetch` request, you can use the `next.revalidate` option.

app/page.tsx

TypeScript

JavaScriptTypeScript

```
export default async function Page() {
  const data = await fetch('https://...', { next: { revalidate: 3600 } })
}
```

This will revalidate the data after a specified amount of seconds.

See the [`fetch` API reference](/docs/15/app/api-reference/functions/fetch) to learn more.

## [`unstable_cache`](#unstable_cache)

`unstable_cache` allows you to cache the result of database queries and other async functions. To use it, wrap `unstable_cache` around the function. For example:

```
import { db } from '@/lib/db'
export async function getUserById(id: string) {
  return db
    .select()
    .from(users)
    .where(eq(users.id, id))
    .then((res) => res[0])
}
```

app/page.tsx

TypeScript

JavaScriptTypeScript

```
import { unstable_cache } from 'next/cache'
import { getUserById } from '@/app/lib/data'
 
export default async function Page({
  params,
}: {
  params: Promise<{ userId: string }>
}) {
  const { userId } = await params
 
  const getCachedUser = unstable_cache(
    async () => {
      return getUserById(userId)
    },
    [userId] // add the user ID to the cache key
  )
}
```

The function accepts a third optional object to define how the cache should be revalidated. It accepts:

- `tags`: an array of tags used by Next.js to revalidate the cache.
- `revalidate`: the number of seconds after cache should be revalidated.

app/page.tsx

TypeScript

JavaScriptTypeScript

```
const getCachedUser = unstable_cache(
  async () => {
    return getUserById(userId)
  },
  [userId],
  {
    tags: ['user'],
    revalidate: 3600,
  }
)
```

See the [`unstable_cache` API reference](/docs/15/app/api-reference/functions/unstable_cache) to learn more.

## [`revalidateTag`](#revalidatetag)

`revalidateTag` is used to revalidate cache entries based on a tag and following an event. To use it with `fetch`, start by tagging the function with the `next.tags` option:

app/lib/data.ts

TypeScript

JavaScriptTypeScript

```
export async function getUserById(id: string) {
  const data = await fetch(`https://...`, {
    next: {
      tags: ['user'],
    },
  })
}
```

Alternatively, you can mark an `unstable_cache` function with the `tags` option:

app/lib/data.ts

TypeScript

JavaScriptTypeScript

```
export const getUserById = unstable_cache(
  async (id: string) => {
    return db.query.users.findFirst({ where: eq(users.id, id) })
  },
  ['user'], // Needed if variables are not passed as parameters
  {
    tags: ['user'],
  }
)
```

Then, call `revalidateTag` in a [Route Handler](/docs/15/app/api-reference/file-conventions/route) or Server Action:

app/lib/actions.ts

TypeScript

JavaScriptTypeScript

```
import { revalidateTag } from 'next/cache'
 
export async function updateUser(id: string) {
  // Mutate data
  revalidateTag('user')
}
```

You can reuse the same tag in multiple functions to revalidate them all at once.

See the [`revalidateTag` API reference](/docs/15/app/api-reference/functions/revalidateTag) to learn more.

## [`revalidatePath`](#revalidatepath)

`revalidatePath` is used to revalidate a route and following an event. To use it, call it in a [Route Handler](/docs/15/app/api-reference/file-conventions/route) or Server Action:

app/lib/actions.ts

TypeScript

JavaScriptTypeScript

```
import { revalidatePath } from 'next/cache'
 
export async function updateUser(id: string) {
  // Mutate data
  revalidatePath('/profile')
```

See the [`revalidatePath` API reference](/docs/15/app/api-reference/functions/revalidatePath) to learn more.

## API Reference

Learn more about the features mentioned in this page by reading the API Reference.

[Next.js Docs  
\
...  
\
Functions  
\
**fetch**  
\
API reference for the extended fetch function.](/docs/15/app/api-reference/functions/fetch)

[Next.js Docs  
\
...  
\
Functions  
\
**unstable\_cache**  
\
API Reference for the unstable\_cache function.](/docs/15/app/api-reference/functions/unstable_cache)

[Next.js Docs  
\
...  
\
Functions  
\
**revalidatePath**  
\
API Reference for the revalidatePath function.](/docs/15/app/api-reference/functions/revalidatePath)

[Next.js Docs  
\
...  
\
Functions  
\
**revalidateTag**  
\
API Reference for the revalidateTag function.](/docs/15/app/api-reference/functions/revalidateTag)

Was this helpful?

supported.

Send
