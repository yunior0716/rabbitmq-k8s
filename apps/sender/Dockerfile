FROM golang:alpine as build

WORKDIR /src 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY sender.go ./

RUN go build sender.go

FROM alpine as runtime

COPY --from=build /src/sender /app/sender

CMD [ "/app/sender" ]

