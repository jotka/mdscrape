---
title: "mdx-components.js"
source: "https://nextjs.org/docs/15/app/api-reference/file-conventions/mdx-components"
---

# mdx-components.js

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[API Reference](/docs/15/app/api-reference)[File-system conventions](/docs/15/app/api-reference/file-conventions)mdx-components.js

You are currently viewing documentation for version 15 of Next.js.

# mdx-components.js

The `mdx-components.js|tsx` file is **required** to use [`@next/mdx` with App Router](/docs/15/app/guides/mdx) and will not work without it. Additionally, you can use it to [customize styles](/docs/15/app/guides/mdx#using-custom-styles-and-components).

Use the file `mdx-components.tsx` (or `.js`) in the root of your project to define MDX Components. For example, at the same level as `pages` or `app`, or inside `src` if applicable.

mdx-components.tsx

TypeScript

JavaScriptTypeScript

```
import type { MDXComponents } from 'mdx/types'
 
const components: MDXComponents = {}
 
export function useMDXComponents(): MDXComponents {
  return components
}
```

## [Exports](#exports)

### [`useMDXComponents` function](#usemdxcomponents-function)

The file must export a single function named `useMDXComponents`. This function does not accept any arguments.

mdx-components.tsx

TypeScript

JavaScriptTypeScript

```
import type { MDXComponents } from 'mdx/types'
 
const components: MDXComponents = {}
 
export function useMDXComponents(): MDXComponents {
  return components
}
```

## [Version History](#version-history)

VersionChanges`v13.1.2`MDX Components added

## Learn more about MDX Components

[Next.js Docs  
\
...  
\
Guides  
\
**MDX**  
\
Learn how to configure MDX and use it in your Next.js apps.](/docs/15/app/guides/mdx)

Was this helpful?

supported.

Send
