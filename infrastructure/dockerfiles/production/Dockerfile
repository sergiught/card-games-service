FROM golang:1.22 AS api-builder

WORKDIR /card-games-service

ENV GOFLAGS="-mod=vendor"

COPY . .

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 VERSION="${VERSION}" make build

FROM ubuntu:noble AS final-image

RUN groupadd dealer && useradd -g dealer dealer

USER dealer

COPY --from=api-builder /card-games-service/dist/card-games-service /

EXPOSE 8080

CMD ["/card-games-service"]
