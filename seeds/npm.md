# npm

Node Package Manager. Install library, run script, publish package.

## Setup

### npm init
**Deskripsi:** Buat `package.json` baru.
```bash
npm init -y
npm init
```

### npm install
**Deskripsi:** Install semua dependensi (dari `package.json`).
```bash
npm install
npm ci  # cleaner install (exact versions, CI)
```

---

## Install Packages

### npm install PACKAGE
**Deskripsi:** Install package production (update `package.json`).
```bash
npm install lodash
npm i lodash  # alias
npm install lodash@latest
npm install lodash@4.17.21
```

### npm install -D PACKAGE
**Deskripsi:** Install dev dependency (testing, build tools).
```bash
npm install -D typescript
npm install -D jest
npm i -D @types/node
```

### npm uninstall
**Deskripsi:** Hapus package.
```bash
npm uninstall lodash
npm remove typescript
npm rm lodash
```

---

## Run Scripts

### npm run
**Deskripsi:** Jalankan script di `package.json` → `"scripts"`.
```bash
npm run build
npm run start
npm run test
npm run dev
```

### npm run with args
**Deskripsi:** Pass argumen ke script.
```bash
npm run start -- --port 3000
npm run build -- --watch
```

### npm exec / npx
**Deskripsi:** Jalankan CLI tool (dari node_modules atau npm registry).
```bash
npm exec tsc
npx tsc --init
npx create-vite my-app
npx @latest some-tool
```

### npm scripts available
**Deskripsi:** Lihat semua script di project.
```bash
npm run
```

---

## Publish / Version

### npm version
**Deskripsi:** Update versi (major/minor/patch).
```bash
npm version patch
npm version minor
npm version major
```

### npm publish
**Deskripsi:** Publish package ke npm registry.
```bash
npm publish
npm publish --access public
npm publish --tag beta
```

---

## Info

### npm list
**Deskripsi:** Lihat installed packages & version.
```bash
npm list
npm list --depth=0
npm list lodash
```

### npm view
**Deskripsi:** Info package dari registry.
```bash
npm view lodash
npm view lodash versions
npm view lodash dist-tags
```

### npm outdated
**Deskripsi:** Lihat package yang bisa di-update.
```bash
npm outdated
npm update
```

---

## Global

### npm install -g
**Deskripsi:** Install global (command line tool).
```bash
npm install -g typescript
npm install -g @angular/cli
npm i -g pnpm
```

### npm list -g
**Deskripsi:** Lihat global packages.
```bash
npm list -g --depth=0
```

---

## Troubleshooting

### npm cache clean
**Deskripsi:** Clear cache (kalau install error).
```bash
npm cache clean --force
```

### npm audit
**Deskripsi:** Check security vulnerabilities.
```bash
npm audit
npm audit fix
npm audit fix --force
```

### rm node_modules + reinstall
**Deskripsi:** Nuclear option (hapus folder, reinstall fresh).
```bash
rm -rf node_modules package-lock.json
npm install
```

---

## Tips

- Lock file: `package-lock.json` (production), commit ke git
- `npm ci` di CI/CD (exact versions), `npm install` di development
- `npm run test` atau `npm test` (alias `npm t`)
- Dev dependency: test, linter, bundler (tidak di production build)
