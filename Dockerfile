# build

FROM golang:1.24 AS builder

WORKDIR /go/src/awalterschulze/checklicense

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

# run

FROM alpine:3 AS production
LABEL maintainer="Walter Schulze <awalterschulze@users.noreply.github.com>"

COPY --from=builder /go/src/awalterschulze/checklicense/checklicense /usr/bin/checklicense

WORKDIR /workdir

RUN set -eux; \
  addgroup -g 1000 checklicense; \
  adduser -u 1000 -G checklicense -s /bin/sh -h /home/checklicense -D checklicense

RUN chown -R checklicense:checklicense /workdir

USER checklicense

ENTRYPOINT ["/usr/bin/checklicense"]
