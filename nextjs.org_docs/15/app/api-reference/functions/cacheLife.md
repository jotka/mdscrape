---
title: "cacheLife"
source: "https://nextjs.org/docs/15/app/api-reference/functions/cacheLife"
---

# cacheLife

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[Functions](/docs/15/app/api-reference/functions)cacheLife

You are currently viewing documentation for version 15 of Next.js.

# cacheLife

This feature is currently available in the canary channel and subject to change. Try it out by [upgrading Next.js](/docs/app/getting-started/upgrading#canary-version), and share your feedback on [GitHub](https://github.com/vercel/next.js/issues).

The `cacheLife` function is used to set the cache lifetime of a function or component. It should be used alongside the [`use cache`](/docs/15/app/api-reference/directives/use-cache) directive, and within the scope of the function or component.

## [Usage](#usage)

To use `cacheLife`, enable the [`cacheComponents` flag](/docs/15/app/api-reference/config/next-config-js/cacheComponents) in your `next.config.js` file:

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

Then, import and invoke the `cacheLife` function within the scope of the function or component:

app/page.tsx

TypeScript

JavaScriptTypeScript

```
'use cache'
import { unstable_cacheLife as cacheLife } from 'next/cache'
 
export default async function Page() {
  cacheLife('hours')
  return <div>Page</div>
}
```

## [Reference](#reference)

### [Default cache profiles](#default-cache-profiles)

Next.js provides a set of named cache profiles modeled on various timescales. If you don't specify a cache profile in the `cacheLife` function alongside the `use cache` directive, Next.js will automatically apply the `default` cache profile.

However, we recommend always adding a cache profile when using the `use cache` directive to explicitly define caching behavior.

**Profile**`stale``revalidate``expire`**Description**`default`5 minutes15 minutes1 yearDefault profile, suitable for content that doesn't need frequent updates`seconds`01 second1 secondFor rapidly changing content requiring near real-time updates`minutes`5 minutes1 minute1 hourFor content that updates frequently within an hour`hours`5 minutes1 hour1 dayFor content that updates daily but can be slightly stale`days`5 minutes1 day1 weekFor content that updates weekly but can be a day old`weeks`5 minutes1 week30 daysFor content that updates monthly but can be a week old`max`5 minutes30 days1 yearFor very stable content that rarely needs updating

The string values used to reference cache profiles don't carry inherent meaning; instead they serve as semantic labels. This allows you to better understand and manage your cached content within your codebase.

> **Good to know:** Updating the [`staleTimes`](/docs/15/app/api-reference/config/next-config-js/staleTimes) and [`expireTime`](/docs/15/app/api-reference/config/next-config-js/expireTime) config options also updates the `stale` and `expire` properties of the `default` cache profile.

### [Custom cache profiles](#custom-cache-profiles)

You can configure custom cache profiles by adding them to the [`cacheLife`](/docs/15/app/api-reference/config/next-config-js/cacheLife) option in your `next.config.ts` file.

Cache profiles are objects that contain the following properties:

**Property****Value****Description****Requirement**`stale``number`Duration the client should cache a value without checking the server.Optional`revalidate``number`Frequency at which the cache should refresh on the server; stale values may be served while revalidating.Optional`expire``number`Maximum duration for which a value can remain stale before switching to dynamic fetching; must be longer than `revalidate`.Optional - Must be longer than `revalidate`

The "stale" property differs from the [`staleTimes`](/docs/15/app/api-reference/config/next-config-js/staleTimes) setting in that it specifically controls client-side router caching. While `staleTimes` is a global setting that affects all instances of both dynamic and static data, the `cacheLife` configuration allows you to define "stale" times on a per-function or per-route basis.

> **Good to know**: The “stale” property does not set the `Cache-control: max-age` header. It instead controls the client-side router cache.

## [Examples](#examples)

### [Defining reusable cache profiles](#defining-reusable-cache-profiles)

You can create a reusable cache profile by defining them in your `next.config.ts` file. Choose a name that suits your use case and set values for the `stale`, `revalidate`, and `expire` properties. You can create as many custom cache profiles as needed. Each profile can be referenced by its name as a string value passed to the `cacheLife` function.

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  experimental: {
    cacheComponents: true,
    cacheLife: {
      biweekly: {
        stale: 60 * 60 * 24 * 14, // 14 days
        revalidate: 60 * 60 * 24, // 1 day
        expire: 60 * 60 * 24 * 14, // 14 days
      },
    },
  },
}
 
module.exports = nextConfig
```

The example above caches for 14 days, checks for updates daily, and expires the cache after 14 days. You can then reference this profile throughout your application by its name:

app/page.tsx

```
'use cache'
import { unstable_cacheLife as cacheLife } from 'next/cache'
 
export default async function Page() {
  cacheLife('biweekly')
  return <div>Page</div>
}
```

### [Overriding the default cache profiles](#overriding-the-default-cache-profiles)

While the default cache profiles provide a useful way to think about how fresh or stale any given part of cacheable output can be, you may prefer different named profiles to better align with your applications caching strategies.

You can override the default named cache profiles by creating a new configuration with the same name as the defaults.

The example below shows how to override the default “days” cache profile:

next.config.ts

```
const nextConfig = {
  experimental: {
    cacheComponents: true,
    cacheLife: {
      days: {
        stale: 3600, // 1 hour
        revalidate: 900, // 15 minutes
        expire: 86400, // 1 day
      },
    },
  },
}
 
module.exports = nextConfig
```

### [Defining cache profiles inline](#defining-cache-profiles-inline)

For specific use cases, you can set a custom cache profile by passing an object to the `cacheLife` function:

app/page.tsx

TypeScript

JavaScriptTypeScript

```
'use cache'
import { unstable_cacheLife as cacheLife } from 'next/cache'
 
export default async function Page() {
  cacheLife({
    stale: 3600, // 1 hour
    revalidate: 900, // 15 minutes
    expire: 86400, // 1 day
  })
 
  return <div>Page</div>
}
```

This inline cache profile will only be applied to the function or file it was created in. If you want to reuse the same profile throughout your application, you can [add the configuration](#defining-reusable-cache-profiles) to the `cacheLife` property of your `next.config.ts` file.

### [Nested usage of `use cache` and `cacheLife`](#nested-usage-of-use-cache-and-cachelife)

When defining multiple caching behaviors in the same route or component tree, if the inner caches specify their own `cacheLife` profile, the outer cache will respect the shortest cache duration among them. **This applies only if the outer cache does not have its own explicit `cacheLife` profile defined.**

For example, if you add the `use cache` directive to your page, without specifying a cache profile, the default cache profile will be applied implicitly (`cacheLife(”default”)`). If a component imported into the page also uses the `use cache` directive with its own cache profile, the outer and inner cache profiles are compared, and shortest duration set in the profiles will be applied.

app/components/parent.tsx

```
// Parent component
import { unstable_cacheLife as cacheLife } from 'next/cache'
import { ChildComponent } from './child'
 
export async function ParentComponent() {
  'use cache'
  cacheLife('days')
 
  return (
    <div>
      <ChildComponent />
    </div>
  )
}
```

And in a separate file, we defined the Child component that was imported:

app/components/child.tsx

```
// Child component
import { unstable_cacheLife as cacheLife } from 'next/cache'
 
export async function ChildComponent() {
  'use cache'
  cacheLife('hours')
  return <div>Child Content</div>
 
  // This component's cache will respect the shorter 'hours' profile
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
**cacheTag**  
\
Learn how to use the cacheTag function to manage cache invalidation in your Next.js application.](/docs/15/app/api-reference/functions/cacheTag)

Was this helpful?

supported.

Send
