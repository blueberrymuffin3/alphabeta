FROM golang:1.16.3-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o alphabeta

FROM scratch
COPY --from=build /app/alphabeta /alphabeta

CMD [ "/alphabeta" ]
