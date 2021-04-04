FROM golang:1.16.2-alpine3.13 as prepare

RUN apk add --update git build-base
WORKDIR /build

COPY go.* /build/
RUN go mod download -x


FROM node:15.12.0-alpine3.13 AS vue-prepare
COPY web/package* /web/
COPY web/yarn* /web/
WORKDIR /web

RUN yarn


FROM prepare AS go-build

COPY cmd /build/cmd
COPY internal /build/internal
WORKDIR /build

RUN go build gitlab.com/swinkelhofer/paperless-office/cmd/paperless-office


FROM vue-prepare AS vue-build

COPY web /web
COPY --from=vue-prepare web/* /web/
WORKDIR /web

RUN yarn app/build


FROM alpine:3.13

RUN apk add --update ocrmypdf tesseract-ocr tesseract-ocr-data-deu tesseract-ocr-data-fra tesseract-ocr-data-por tesseract-ocr-data-spa poppler-utils py3-pluggy py3-pip && \
    pip install coloredlogs && \
    apk del py3-pip && \
    chown 1000:1000 /srv

COPY build/docker/init /usr/bin/init
COPY --from=go-build /build/paperless-office /usr/bin/paperless-office
COPY --from=vue-build /web/index.html /html/
COPY --from=vue-build /web/dist /html/dist

USER 1000

WORKDIR /srv

EXPOSE 8000

VOLUME [ "/srv/data/processed", "/srv/data/raw" ]

ENTRYPOINT ["/usr/bin/init"]
CMD ["-serve-dir", "/srv"]
