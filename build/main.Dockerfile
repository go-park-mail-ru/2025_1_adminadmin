FROM golang:1.24.2-alpine AS builder

COPY . /github.com/go-park-mail-ru/2025_1_adminadmin
WORKDIR /github.com/go-park-mail-ru/2025_1_adminadmin

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -o ./.bin ./cmd/main/main.go
RUN go clean --modcache

FROM scratch AS runner

WORKDIR /build_v1/

COPY --from=builder /github.com/go-park-mail-ru/2025_1_adminadmin/.bin .

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /
ENV TZ="Europe/Moscow"
ENV ZONEINFO=/zoneinfo.zip

EXPOSE 5458
EXPOSE 5459

ENTRYPOINT ["./.bin"]