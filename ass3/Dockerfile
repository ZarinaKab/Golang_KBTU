FROM golang:latest


WORKDIR /app
COPY ./ ./

EXPOSE 8080


RUN go mod download
RUN go build -o ass3 ./main.go

CMD [ "./ass3" ]
