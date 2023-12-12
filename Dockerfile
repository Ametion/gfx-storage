FROM golang:1.20

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o /gfx-storage cmd/app/main.go

EXPOSE 5783

CMD [ "/gfx-storage" ]