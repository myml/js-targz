# js-targz

Preview tar.gz in the browser

## HTML script

### Import

`<script src="./index.js"></script>`

### Use

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
  return true;
});
```
