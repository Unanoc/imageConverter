FROM golang:latest 
RUN mkdir /imageConverter 
ADD . /imageConverter/ 
WORKDIR /imageConverter 
RUN go build -o main . 
CMD ["/app/main"]