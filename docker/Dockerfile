FROM golang:1.9.2 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 
EXPOSE 8080
# CMD ["/app/main"]

ENTRYPOINT ["/app/main"]