FROM golang:alpine as build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -v -o server

FROM alpine as production
COPY --from=build /app/server /usr/bin
COPY --from=build /app/.env /usr/bin
ENV RUN_ENV production
EXPOSE 8080
ENTRYPOINT [ "server" ]