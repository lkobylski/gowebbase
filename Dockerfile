FROM golang:alpine as build

WORKDIR /go/src/app
COPY . .
RUN go install -v ./...

FROM alpine
COPY --from=build /go/src/app/templates /templates
COPY --from=build /go/src/app/public /public
COPY --from=build /go/bin/gowebbase /usr/bin/gowebbase
#COPY --from=build /go/src/app/application /application

ENV DATABASE_URL=postgres://postgres:portgres@postgres/postgres

RUN apk add --no-cache curl
#HEALTCHECK CMD todo

EXPOSE 3000

CMD ["gowebbase"]

