FROM golang:1.16.3 AS builder
LABEL maintainer=SnowWarrior email=gerrytranCHINA@gmail.com

WORKDIR /build
RUN useradd -u 10001 -ms /bin/bash app-runner

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE="on"

COPY ./ /build 

RUN go mod download \
  && CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o /build/ri_ji /build/main.go

FROM alpine:3.13.4 AS final

WORKDIR /app

COPY --from=builder /build/ri_ji /app
COPY --from=builder /build/.env /app
COPY --from=builder /build/key /app/key

USER app-runner
ENTRYPOINT [ "/app/ri_ji" ]
