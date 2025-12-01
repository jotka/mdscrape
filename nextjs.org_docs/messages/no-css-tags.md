---
title: "No CSS Tags"
source: "https://nextjs.org/docs/messages/no-css-tags"
---

# No CSS Tags

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)No CSS Tags

# No CSS Tags

> Prevent manual stylesheet tags.

## [Why This Error Occurred](#why-this-error-occurred)

A `link` element was used to link to an external stylesheet. This can negatively affect CSS resource loading on your webpage.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

There are multiple ways to include styles using Next.js' built-in CSS support, including the option to use `@import` within the root stylesheet that is imported in `pages/_app.js`:

styles.css

```
/* Root stylesheet */
@import 'extra.css';
 
body {
  /* ... */
}
```

Another option is to use CSS Modules to import the CSS file scoped specifically to the component.

pages/index.js

```
import styles from './extra.module.css'
 
export class Home {
  render() {
    return (
      <div>
        <button type="button" className={styles.active}>
          Open
        </button>
      </div>
    )
  }
}
```

Refer to the [Built-In CSS Support](/docs/app/getting-started/css) documentation to learn about all the ways to include CSS to your application.

Was this helpful?

supported.

Send
