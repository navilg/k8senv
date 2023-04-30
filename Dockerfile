FROM golang:1.20.3-alpine3.17 as build
ARG OS
ARG ARCH
WORKDIR /build
COPY . .
RUN apk add git
RUN go mod download && \
    CGO_ENABLED=0 go build -o k8senv

FROM alpine:3.17
ARG VERSION
ARG user=k8senv
ARG group=k8senv
ARG uid=1000
ARG gid=1000
USER root
WORKDIR /app
COPY --from=build /build/k8senv /app/.k8senv/bin/k8senv
RUN apk update && apk --no-cache add bash vim && addgroup -g ${gid} ${group} && adduser -h /app -u ${uid} -G ${group} -s /bin/bash -D ${user}
RUN chown -R k8senv:k8senv /app/.k8senv && chmod -R u+rx /app/.k8senv
USER k8senv
ENV PATH="/app/.k8senv/bin:$PATH"
VOLUME /app/.k8senv/
ENTRYPOINT [ "k8senv"]
CMD ["help"]