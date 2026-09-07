# Yarn

Fast, reliable package manager. Alternative ke npm.

## Setup

### yarn init
**Deskripsi:** Buat `package.json` baru.
```bash
yarn init -y
yarn init
```

### yarn install
**Deskripsi:** Install dependensi (dari `package.json` & lock file).
```bash
yarn install
yarn
```

---

## Install Packages

### yarn add PACKAGE
**Deskripsi:** Install package production (update `package.json`).
```bash
yarn add lodash
yarn add lodash@latest
yarn add lodash@4.17.21
```

### yarn add -D PACKAGE
**Deskripsi:** Install dev dependency.
```bash
yarn add -D typescript
yarn add -D jest
yarn add -D @types/node
```

### yarn remove
**Deskripsi:** Hapus package.
```bash
yarn remove lodash
yarn remove typescript
```

---

## Run Scripts

### yarn run
**Deskripsi:** Jalankan script di `package.json`.
```bash
yarn run build
yarn run start
yarn run test
yarn run dev
```

### yarn [SCRIPT]
**Deskripsi:** Shorthand (tanpa `run`).
```bash
yarn build
yarn start
yarn test
```

### yarn with args
**Deskripsi:** Pass argumen.
```bash
yarn start --port 3000
yarn build --watch
```

---

## Info

### yarn list
**Deskripsi:** Lihat installed packages.
```bash
yarn list
yarn list --depth=0
yarn list lodash
```

### yarn why
**Deskripsi:** Lihat why package installed (dependency chain).
```bash
yarn why lodash
```

---

## Global

### yarn global add
**Deskripsi:** Install global.
```bash
yarn global add typescript
yarn global add @angular/cli
```

### yarn global list
**Deskripsi:** Lihat global packages.
```bash
yarn global list
```

---

## Upgrade / Update

### yarn up
**Deskripsi:** Update package ke latest (dalam semver range).
```bash
yarn up lodash
yarn up
```

### yarn upgrade
**Deskripsi:** Upgrade dengan flexibility (bisa minor/major).
```bash
yarn upgrade lodash@latest
```

---

## Workspaces (Monorepo)

### yarn workspaces list
**Deskripsi:** Lihat workspace di monorepo.
```bash
yarn workspaces list
```

### yarn workspace WORKSPACE add
**Deskripsi:** Add dependency ke specific workspace.
```bash
yarn workspace @myapp/core add lodash
```

---

## Troubleshooting

### yarn cache clean
**Deskripsi:** Clear cache.
```bash
yarn cache clean
```

### yarn autoclean
**Deskripsi:** Remove unused packages.
```bash
yarn autoclean --force
```

### rm node_modules + reinstall
**Deskripsi:** Clean reinstall.
```bash
rm -rf node_modules yarn.lock
yarn install
```

---

## Tips

- Lock file: `yarn.lock` (commit ke git)
- Lebih cepat dari npm (parallel, better caching)
- `yarn add` langsung tanpa `-D` untuk prod dependency
- Monorepo support (workspaces) built-in
