FROM golang:1.20.5-alpine AS build
WORKDIR /go/src/kronologos
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/kronologos ./cmd/kronologos

FROM scratch
COPY --from=build /go
