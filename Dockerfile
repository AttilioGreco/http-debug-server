FROM golang:1.18-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /http-debug-server

FROM alpine:3.16
ARG USER=nonroot
ENV HTTP_PORT=":8080"
ENV HOME /home/$USER

# install sudo as root
RUN apk add sudo

# add new user
RUN adduser -D $USER \
        && echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER \
        && chmod 0440 /etc/sudoers.d/$USER

WORKDIR /

COPY --from=build /http-debug-server /http-debug-server

EXPOSE 8080

USER nonroot

ENTRYPOINT ["/http-debug-server"]