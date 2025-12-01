---
title: "Google Font Display"
source: "https://nextjs.org/docs/messages/google-font-display"
---

# Google Font Display

Menu

Using App Router

Features available in /app

Latest Version

16.0.5

[Docs](/docs)[Errors](/docs)Google Font Display

# Google Font Display

> Enforce font-display behavior with Google Fonts.

## [Why This Error Occurred](#why-this-error-occurred)

For a Google Font, the font-display descriptor was either missing or set to `auto`, `block`, or `fallback`, which are not recommended.

## [Possible Ways to Fix It](#possible-ways-to-fix-it)

For most cases, the best font display strategy for custom fonts is `optional`.

pages/index.js

```
import Head from 'next/head'
 
export default function IndexPage() {
  return (
    <div>
      <Head>
        <link
          href="https://fonts.googleapis.com/css2?family=Krona+One&display=optional"
          rel="stylesheet"
        />
      </Head>
    </div>
  )
}
```

Specifying `display=optional` minimizes the risk of invisible text or layout shift. If swapping to the custom font after it has loaded is important to you, then use `display=swap` instead.

### [When Not To Use It](#when-not-to-use-it)

If you want to specifically display a font using an `auto`, `block`, or `fallback` strategy, then you can disable this rule.

## [Useful Links](#useful-links)

- [Controlling Font Performance with font-display](https://developer.chrome.com/blog/font-display/)
- [Google Fonts API Docs](https://developers.google.com/fonts/docs/css2#use_font-display)
- [CSS `font-display` property](https://www.w3.org/TR/css-fonts-4/#font-display-desc)

Was this helpful?

supported.

Send
