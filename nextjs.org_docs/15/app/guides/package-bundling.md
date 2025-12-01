---
title: "How to optimize package bundling"
source: "https://nextjs.org/docs/15/app/guides/package-bundling"
---

# How to optimize package bundling

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[App Router](/docs/15/app)[Guides](/docs/15/app/guides)Package Bundling

You are currently viewing documentation for version 15 of Next.js.

# How to optimize package bundling

Bundling external packages can significantly improve the performance of your application. By default, packages imported inside Server Components and Route Handlers are automatically bundled by Next.js. This page will guide you through how to analyze and further optimize package bundling.

## [Analyzing JavaScript bundles](#analyzing-javascript-bundles)

[`@next/bundle-analyzer`](https://www.npmjs.com/package/@next/bundle-analyzer) is a plugin for Next.js that helps you manage the size of your application bundles. It generates a visual report of the size of each package and their dependencies. You can use the information to remove large dependencies, split, or [lazy-load](/docs/15/app/guides/lazy-loading) your code.

### [Installation](#installation)

Install the plugin by running the following command:

```
npm i @next/bundle-analyzer
# or
yarn add @next/bundle-analyzer
# or
pnpm add @next/bundle-analyzer
```

Then, add the bundle analyzer's settings to your `next.config.js`.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {}
 
const withBundleAnalyzer = require('@next/bundle-analyzer')({
  enabled: process.env.ANALYZE === 'true',
})
 
module.exports = withBundleAnalyzer(nextConfig)
```

### [Generating a report](#generating-a-report)

Run the following command to analyze your bundles:

```
ANALYZE=true npm run build
# or
ANALYZE=true yarn build
# or
ANALYZE=true pnpm build
```

The report will open three new tabs in your browser, which you can inspect. Periodically evaluating your application's bundles can help you maintain application performance over time.

## [Optimizing package imports](#optimizing-package-imports)

Some packages, such as icon libraries, can export hundreds of modules, which can cause performance issues in development and production.

You can optimize how these packages are imported by adding the [`optimizePackageImports`](/docs/15/app/api-reference/config/next-config-js/optimizePackageImports) option to your `next.config.js`. This option will only load the modules you *actually* use, while still giving you the convenience of writing import statements with many named exports.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  experimental: {
    optimizePackageImports: ['icon-library'],
  },
}
 
module.exports = nextConfig
```

Next.js also optimizes some libraries automatically, thus they do not need to be included in the optimizePackageImports list. See the [full list](/docs/15/app/api-reference/config/next-config-js/optimizePackageImports).

## [Opting specific packages out of bundling](#opting-specific-packages-out-of-bundling-1)

Since packages imported inside Server Components and Route Handlers are automatically bundled by Next.js, you can opt specific packages out of bundling using the [`serverExternalPackages`](/docs/15/app/api-reference/config/next-config-js/serverExternalPackages) option in your `next.config.js`.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  serverExternalPackages: ['package-name'],
}
 
module.exports = nextConfig
```

Next.js includes a list of popular packages that currently are working on compatibility and automatically opt-ed out. See the [full list](/docs/15/app/api-reference/config/next-config-js/serverExternalPackages).

## Next Steps

Learn more about optimizing your application for production.

[Next.js Docs  
\
...  
\
Guides  
\
**Production**  
\
Recommendations to ensure the best performance and user experience before taking your Next.js application to production.](/docs/15/app/guides/production-checklist)

Was this helpful?

supported.

Send
