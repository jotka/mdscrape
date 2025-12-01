---
title: "default.js"
source: "https://nextjs.org/docs/15/app/api-reference/file-conventions/default"
---

# default.js

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[File-system conventions](/docs/15/app/api-reference/file-conventions)default.js

You are currently viewing documentation for version 15 of Next.js.

# default.js

The `default.js` file is used to render a fallback within [Parallel Routes](/docs/15/app/api-reference/file-conventions/parallel-routes) when Next.js cannot recover a [slot's](/docs/15/app/api-reference/file-conventions/parallel-routes#slots) active state after a full-page load.

During [soft navigation](/docs/15/app/getting-started/linking-and-navigating#client-side-transitions), Next.js keeps track of the active *state* (subpage) for each slot. However, for hard navigations (full-page load), Next.js cannot recover the active state. In this case, a `default.js` file can be rendered for subpages that don't match the current URL.

Consider the following folder structure. The `@team` slot has a `settings` page, but `@analytics` does not.

![Parallel Routes unmatched routes](/_next/image?url=https%3A%2F%2Fh8DxKfmAPhn8O0p3.public.blob.vercel-storage.com%2Fdocs%2Flight%2Fparallel-routes-unmatched-routes.png&w=3840&q=75)![Parallel Routes unmatched routes](/_next/image?url=https%3A%2F%2Fh8DxKfmAPhn8O0p3.public.blob.vercel-storage.com%2Fdocs%2Fdark%2Fparallel-routes-unmatched-routes.png&w=3840&q=75)

When navigating to `/settings`, the `@team` slot will render the `settings` page while maintaining the currently active page for the `@analytics` slot.

On refresh, Next.js will render a `default.js` for `@analytics`. If `default.js` doesn't exist, a `404` is rendered instead.

Additionally, since `children` is an implicit slot, you also need to create a `default.js` file to render a fallback for `children` when Next.js cannot recover the active state of the parent page.

## [Reference](#reference)

### [`params` (optional)](#params-optional)

A promise that resolves to an object containing the [dynamic route parameters](/docs/15/app/api-reference/file-conventions/dynamic-routes) from the root segment down to the slot's subpages. For example:

app/\[artist]/@sidebar/default.js

TypeScript

JavaScriptTypeScript

```
export default async function Default({
  params,
}: {
  params: Promise<{ artist: string }>
}) {
  const { artist } = await params
}
```

ExampleURL`params``app/[artist]/@sidebar/default.js``/zack``Promise<{ artist: 'zack' }>``app/[artist]/[album]/@sidebar/default.js``/zack/next``Promise<{ artist: 'zack', album: 'next' }>`

- Since the `params` prop is a promise. You must use `async/await` or React's [`use`](https://react.dev/reference/react/use) function to access the values.
  
  - In version 14 and earlier, `params` was a synchronous prop. To help with backwards compatibility, you can still access it synchronously in Next.js 15, but this behavior will be deprecated in the future.

## Learn more about Parallel Routes

[Next.js Docs  
\
...  
\
File-system conventions  
\
**Parallel Routes**  
\
Simultaneously render one or more pages in the same view that can be navigated independently. A pattern for highly dynamic applications.](/docs/15/app/api-reference/file-conventions/parallel-routes)

Was this helpful?

supported.

Send
