# Ejemplos

## Puertos

- Expone un puerto en Docker

```bash
docker run -d --name web -p 8080:80 nginx
```

## Ejercicio:

- Corre un contenedor de nginx(html), users-api y orders-api en los que se pueda hacer una peticion http.

## Redes

````bash
docker inspect nginx-web --format='{{json .NetworkSettings.Networks}}' | jq
docker inspect users-api --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
docker inspect orders-api --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}'
```
```bash
docker network ls
```

## Crear una red personalizada
```bash
docker network create app_network
docker run -d --name users-api \
  --network app_net \
  -p 8081:8081 users-api:dev
docker run -d --name orders-api \
  --network app_net \
  -p 8082:8082 orders-api:dev
```

- Ejercicio: Levantar el container curl en la misma red y comprobar la conectividad, tanto por resolucion de nombres como por ip.
````
