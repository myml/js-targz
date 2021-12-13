# js-targz

Preview tar.gz in the browser

## HTML script

### CDN

`<script src="https://cdn.jsdelivr.net/npm/@myml/js-targz@0.0.4/index.js"></script>`

### Demo

See [./demo.html](./demo.html)

## Module import

### Install

`yarn add @myml/js-targz`

### Import

`import { targz } from '@myml/js-targz'`

### Use

```ts
targz(files[0], (header) => {
  console.log(header);
});
```

```ts
targz(files[0], (header) => {
  // stop loop
  if (header.Name == "found_file") {
    return true;
  }
  console.log(header);
});
```
