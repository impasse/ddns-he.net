FROM golang:alpine AS builder

COPY . /build

RUN cd /build \
    && go build -o ddns-he.net .


FROM gcr.io/distroless/base

COPY --from=builder /build/ddns-he.net /ddns-he.net

ENTRYPOINT ["/ddns-he.net"]