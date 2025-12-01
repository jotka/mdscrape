---
title: "cacheTag"
source: "https://nextjs.org/docs/15/app/api-reference/functions/cacheTag"
---

# cacheTag

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)cacheTag

You are currently viewing documentation for version 15 of Next.js.

# cacheTag

This feature is currently available in the canary channel and subject to change. Try it out by [upgrading Next.js](/docs/app/getting-started/upgrading#canary-version), and share your feedback on [GitHub](https://github.com/vercel/next.js/issues).

The `cacheTag` function allows you to tag cached data for on-demand invalidation. By associating tags with cache entries, you can selectively purge or revalidate specific cache entries without affecting other cached data.

## [Usage](#usage)

To use `cacheTag`, enable the [`cacheComponents` flag](/docs/15/app/api-reference/config/next-config-js/cacheComponents) in your `next.config.js` file:

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    cacheComponents: true,
  },
}
 
export default nextConfig
```

The `cacheTag` function takes one or more string values.

app/data.ts

TypeScript

JavaScriptTypeScript

```
import { unstable_cacheTag as cacheTag } from 'next/cache'
 
export async function getData() {
  'use cache'
  cacheTag('my-data')
  const data = await fetch('/api/data')
  return data
}
```

You can then purge the cache on-demand using [`revalidateTag`](/docs/15/app/api-reference/functions/revalidateTag) API in another function, for example, a [route handler](/docs/15/app/api-reference/file-conventions/route) or [Server Action](/docs/15/app/getting-started/updating-data):

app/action.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { revalidateTag } from 'next/cache'
 
export default async function submit() {
  await addPost()
  revalidateTag('my-data')
}
```

## [Good to know](#good-to-know)

- **Idempotent Tags**: Applying the same tag multiple times has no additional effect.
- **Multiple Tags**: You can assign multiple tags to a single cache entry by passing multiple string values to `cacheTag`.

```
cacheTag('tag-one', 'tag-two')
```

## [Examples](#examples)

### [Tagging components or functions](#tagging-components-or-functions)

Tag your cached data by calling `cacheTag` within a cached function or component:

app/components/bookings.tsx

TypeScript

JavaScriptTypeScript

```
import { unstable_cacheTag as cacheTag } from 'next/cache'
 
interface BookingsProps {
  type: string
}
 
export async function Bookings({ type = 'haircut' }: BookingsProps) {
  'use cache'
  cacheTag('bookings-data')
 
  async function getBookingsData() {
    const data = await fetch(`/api/bookings?type=${encodeURIComponent(type)}`)
    return data
  }
 
  return //...
}
```

### [Creating tags from external data](#creating-tags-from-external-data)

You can use the data returned from an async function to tag the cache entry.

app/components/bookings.tsx

TypeScript

JavaScriptTypeScript

```
import { unstable_cacheTag as cacheTag } from 'next/cache'
 
interface BookingsProps {
  type: string
}
 
export async function Bookings({ type = 'haircut' }: BookingsProps) {
  async function getBookingsData() {
    'use cache'
    const data = await fetch(`/api/bookings?type=${encodeURIComponent(type)}`)
    cacheTag('bookings-data', data.id)
    return data
  }
  return //...
}
```

### [Invalidating tagged cache](#invalidating-tagged-cache)

Using [`revalidateTag`](/docs/15/app/api-reference/functions/revalidateTag), you can invalidate the cache for a specific tag when needed:

app/actions.ts

TypeScript

JavaScriptTypeScript

```
'use server'
 
import { revalidateTag } from 'next/cache'
 
export async function updateBookings() {
  await updateBookingData()
  revalidateTag('bookings-data')
}
```

## Related

View related API references.

[Next.js Docs  
\
...  
\
next.config.js  
\
**cacheComponents**  
\
Learn how to enable the cacheComponents flag in Next.js.](/docs/15/app/api-reference/config/next-config-js/cacheComponents)

[Next.js Docs  
\
...  
\
Directives  
\
**use cache**  
\
Learn how to use the use cache directive to cache data in your Next.js application.](/docs/15/app/api-reference/directives/use-cache)

[Next.js Docs  
\
...  
\
Functions  
\
**revalidateTag**  
\
API Reference for the revalidateTag function.](/docs/15/app/api-reference/functions/revalidateTag)

[Next.js Docs  
\
...  
\
Functions  
\
**cacheLife**  
\
Learn how to use the cacheLife function to set the cache expiration time for a cached function or component.](/docs/15/app/api-reference/functions/cacheLife)

Was this helpful?

supported.

Send
