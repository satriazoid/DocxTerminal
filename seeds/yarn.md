# Yarn

Cheatsheet. Edit this file. Yarn 1 / Berry both work for most of this.

## Init / install

```
yarn init -y
yarn
yarn add lodash
yarn add -D typescript
yarn remove lodash
```

## Run

```
yarn build
yarn tsc
yarn dlx create-vite
```

## Workspaces (Berry)

```
yarn workspaces foreach -A run build
yarn workspace pkg-a add lodash
```
