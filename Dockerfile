FROM golang:1.21-alpine

# membuat direktori folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY . .

RUN go mod tidy

# create executable
RUN go build -o beapi ./cmd/httpserver

EXPOSE 8080

CMD [ "./beapi" ]