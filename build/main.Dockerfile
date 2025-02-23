FROM golang:1.24.0-alpine AS builder

COPY . /github.com/go-park-mail-ru/2025_1_adminadmin
WORKDIR /github.com/go-park-mail-ru/2025_1_adminadmin

RUN go mod download
RUN go clean --modcache
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/main.go

FROM scratch AS runner

WORKDIR /build_v1/

COPY --from=builder /github.com/go-park-mail-ru/2025_1_adminadmin/.bin .

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ="Europe/Moscow"
ENV ZONEINFO=/zoneinfo.zip

EXPOSE 5458

ENTRYPOINT ["./.bin"]