FROM golang:alpine as build

WORKDIR /src 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY saver.go ./

RUN go build saver.go

FROM alpine as runtime

COPY --from=build /src/saver /app/saver

CMD [ "/app/saver" ]
