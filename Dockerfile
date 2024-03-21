FROM golang:1.21 AS build

WORKDIR /go/src
COPY . ./go
COPY ./app.go .

COPY ./go.mod .
#RUN go mod tidy
#RUN go mod vendor

RUN go mod download
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64 

#RUN go get -d -v ./...
RUN go build -a -installsuffix cgo -o app ./go 

FROM scratch AS runtime
COPY --from=build /go/src/mware ./go
EXPOSE 8080/tcp
ENTRYPOINT ["./mware"]


