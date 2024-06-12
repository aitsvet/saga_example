FROM golang:alpine AS builder

RUN apk add --no-cache --update gcc musl-dev

WORKDIR /src

COPY . .

ENV CGO_ENABLED=1

RUN go mod download

FROM builder AS apps

ENV GOBUILD='go build -ldflags=-s -ldflags=-w -trimpath -tags musl -a -o'

RUN $GOBUILD /dst/gateway ./cmd/gateway/main.go
RUN $GOBUILD /dst/ad ./cmd/ad/main.go
RUN $GOBUILD /dst/order ./cmd/order/main.go
RUN $GOBUILD /dst/shipment ./cmd/shipment/main.go
RUN $GOBUILD /dst/dictionary ./cmd/dictionary/main.go

RUN ldd /dst/gateway | tr -s [:blank:] '\n' | grep ^/ | xargs -I % install -D % /dst/%
RUN ldd /dst/ad | tr -s [:blank:] '\n' | grep ^/ | xargs -I % install -D % /dst/%

FROM alpine:latest AS production

COPY --from=apps /dst /

EXPOSE 8080
