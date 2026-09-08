# Podman

Daemonless container runtime. Rootless containers by default.

## Basic

### podman run
**Deskripsi:** Run container.
```bash
podman run hello-world
podman run -d nginx
podman run -p 8080:80 nginx
podman run -v /local:/container nginx
```

### podman ps
**Deskripsi:** List running containers.
```bash
podman ps
podman ps -a
```

### podman stop/start/restart
**Deskripsi:** Control containers.
```bash
podman stop CONTAINER
podman start CONTAINER
podman restart CONTAINER
```

### podman rm
**Deskripsi:** Remove container.
```bash
podman rm CONTAINER
podman rm -f CONTAINER
```

---

## Images

### podman images
**Deskripsi:** List images.
```bash
podman images
podman images -a
```

### podman pull
**Deskripsi:** Download image.
```bash
podman pull nginx
podman pull docker.io/library/nginx
```

### podman rmi
**Deskripsi:** Remove image.
```bash
podman rmi IMAGE
podman rmi $(podman images -q)  # force
```

---

## Rootless

### podman login
**Deskripsi:** Login to registry.
```bash
podman login docker.io
```

### podman system
**Deskripsi:** Rootless daemon.
```bash
podman system service &
podman ps  # rootless by default
```

---

## Volumes

### podman volume
**Deskripsi:** Volume commands.
```bash
podman volume create myvol
podman volume ls
podman volume inspect myvol
podman volume rm myvol
```

---

## Rootless daemon mode
```bash
# Background daemon (no root)
podman system service --time=0

# Use via systemd socket activation
systemctl --user start podman.socket
```

---

## Kube Compatibility

Podman can run as k3s/containerd backend:
```bash
podman kube play -f deployment.yaml
podman kube run -f pod.yaml
```

---

## Best Practices

- Rootless by default (safer, no root required)
- Use `podman system service` for background daemon
- Similar syntax to Docker (interchangeable where supported)
- No daemon process running continuously (daemonless architecture)
