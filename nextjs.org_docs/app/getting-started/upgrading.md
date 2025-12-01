---
title: "Upgrading"
source: "https://nextjs.org/docs/app/getting-started/upgrading"
---

# Upgrading

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[App Router](/docs/app)[Getting Started](/docs/app/getting-started)Upgrading

Copy page

# Upgrading

## [Latest version](#latest-version)

To update to the latest version of Next.js, you can use the `upgrade` command:

Terminal

```
next upgrade
```

Next.js 15 and earlier do not support the `upgrade` command and need to use a separate package instead:

Terminal

```
npx @next/codemod@canary upgrade latest
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

**Authentication**:

- [`forbidden`](/docs/app/api-reference/functions/forbidden)
- [`unauthorized`](/docs/app/api-reference/functions/unauthorized)
- [`forbidden.js`](/docs/app/api-reference/file-conventions/forbidden)
- [`unauthorized.js`](/docs/app/api-reference/file-conventions/unauthorized)
- [`authInterrupts`](/docs/app/api-reference/config/next-config-js/authInterrupts)

## Version guides

See the version guides for in-depth upgrade instructions.

[**Version 16**  
\
Upgrade your Next.js Application from Version 15 to 16.](/docs/app/guides/upgrading/version-16)

[**Version 15**  
\
Upgrade your Next.js Application from Version 14 to 15.](/docs/app/guides/upgrading/version-15)

[**Version 14**  
\
Upgrade your Next.js Application from Version 13 to 14.](/docs/app/guides/upgrading/version-14)

Was this helpful?

supported.

Send
