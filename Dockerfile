FROM golang:1.22.5-alpine AS golang

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . ./

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /go-rest-template

FROM scratch

# Will only work on linux
COPY --from=golang /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=golang /app/.env .
COPY --from=golang /app/config/migrations /config/migrations/
# COPY --from=golang /app/assets /assets/


COPY --from=golang /go-rest-template .

CMD ["/go-rest-template"]