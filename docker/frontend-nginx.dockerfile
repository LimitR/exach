FROM node:16-alpine as builder_frontend

WORKDIR /app

COPY ./frontend .

RUN npm i 

RUN npm run build


FROM nginx:alpine

COPY ./config/nginx.conf /etc/nginx/nginx.conf

RUN rm -rf /usr/share/nginx/html/*

COPY --from=builder_frontend /app/build /usr/share/nginx/html

ENTRYPOINT ["nginx", "-g", "daemon off;"]