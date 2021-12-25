FROM golang:1.15.7-buster
ADD main.go /go/src/main.go
ENTRYPOINT [ "go", "run", "/go/src/main.go" ]
CMD [ "9000", "echo container" ]
EXPOSE 9000