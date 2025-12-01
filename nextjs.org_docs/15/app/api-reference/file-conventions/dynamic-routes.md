---
title: "Dynamic Route Segments"
source: "https://nextjs.org/docs/15/app/api-reference/file-conventions/dynamic-routes"
---

# Dynamic Route Segments

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[File-system conventions](/docs/15/app/api-reference/file-conventions)Dynamic Segments

You are currently viewing documentation for version 15 of Next.js.

# Dynamic Route Segments

When you don't know the exact route segment names ahead of time and want to create routes from dynamic data, you can use Dynamic Segments that are filled in at request time or prerendered at build time.

## [Convention](#convention)

A Dynamic Segment can be created by wrapping a folder's name in square brackets: `[folderName]`. For example, a blog could include the following route `app/blog/[slug]/page.js` where `[slug]` is the Dynamic Segment for blog posts.

app/blog/\[slug]/page.tsx

TypeScript

JavaScriptTypeScript

```
export default async function Page({
  params,
}: {
  params: Promise<{ slug: string }>
}) {
  const { slug } = await params
  return <div>My Post: {slug}</div>
}
```

Dynamic Segments are passed as the `params` prop to [`layout`](/docs/15/app/api-reference/file-conventions/layout), [`page`](/docs/15/app/api-reference/file-conventions/page), [`route`](/docs/15/app/api-reference/file-conventions/route), and [`generateMetadata`](/docs/15/app/api-reference/functions/generate-metadata#generatemetadata-function) functions.

RouteExample URL`params``app/blog/[slug]/page.js``/blog/a``{ slug: 'a' }``app/blog/[slug]/page.js``/blog/b``{ slug: 'b' }``app/blog/[slug]/page.js``/blog/c``{ slug: 'c' }`

### [In Client Components](#in-client-components)

In a Client Component **page**, dynamic segments from props can be accessed using the [`use`](https://react.dev/reference/react/use) hook.

app/blog/\[slug]/page.tsx

TypeScript

JavaScriptTypeScript

```
'use client'
import { use } from 'react'
 
export default function BlogPostPage({
  params,
}: {
  params: Promise<{ slug: string }>
}) {
  const { slug } = use(params)
 
  return (
    <div>
      <p>{slug}</p>
    </div>
  )
}
```

Alternatively Client Components can use the [`useParams`](/docs/15/app/api-reference/functions/use-params) hook to access the `params` anywhere in the Client Component tree.

### [Catch-all Segments](#catch-all-segments)

Dynamic Segments can be extended to **catch-all** subsequent segments by adding an ellipsis inside the brackets `[...folderName]`.

For example, `app/shop/[...slug]/page.js` will match `/shop/clothes`, but also `/shop/clothes/tops`, `/shop/clothes/tops/t-shirts`, and so on.

RouteExample URL`params``app/shop/[...slug]/page.js``/shop/a``{ slug: ['a'] }``app/shop/[...slug]/page.js``/shop/a/b``{ slug: ['a', 'b'] }``app/shop/[...slug]/page.js``/shop/a/b/c``{ slug: ['a', 'b', 'c'] }`

### [Optional Catch-all Segments](#optional-catch-all-segments)

Catch-all Segments can be made **optional** by including the parameter in double square brackets: `[[...folderName]]`.

For example, `app/shop/[[...slug]]/page.js` will **also** match `/shop`, in addition to `/shop/clothes`, `/shop/clothes/tops`, `/shop/clothes/tops/t-shirts`.

The difference between **catch-all** and **optional catch-all** segments is that with optional, the route without the parameter is also matched (`/shop` in the example above).

RouteExample URL`params``app/shop/[[...slug]]/page.js``/shop``{ slug: undefined }``app/shop/[[...slug]]/page.js``/shop/a``{ slug: ['a'] }``app/shop/[[...slug]]/page.js``/shop/a/b``{ slug: ['a', 'b'] }``app/shop/[[...slug]]/page.js``/shop/a/b/c``{ slug: ['a', 'b', 'c'] }`

### [TypeScript](#typescript)

When using TypeScript, you can add types for `params` depending on your configured route segment — use [`PageProps<'/route'>`](/docs/15/app/api-reference/file-conventions/page#page-props-helper), [`LayoutProps<'/route'>`](/docs/15/app/api-reference/file-conventions/layout#layout-props-helper), or [`RouteContext<'/route'>`](/docs/15/app/api-reference/file-conventions/route#route-context-helper) to type `params` in `page`, `layout`, and `route` respectively.

Route `params` values are typed as `string`, `string[]`, or `undefined` (for optional catch-all segments), because their values aren't known until runtime. Users can enter any URL into the address bar, and these broad types help ensure that your application code handles all these possible cases.

Route`params` Type Definition`app/blog/[slug]/page.js``{ slug: string }``app/shop/[...slug]/page.js``{ slug: string[] }``app/shop/[[...slug]]/page.js``{ slug?: string[] }``app/[categoryId]/[itemId]/page.js``{ categoryId: string, itemId: string }`

If you're working on a route where `params` can only have a fixed number of valid values, such as a `[locale]` param with a known set of language codes, you can use runtime validation to handle any invalid params a user may enter, and let the rest of your application work with the narrower type from your known set.

/app/\[locale]/page.tsx

```
import { notFound } from 'next/navigation'
import type { Locale } from '@i18n/types'
import { isValidLocale } from '@i18n/utils'
 
function assertValidLocale(value: string): asserts value is Locale {
  if (!isValidLocale(value)) notFound()
}
 
export default async function Page(props: PageProps<'/[locale]'>) {
  const { locale } = await props.params // locale is typed as string
  assertValidLocale(locale)
  // locale is now typed as Locale
}
```

## [Behavior](#behavior)

- Since the `params` prop is a promise. You must use `async`/`await` or React's use function to access the values.
  
  - In version 14 and earlier, `params` was a synchronous prop. To help with backwards compatibility, you can still access it synchronously in Next.js 15, but this behavior will be deprecated in the future.

## [Examples](#examples)

### [With `generateStaticParams`](#with-generatestaticparams)

The [`generateStaticParams`](/docs/15/app/api-reference/functions/generate-static-params) function can be used to [statically generate](/docs/15/app/getting-started/partial-prerendering#static-rendering) routes at build time instead of on-demand at request time.

app/blog/\[slug]/page.tsx

TypeScript

JavaScriptTypeScript

```
export async function generateStaticParams() {
  const posts = await fetch('https://.../posts').then((res) => res.json())
 
  return posts.map((post) => ({
    slug: post.slug,
  }))
}
```

When using `fetch` inside the `generateStaticParams` function, the requests are [automatically deduplicated](/docs/15/app/guides/caching#request-memoization). This avoids multiple network calls for the same data Layouts, Pages, and other `generateStaticParams` functions, speeding up build time.

## Next Steps

For more information on what to do next, we recommend the following sections

[Next.js Docs  
\
...  
\
Functions  
\
**generateStaticParams**  
\
API reference for the generateStaticParams function.](/docs/15/app/api-reference/functions/generate-static-params)

Was this helpful?

supported.

Send
