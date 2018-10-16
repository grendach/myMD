FROM golang:alpine
ADD ./myMD /go/src/app
WORKDIR /go/src/app
ENV PORT=3001
CMD ["go", "run", "myMD.go"]
