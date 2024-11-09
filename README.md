# tonx_backend

## Quick Start

Clone from this repo, then run the command: 

```bash
docker build -t tonx_backend .
```

**Make sure you already build up:**
* [tonx_frontend docker image](https://github.com/tn00869679/tonx_frontend)
* tonx_backend docker image (this repo)

```bash
docker-compose up -d
```

View frontend at 
```txt
http://localhost:5173/
```

View API docs (Swagger) at
```txt
http://localhost:3310/
```
