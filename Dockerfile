FROM golang:1.11.0 
ARG APP_NAME

RUN mkdir /app 
COPY src/${APP_NAME}.go /app/${APP_NAME}.go
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
