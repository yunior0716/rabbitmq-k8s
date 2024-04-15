FROM golang:alpine as build

WORKDIR /src 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY receiver.go ./

RUN go build receiver.go

FROM alpine as runtime

COPY --from=build /src/receiver /app/receiver

CMD [ "/app/receiver" ]



