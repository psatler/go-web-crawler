FROM golang:latest

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

# $ docker build -t my-golang-app .
# $ docker run -it --rm --name my-running-app my-golang-app