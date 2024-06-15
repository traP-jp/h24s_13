FROM golang:1-alpine AS builder

WORKDIR /work

COPY ./go.* .
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    go build -o /work/app -ldflags "-s -w" .

FROM alpine:3 AS prod

WORKDIR /work

COPY --from=builder /work/app .

ENTRYPOINT ["/work/app"]
