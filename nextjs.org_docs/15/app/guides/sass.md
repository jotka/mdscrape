---
title: "How to use Sass"
source: "https://nextjs.org/docs/15/app/guides/sass"
---

# How to use Sass

Menu

App Router 15

Using App Router

Features available in /app

Version 15

15.5.6

[App Router](/docs/15/app)[Guides](/docs/15/app/guides)Sass

You are currently viewing documentation for version 15 of Next.js.

# How to use Sass

Next.js has built-in support for integrating with Sass after the package is installed using both the `.scss` and `.sass` extensions. You can use component-level Sass via CSS Modules and the `.module.scss`or `.module.sass` extension.

First, install [`sass`](https://github.com/sass/sass):

Terminal

```
npm install --save-dev sass
```

> **Good to know**:
> 
> Sass supports [two different syntaxes](https://sass-lang.com/documentation/syntax), each with their own extension. The `.scss` extension requires you use the [SCSS syntax](https://sass-lang.com/documentation/syntax#scss), while the `.sass` extension requires you use the [Indented Syntax ("Sass")](https://sass-lang.com/documentation/syntax#the-indented-syntax).
> 
> If you're not sure which to choose, start with the `.scss` extension which is a superset of CSS, and doesn't require you learn the Indented Syntax ("Sass").

### [Customizing Sass Options](#customizing-sass-options)

If you want to configure your Sass options, use `sassOptions` in `next.config`.

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  sassOptions: {
    additionalData: `$var: red;`,
  },
}
 
export default nextConfig
```

#### [Implementation](#implementation)

You can use the `implementation` property to specify the Sass implementation to use. By default, Next.js uses the [`sass`](https://www.npmjs.com/package/sass) package.

next.config.ts

TypeScript

JavaScriptTypeScript

```
import type { NextConfig } from 'next'
 
const nextConfig: NextConfig = {
  sassOptions: {
    implementation: 'sass-embedded',
  },
}
 
export default nextConfig
```

### [Sass Variables](#sass-variables)

Next.js supports Sass variables exported from CSS Module files.

For example, using the exported `primaryColor` Sass variable:

app/variables.module.scss

```
$primary-color: #64ff00;
 
:export {
  primaryColor: $primary-color;
}
```

app/page.js

```
// maps to root `/` URL
 
import variables from './variables.module.scss'
 
export default function Page() {
  return <h1 style={{ color: variables.primaryColor }}>Hello, Next.js!</h1>
}
```

Was this helpful?

supported.

Send
