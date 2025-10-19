FROM nginx:1.27-alpine

# Opcional: limpiar la página por defecto
RUN rm -f /usr/share/nginx/html/index.html

# Copiar nuestro index.html
COPY index.html /usr/share/nginx/html/index.html

EXPOSE 80

# Mantener NGINX en primer plano
CMD ["nginx", "-g", "daemon off;"]
