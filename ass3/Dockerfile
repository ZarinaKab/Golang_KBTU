FROM golang:latest


WORKDIR /app
COPY ./ ./

EXPOSE 8080


RUN go mod download
RUN go build -o golang_ass_2 ./main.go

CMD [ "./golang_ass_2" ]
