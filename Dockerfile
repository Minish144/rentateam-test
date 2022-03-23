FROM golang:alpine3.15 AS build

WORKDIR /go/src/app

RUN apk add git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o api .

FROM scratch AS bin

COPY --from=build /go/src/app/api /api
# example config for tests
COPY --from=build /go/src/app/config.example.yaml /config.yaml

CMD ["/api", "runserver"]