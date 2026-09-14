```markdown
# Git

Version control. Track changes, collaborate, manage history.

## Setup

Konfigurasi awal (satu kali):

```bash
git config --global user.name "Your Name"
git config --global user.email "you@example.com"
```

---

## Daily Workflow

### git status
**Deskripsi:** Lihat file yang berubah, staged, untracked.
```bash
git status
```

### git add
**Deskripsi:** Stage file untuk commit. `-p` = pilih per chunk.
```bash
git add FILE
git add .
git add -p
```

### git commit
**Deskripsi:** Simpan perubahan dengan pesan deskriptif.
```bash
git commit -m "pesan singkat"
git commit -am "auto-stage tracked files + commit"
```

### git pull --rebase
**Deskripsi:** Fetch + rebase (history rapi, tanpa merge commit).
```bash
git pull --rebase origin main
```

### git push
**Deskripsi:** Kirim commit ke remote.
```bash
git push
git push origin feature/x
```

---

## Branch

### git switch -c
**Deskripsi:** Buat & pindah branch baru.
```bash
git switch -c feature/nama
git switch -c bugfix/issue-123
```

### git switch
**Deskripsi:** Pindah ke branch existing.
```bash
git switch main
git switch develop
```

### git branch -d
**Deskripsi:** Hapus branch lokal (aman, cek merged).
```bash
git branch -d feature/x
git branch -D feature/x  # force delete
```

### git branch
**Deskripsi:** Lihat semua branch.
```bash
git branch -a
git branch -v
```

---

## Undo / Fix

### git restore --staged
**Deskripsi:** Unstage file (batalkan `git add`).
```bash
git restore --staged FILE
git restore --staged .
```

### git restore
**Deskripsi:** Discard perubahan di file (kembali ke HEAD).
```bash
git restore FILE
git restore .
```

### git reset --soft
**Deskripsi:** Undo commit terakhir, keep staged.
```bash
git reset --soft HEAD~1
```

### git reset --hard
**Deskripsi:** Undo commit & buang perubahan (hati-hati!).
```bash
git reset --hard HEAD~1
```

### git revert
**Deskripsi:** Buat commit baru yang undo commit lama (history tetap).
```bash
git revert HEAD
git revert abc123
```

---

## View History

### git log
**Deskripsi:** Lihat daftar commit.
```bash
git log --oneline -20
git log --graph --oneline --all
git log --author="name"
```

### git diff
**Deskripsi:** Lihat perubahan (unstaged vs HEAD).
```bash
git diff
git diff --staged
git diff main..feature/x
```

### git show
**Deskripsi:** Lihat detail commit tertentu.
```bash
git show HEAD
git show abc123
git show HEAD:file.txt
```

---

## Remote

### git clone
**Deskripsi:** Clone repo dari remote.
```bash
git clone https://github.com/user/repo.git
git clone git@github.com:user/repo.git
```

### git remote
**Deskripsi:** Kelola remote URL.
```bash
git remote -v
git remote add origin https://...
git remote set-url origin https://...
```

### git fetch
**Deskripsi:** Download changes dari remote (no merge).
```bash
git fetch origin
git fetch --all
```

---

## Tips

- Commit sering, pesan jelas: `git add -p` → review per baris
- Rebase sebelum push: `git pull --rebase`
- Cek branch sebelum commit: `git status` + `git branch`
- Backup branch penting sebelum reset: `git branch backup-name`
```