---
title: "How to optimize package bundling"
source: "https://nextjs.org/docs/pages/guides/package-bundling"
---

# How to optimize package bundling

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Pages Router](/docs/pages)[Guides](/docs/pages/guides)Package Bundling

Copy page

# How to optimize package bundling

Bundling external packages can significantly improve the performance of your application. By default, packages imported into your application are not bundled. This can impact performance or might not work if external packages are not pre-bundled, for example, if imported from a monorepo or `node_modules`. This page will guide you through how to analyze and configure package bundling.

## [Analyzing JavaScript bundles](#analyzing-javascript-bundles)

[`@next/bundle-analyzer`](https://www.npmjs.com/package/@next/bundle-analyzer) is a plugin for Next.js that helps you manage the size of your application bundles. It generates a visual report of the size of each package and their dependencies. You can use the information to remove large dependencies, split, or [lazy-load](/docs/app/guides/lazy-loading) your code.

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

You can optimize how these packages are imported by adding the [`optimizePackageImports`](/docs/app/api-reference/config/next-config-js/optimizePackageImports) option to your `next.config.js`. This option will only load the modules you *actually* use, while still giving you the convenience of writing import statements with many named exports.

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

Next.js also optimizes some libraries automatically, thus they do not need to be included in the optimizePackageImports list. See the [full list](/docs/app/api-reference/config/next-config-js/optimizePackageImports).

## [Bundling specific packages](#bundling-specific-packages)

To bundle specific packages, you can use the [`transpilePackages`](/docs/app/api-reference/config/next-config-js/transpilePackages) option in your `next.config.js`. This option is useful for bundling external packages that are not pre-bundled, for example, in a monorepo or imported from `node_modules`.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  transpilePackages: ['package-name'],
}
 
module.exports = nextConfig
```

## [Bundling all packages](#bundling-all-packages)

To automatically bundle all packages (default behavior in the App Router), you can use the [`bundlePagesRouterDependencies`](/docs/pages/api-reference/config/next-config-js/bundlePagesRouterDependencies) option in your `next.config.js`.

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  bundlePagesRouterDependencies: true,
}
 
module.exports = nextConfig
```

## [Opting specific packages out of bundling](#opting-specific-packages-out-of-bundling)

If you have the [`bundlePagesRouterDependencies`](/docs/pages/api-reference/config/next-config-js/bundlePagesRouterDependencies) option enabled, you can opt specific packages out of automatic bundling using the [`serverExternalPackages`](/docs/pages/api-reference/config/next-config-js/serverExternalPackages) option in your `next.config.js`:

next.config.js

```
/** @type {import('next').NextConfig} */
const nextConfig = {
  // Automatically bundle external packages in the Pages Router:
  bundlePagesRouterDependencies: true,
  // Opt specific packages out of bundling for both App and Pages Router:
  serverExternalPackages: ['package-name'],
}
 
module.exports = nextConfig
```

## Next Steps

Learn more about optimizing your application for production.

[**Production**  
\
Recommendations to ensure the best performance and user experience before taking your Next.js application to production.](/docs/pages/guides/production-checklist)

Was this helpful?

supported.

Send
