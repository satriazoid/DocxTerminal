# SSH

Secure shell. Remote access, port forwarding, tunneling.

## Basic

### ssh
**Deskripsi:** Connect to remote host.
```bash
ssh user@host
ssh user@192.168.1.100
ssh -p 2222 user@host
```

### ssh-keygen
**Deskripsi:** Generate SSH key pair.
```bash
ssh-keygen -t ed25519 -C "user@email.com"
ssh-keygen -t rsa -b 4096
```

### ssh-copy-id
**Deskripsi:** Copy public key to remote.
```bash
ssh-copy-id user@host
ssh-copy-id -i ~/.ssh/id_ed25519.pub user@host
```

---

## Keys

### ssh-add
**Deskripsi:** Add key to agent.
```bash
ssh-add ~/.ssh/id_ed25519
ssh-add -l  # list keys
ssh-add -D  # delete all
```

### Authorized Keys
**Deskripsi:** Manage authorized_keys file.
```bash
cat ~/.ssh/authorized_keys
echo "ssh-ed25519 AAAA... key@example.com" >> ~/.ssh/authorized_keys
```

---

## Config

### ~/.ssh/config
```ssh
Host myserver
    HostName 192.168.1.100
    User admin
    Port 2222
    IdentityFile ~/.ssh/id_ed25519

Host github.com
    HostName github.com
    User git
    IdentityFile ~/.ssh/id_ed25519_github
```

### ssh -F
**Deskripsi:** Use alternative config file.
```bash
ssh -F ~/.ssh/config.custom user@host
```

---

## Tunneling

### SSH Forwarding
```bash
# Local forward (localhost:localPort -> remote:host:port)
ssh -L 8080:localhost:80 user@remote
ssh -L 3307:db.example.com:3306 user@db

# Remote forward (remote:host:port -> localhost:localPort)
ssh -R 8080:localhost:80 user@remote

# Dynamic SOCKS proxy
ssh -D 1080 user@proxy
```

### SSH ProxyJump
```bash
# Jump through intermediate host
ssh -J jumphost user@target
ssh -J user@jumphost:user@target host
```

### SSH ProxyCommand
```bash
# ~/.ssh/config
Host target
    ProxyCommand ssh -W %h:%p jumphost
```

---

## Options

### ssh -t
**Deskripsi:** Request pseudo-terminal allocation (useful for interactive apps).
```bash
ssh -t user@host
ssh -t -T  # disable pty allocation
```

### ssh -N
**Deskripsi:** Don't execute remote command (for forwarding only).
```bash
ssh -N -L 8080:localhost:80 user@remote
ssh -N -D 1080 user@proxy
```

### ssh -C
**Deskripsi:** Enable compression.
```bash
ssh -C user@slow-connection
```

### ssh -o
**Deskripsi:** Set options.
```bash
ssh -o StrictHostKeyChecking=no user@host
ssh -o ServerAliveInterval=60 user@host
ssh -o ConnectTimeout=10 user@host
```

---

## Security

### ~/.ssh/config security
```ssh
Host *
    ServerAliveInterval 60
    ServerAliveCountMax 3
    TCPKeepAlive yes
    IdentitiesOnly yes
    ForwardAgent no
```

### Disable password auth
```bash
# On server: /etc/ssh/sshd_config
PasswordAuthentication no
PubkeyAuthentication yes

# Restart SSH
sudo systemctl restart sshd
```

### Limit user access
```bash
# /etc/ssh/sshd_config
AllowUsers admin
AllowUsers user1 user2
DenyUsers attacker
```

---

## Troubleshooting

### ssh -v
**Deskripsi:** Verbose mode.
```bash
ssh -v user@host
ssh -vvv user@host  # extra verbose
```

### Connection refused
```bash
# Check if SSH is running
sudo systemctl status sshd
sudo systemctl start sshd

# Check firewall
sudo ufw status
sudo firewall-cmd --list-all
```

### Permission denied
```bash
# Check permissions
ls -la ~/.ssh/id_*
chmod 600 ~/.ssh/id_*
chmod 644 ~/.ssh/id_*.pub
chmod 700 ~/.ssh/
chmod 600 ~/.ssh/authorized_keys
```

---

## Best Practices

- Use key-based authentication (no passwords)
- Limit user permissions in SSH config
- Disable password auth for remote servers
- Set SSH key passphrase for sensitive keys
- Rotate SSH keys periodically
- Monitor `~/.ssh/known_hosts` for changes
- Use `StrictHostKeyChecking` to prevent MITM
