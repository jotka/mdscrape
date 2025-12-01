---
title: "use cache: private"
source: "https://nextjs.org/docs/app/api-reference/directives/use-cache-private"
---

# use cache: private

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[API Reference](/docs/app/api-reference)[Directives](/docs/app/api-reference/directives)use cache: private

Copy page

# use cache: private

The `'use cache: private'` directive works just like [`use cache`](/docs/app/api-reference/directives/use-cache), but allows you to use runtime APIs like cookies, headers, or search params.

> **Good to know:** Unlike `use cache`, private caches are not prerendered statically as they contain personalized data that is not shared between users.

## [Usage](#usage)

To use `'use cache: private'`, enable the [`cacheComponents`](/docs/app/api-reference/config/next-config-js/cacheComponents) flag in your `next.config.ts` file:

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  cacheComponents: true,
}
 
export default nextConfig
```

Then add `'use cache: private'` to your function along with a `cacheLife` configuration.

### [Basic example](#basic-example)

app/product/\[id]/page.tsx

TypeScript

JavaScriptTypeScript

```
import { Suspense } from 'react'
import { cookies } from 'next/headers'
import { cacheLife, cacheTag } from 'next/cache'
 
export default async function ProductPage({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params
 
  return (
    <div>
      <ProductDetails id={id} />
      <Suspense fallback={<div>Loading recommendations...</div>}>
        <Recommendations productId={id} />
      </Suspense>
    </div>
  )
}
 
async function Recommendations({ productId }: { productId: string }) {
  const recommendations = await getRecommendations(productId)
 
  return (
    <div>
      {recommendations.map((rec) => (
        <ProductCard key={rec.id} product={rec} />
      ))}
    </div>
  )
}
 
async function getRecommendations(productId: string) {
  'use cache: private'
  cacheTag(`recommendations-${productId}`)
  cacheLife({ stale: 60 }) // Minimum 30 seconds required for runtime prefetch
 
  // Access cookies within private cache functions
  const sessionId = (await cookies()).get('session-id')?.value || 'guest'
 
  return getPersonalizedRecommendations(productId, sessionId)
}
```

## [Request APIs allowed in private caches](#request-apis-allowed-in-private-caches)

The following request-specific APIs can be used inside `'use cache: private'` functions:

APIAllowed in `use cache`Allowed in `'use cache: private'``cookies()`NoYes`headers()`NoYes`searchParams`NoYes`connection()`NoNo

> **Note:** The [`connection()`](https://nextjs.org/docs/app/api-reference/functions/connection) API is prohibited in both `use cache` and `'use cache: private'` as it provides connection-specific information that cannot be safely cached.

## [Version History](#version-history)

VersionChanges`v16.0.0``"use cache: private"` is enabled with the Cache Components feature.

## Related

View related API references.

[**use cache**  
\
Learn how to use the use cache directive to cache data in your Next.js application.](/docs/app/api-reference/directives/use-cache)

[**cacheComponents**  
\
Learn how to enable the cacheComponents flag in Next.js.](/docs/app/api-reference/config/next-config-js/cacheComponents)

[**cacheHandlers**  
\
Configure custom cache handlers for use cache directives in Next.js.](/docs/app/api-reference/config/next-config-js/cacheHandlers)

[**cacheLife**  
\
Learn how to use the cacheLife function to set the cache expiration time for a cached function or component.](/docs/app/api-reference/functions/cacheLife)

[**cacheTag**  
\
Learn how to use the cacheTag function to manage cache invalidation in your Next.js application.](/docs/app/api-reference/functions/cacheTag)

[**Prefetching**  
\
Learn how to configure prefetching in Next.js](/docs/app/guides/prefetching)

Was this helpful?

supported.

Send
