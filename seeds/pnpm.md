# pnpm

Fast, disk-space efficient package manager. Symlink + hardlink storage.

## Setup

### pnpm init
**Deskripsi:** Buat `package.json` baru.
```bash
pnpm init -y
pnpm init
```

### pnpm install
**Deskripsi:** Install dependensi (buat symlink dari `.pnpm`).
```bash
pnpm install
pnpm i
```

---

## Install Packages

### pnpm add PACKAGE
**Deskripsi:** Install package production.
```bash
pnpm add lodash
pnpm add lodash@latest
pnpm add lodash@4.17.21
```

### pnpm add -D PACKAGE
**Deskripsi:** Install dev dependency.
```bash
pnpm add -D typescript
pnpm add -D jest
pnpm add -D @types/node
```

### pnpm remove
**Deskripsi:** Hapus package.
```bash
pnpm remove lodash
pnpm rm typescript
```

---

## Run Scripts

### pnpm run
**Deskripsi:** Jalankan script di `package.json`.
```bash
pnpm run build
pnpm run start
pnpm run test
pnpm run dev
```

### pnpm [SCRIPT]
**Deskripsi:** Shorthand.
```bash
pnpm build
pnpm start
pnpm test
```

### pnpm with args
**Deskripsi:** Pass argumen.
```bash
pnpm start -- --port 3000
pnpm build -- --watch
```

---

## Info

### pnpm list
**Deskripsi:** Lihat installed packages.
```bash
pnpm list
pnpm list --depth=0
pnpm list lodash
```

### pnpm why
**Deskripsi:** Lihat dependency chain.
```bash
pnpm why lodash
```

### pnpm store
**Deskripsi:** Info tentang store (~/.pnpm-store).
```bash
pnpm store status
pnpm store prune
```

---

## Global

### pnpm add -g
**Deskripsi:** Install global.
```bash
pnpm add -g typescript
pnpm add -g @angular/cli
```

### pnpm list -g
**Deskripsi:** Lihat global packages.
```bash
pnpm list -g
```

---

## Update / Upgrade

### pnpm up
**Deskripsi:** Update packages.
```bash
pnpm up
pnpm up lodash
pnpm up --latest
```

---

## Monorepo (Workspaces)

### pnpm recursive
**Deskripsi:** Run command di semua workspace.
```bash
pnpm -r install
pnpm -r build
pnpm -r test
```

### pnpm filter
**Deskripsi:** Run di specific workspace.
```bash
pnpm --filter @myapp/core add lodash
pnpm --filter @myapp/web build
```

---

## Troubleshooting

### pnpm prune
**Deskripsi:** Remove unused packages.
```bash
pnpm prune
```

### pnpm store prune
**Deskripsi:** Clean store cache.
```bash
pnpm store prune
```

### rm node_modules + reinstall
**Deskripsi:** Clean reinstall.
```bash
rm -rf node_modules pnpm-lock.yaml
pnpm install
```

---

## Tips

- Lock file: `pnpm-lock.yaml` (commit ke git)
- Strict mode: sub-dependencies harus explicit di `package.json`
- Storage: `~/.pnpm-store/` (shared, symlinked)
- Lebih cepat + lebih hemat disk dibanding npm & yarn
- Monorepo: `pnpm-workspace.yaml` untuk filter
