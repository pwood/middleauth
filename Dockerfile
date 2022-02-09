ARG GO_VERSION=1.18-rc

FROM golang:${GO_VERSION}-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
RUN apk add --no-cache ca-certificates git

WORKDIR /src
COPY ./ ./

ENV CGO_ENABLED=0
RUN go build -o /middleauth .

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /middleauth /middleauth

USER nobody:nobody

ENTRYPOINT ["/middleauth"]
