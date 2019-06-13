FROM node:11.15.0 AS node-builder

WORKDIR /work
COPY ./client /work/client
COPY ./package.json /work/
COPY ./package-lock.json /work/
COPY ./config/webpack.config.js /work/config/

RUN npm install
RUN ./node_modules/webpack/bin/webpack.js --config config/webpack.config.js

FROM golang:latest AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 0

WORKDIR /go/src/github.com/MrHuxu/react-go-ssr-boilerplate
COPY ./main.go /go/src/github.com/MrHuxu/react-go-ssr-boilerplate/
COPY ./server /go/src/github.com/MrHuxu/react-go-ssr-boilerplate/server
COPY ./go.mod /go/src/github.com/MrHuxu/react-go-ssr-boilerplate/
COPY ./go.sum /go/src/github.com/MrHuxu/react-go-ssr-boilerplate/

RUN go mod download
RUN go build main.go

FROM scratch

ENV GIN_MODE release
ENV INSIDE_DOCKER true

WORKDIR /output
COPY ./config/server.json /output/config/
COPY ./server/templates /output/server/templates
COPY --from=node-builder /work/client/public/bundle.js /output/client/public/
COPY --from=go-builder /go/src/github.com/MrHuxu/react-go-ssr-boilerplate/main /output/

EXPOSE 13109
ENTRYPOINT [ "./main" ]