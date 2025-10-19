# Almacenamiento:

- Montar una carpeta local

```bash
docker run -d --name miweb \
  -p 8080:80 \
  -v $(pwd):/usr/share/nginx/html \
  nginx
```

- Crear un volumen

```bash
docker volume create webdata
```

- Usar el Volumen

```bash
docker run -d --name miweb \
  -p 8080:80 \
  -v webdata:/usr/share/nginx/html \
  nginx
```
