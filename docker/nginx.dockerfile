FROM nginx:alpine

COPY ./config/nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder_frontend /app/build /usr/share/nginx/html

ENTRYPOINT ["nginx", "-g", "daemon off;"]