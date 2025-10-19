docker compose up # Levanta todo
docker compose up -d # En segundo plano
docker compose ps # Ver contenedores
docker compose logs # Logs de todos
docker compose logs users-api
docker compose down # Detiene y elimina contenedores y red
docker compose down -v # También borra volúmenes
