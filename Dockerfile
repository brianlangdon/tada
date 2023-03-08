FROM node:18.14.2 AS JS_BUILD
COPY dating /dating
WORKDIR dating
RUN npm install && npm run build --prod

FROM golang:1.19.5-alpine AS GO_BUILD
COPY server /server
WORKDIR /server
RUN apk add build-base
RUN go build -o /go/bin/server

FROM alpine:3.14
COPY --from=JS_BUILD /dating/build* ./dating/
COPY --from=GO_BUILD /go/bin/server ./
CMD ./server
