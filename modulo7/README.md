echo "TuPasswordSegura" | docker secret create db_password -

docker swarm init # si aún no está iniciado
docker stack deploy -c docker-stack.yml wordpress
docker stack ps wordpress
docker service ls
