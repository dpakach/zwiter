## Running using Docker

### 1. Build the docker image

```bash
cd /path/to/project
docker build -t zwiter .
```

### 2. Create a new Network
```bash
docker network create zwiter
```

### 3. Run the users service
```bash
docker run --rm -e service=users -p 8002:8002 --name zusers --net zwiter zwiter
```

### 4. Run the posts service
```
docker run --rm -e service=posts -p 8001:8001 --name zposts --net zwiter zwiter
```

