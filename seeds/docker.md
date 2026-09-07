# Docker

Container runtime. Deploy apps in isolated environments.

## Basic

### docker run
**Deskripsi:** Run container.
```bash
docker run hello-world
docker run -d nginx
docker run -p 8080:80 nginx
docker run -v /local:/container nginx
```

### docker ps
**Deskripsi:** List running containers.
```bash
docker ps
docker ps -a
docker ps --format "table {{.Names}}\t{{.Status}}\t{{.Ports}}"
```

### docker stop/start/restart
**Deskripsi:** Control containers.
```bash
docker stop CONTAINER
docker start CONTAINER
docker restart CONTAINER
```

### docker rm
**Deskripsi:** Remove container.
```bash
docker rm CONTAINER
docker rm -f CONTAINER  # force kill first
docker rm $(docker ps -aq)  # remove all
```

---

## Images

### docker images
**Deskripsi:** List images.
```bash
docker images
docker images -a
docker images --filter "dangling=true"
```

### docker pull
**Deskripsi:** Download image.
```bash
docker pull nginx
docker pull ubuntu:22.04
```

### docker rmi
**Deskripsi:** Remove image.
```bash
docker rmi IMAGE
docker rmi $(docker images -q)  # force
```

---

## Logs

### docker logs
**Deskripsi:** View logs.
```bash
docker logs CONTAINER
docker logs -f CONTAINER  # follow
docker logs --tail 50 CONTAINER
```

### docker inspect
**Deskripsi:** Inspect container/image.
```bash
docker inspect CONTAINER
docker inspect --format '{{.NetworkSettings.IPAddress}}' CONTAINER
```

---

## Networks

### docker network ls
**Deskripsi:** List networks.
```bash
docker network ls
docker network inspect bridge
```

### docker network create
**Deskripsi:** Create network.
```bash
docker network create mynet
docker network connect mynet CONTAINER
```

---

## Volumes

### docker volume ls
**Deskripsi:** List volumes.
```bash
docker volume ls
docker volume inspect myvolume
```

### docker volume create
**Deskripsi:** Create volume.
```bash
docker volume create myvolume
docker run -v myvolume:/data nginx
```

---

## Compose

### docker-compose.yml
```yaml
version: '3.8'
services:
  web:
    image: nginx
    ports:
      - "8080:80"
    volumes:
      - ./html:/usr/share/nginx/html
  db:
    image: postgres:15
    environment:
      POSTGRES_PASSWORD: example
```

### docker-compose
**Deskripsi:** Compose commands.
```bash
docker-compose up -d
docker-compose down
docker-compose logs -f
docker-compose ps
```

---

## Best Practices

- Use minimal base image: `alpine` or `distroless`
- Keep containers ephemeral: no persistence inside container
- Use explicit volumes for data
- Tag images: `docker tag myimg:v1.0 myregistry/myimg:v1.0`
- Scan for vulnerabilities: `docker scan IMAGE`
