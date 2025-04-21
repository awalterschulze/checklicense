# build

FROM golang:1.24 AS builder

WORKDIR /go/src/awalterschulze/checklicense

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" .

# run

FROM alpine:3 AS production
LABEL maintainer="Walter Schulze <awalterschulze@users.noreply.github.com>"

COPY --from=builder /go/src/awalterschulze/checklicense/checklicense /usr/bin/checklicense

ENTRYPOINT ["/usr/bin/checklicense"]
