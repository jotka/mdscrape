---
title: "Upgrading"
source: "https://nextjs.org/docs/15/app/getting-started/upgrading"
---

# Upgrading

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[App Router](/docs/15/app)[Getting Started](/docs/15/app/getting-started)Upgrading

You are currently viewing documentation for version 15 of Next.js.

# Upgrading

## [Latest version](#latest-version)

To update to the latest version of Next.js, you can use the `upgrade` codemod:

Terminal

```
npx @next/codemod@latest upgrade latest
```

If you prefer to upgrade manually, install the latest Next.js and React versions:

pnpmnpmyarnbun

Terminal

```
pnpm i next@latest react@latest react-dom@latest eslint-config-next@latest
```

## [Canary version](#canary-version)

To update to the latest canary, make sure you're on the latest version of Next.js and everything is working as expected. Then, run the following command:

Terminal

```
npm i next@canary
```

### [Features available in canary](#features-available-in-canary)

The following features are currently available in canary:

**Caching**:

- [`"use cache"`](/docs/15/app/api-reference/directives/use-cache)
- [`cacheLife`](/docs/15/app/api-reference/functions/cacheLife)
- [`cacheTag`](/docs/15/app/api-reference/functions/cacheTag)
- [`cacheComponents`](/docs/15/app/api-reference/config/next-config-js/cacheComponents)

**Authentication**:

- [`forbidden`](/docs/15/app/api-reference/functions/forbidden)
- [`unauthorized`](/docs/15/app/api-reference/functions/unauthorized)
- [`forbidden.js`](/docs/15/app/api-reference/file-conventions/forbidden)
- [`unauthorized.js`](/docs/15/app/api-reference/file-conventions/unauthorized)
- [`authInterrupts`](/docs/15/app/api-reference/config/next-config-js/authInterrupts)

## Version guides

See the version guides for in-depth upgrade instructions.

[Next.js Docs  
\
...  
\
Upgrading  
\
**Version 15**  
\
Upgrade your Next.js Application from Version 14 to 15.](/docs/15/app/guides/upgrading/version-15)

[Next.js Docs  
\
...  
\
Upgrading  
\
**Version 14**  
\
Upgrade your Next.js Application from Version 13 to 14.](/docs/15/app/guides/upgrading/version-14)

Was this helpful?

supported.

Send
