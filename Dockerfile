FROM golang:1.10

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]

# $ docker build -t my-golang-app .
# $ docker run -it --rm --name my-running-app my-golang-app