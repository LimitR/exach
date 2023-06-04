FROM golang:alpine AS builder

WORKDIR /build
ADD .env /build/
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

FROM alpine

WORKDIR /

COPY --from=builder /build .
CMD [ "./main" ]

FROM node:16-alpine as builder_frontend

WORKDIR /app

COPY ./frontend .

RUN npm i 

RUN npm run build


FROM nginx:alpine

#!/bin/sh

COPY ./config/nginx.conf /etc/nginx/nginx.conf

## Remove default nginx index page
RUN rm -rf /usr/share/nginx/html/*

# Copy from the stahg 1
COPY --from=builder_frontend /app/build /usr/share/nginx/html

ENTRYPOINT ["nginx", "-g", "daemon off;"]