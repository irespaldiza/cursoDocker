# Comandos Basicos

- Ejecutar un contenedor en primer plano:

```bash
docker run -it ubuntu bash
```

- Ejecutar un contenedor en segundo plano

```bash
docker run -d --name web  nginx
```

- Listar los contenedores en ejecucion

```bash
docker ps
```

- Mostrar los logs logs de un contenedor

```bash
docker logs web
```

- Ejecutar un nuevo comando en el contenedor

```bash
docker exec -it web bash
```

| No es el comando principal

- Parar un Contenedor

```bash
docker stop we-b
```

- Arrancar un Contenedor parado

```bash
docker start web
```

- Eliminar un contenedor

````bash
docker rm web
```

# Practica:
Comprobar si lo:wq
s cambios realizados dentro de un contenedor (por ejemplo, crear un archivo) se mantienen cuando el contenedor se detiene, se reinicia o se elimina.
````
