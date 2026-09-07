# DocxTerminal

Markdown cheat sheets di terminal. File-based CRUD. TUI + CLI.

Ketik `dt` → menu. Edit pakai `$EDITOR` / vim / nano.

![DocxTerminal TUI](.image/docxterminal-ui.png)

Screenshot lain taruh di [`.image/`](.image/).

## Fitur

- TUI: List / Add / Edit / Delete / Quit
- CLI: `list`, `show`, `add`, `edit`, `delete`
- Storage: `~/.docxterminal/*.md` (satu file = satu dokumen)
- Seed pertama kali: `git`, `npm`, `pnpm`, `yarn` (dengan template lengkap & deskripsi)
- Editor: `$EDITOR` → `$VISUAL` → `vim` → `nano` → `notepad` (Windows)
- Nama dokumen: huruf, angka, `.` `_` `-` (path traversal ditolak)
- TUI List → Enter: markdown di-render Glamour (heading, list, code fence, auto light/dark, wrap ulang saat resize). `dt show NAME` tetap print source mentah, aman untuk pipe/script.

## Install

Butuh [Go](https://go.dev/dl/) 1.24+.

### Windows (Git Bash / CMD)

```bash
git clone https://github.com/satriazoid/DocxTerminal.git
cd DocxTerminal
bash install.sh
# atau: install.bat
# atau: powershell -ExecutionPolicy Bypass -File install.ps1
```

### Linux / macOS

```bash
git clone https://github.com/satriazoid/DocxTerminal.git
cd DocxTerminal
bash install.sh
```

### Manual (kalau script tidak bisa)

```bash
git clone https://github.com/satriazoid/DocxTerminal.git
cd DocxTerminal
go build -o dt .
# Copy dt (atau dt.exe) ke folder yang ada di PATH
# Contoh: ~/.local/bin/ (Linux/macOS) atau %USERPROFILE%\go\bin (Windows)
```

### Verifikasi

Restart terminal, lalu:

```bash
dt
dt list
```

### Rebuild setelah ubah kode

```bash
bash install.sh
# atau manual: go build -o dt . && mv dt ~/.local/bin/dt
```

## Usage

### TUI

```bash
dt
```

| Key | Aksi |
|-----|------|
| `j` / `k` atau `↓` / `↑` | pindah |
| `Enter` | pilih |
| `e` | edit (saat view) |
| `y` / `n` | konfirmasi delete |
| `q` / `Esc` | kembali / keluar |
| `Ctrl+C` | keluar |

Menu:

- **List** lihat dokumen, Enter = baca (Glamour), `e` = buka editor
- **Add** ketik nama, Enter = buat + buka editor
- **Edit** pilih dokumen, buka editor
- **Delete** pilih, konfirmasi `y`
- **Quit**

### CLI

```bash
dt list              # daftar dokumen
dt show git          # print markdown
dt add docker        # buat + buka editor
dt edit git          # buka editor
dt delete docker     # hapus file
dt help
```

Alias: `ls`, `cat`, `rm`.

## Storage

```
~/.docxterminal/
├── git.md
├── npm.md
├── pnpm.md
└── yarn.md
```

Windows: `%USERPROFILE%\.docxterminal\`

Seed otomatis saat pertama kali (kalau folder kosong). Dokumen yang sudah ada tidak ditimpa.

Edit langsung file `.md` juga valid — TUI/CLI baca ulang dari disk.

## Layout repo

```
DocxTerminal/
├── main.go          CLI commands
├── tui.go           menu TUI (Bubble Tea)
├── store.go         CRUD file
├── editor.go        $EDITOR / vim / nano
├── store_test.go
├── tui_test.go
├── install.sh       Installer (Linux/macOS/Git Bash)
├── install.bat      Installer (Windows CMD)
├── install.ps1      Installer (Windows PowerShell)
├── seeds/           template awal
│   ├── git.md
│   ├── npm.md
│   ├── pnpm.md
│   └── yarn.md
├── .image/          screenshot README
└── README.md
```

## Develop

```bash
go test ./...
go build -o dt .
bash install.sh
```

Atau direct run:

```bash
go run . list
go run . add myguide
```

## License

MIT License

Copyright (c) 2026 Akujejo.
