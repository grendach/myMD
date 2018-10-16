FROM golang:1.10.1
ADD ./src /go/src/app
CMD ["go", "get", "github.com/gin-contrib/static"]
CMD ["go", "get", "github.com/gin-gonic/gin"] 
CMD ["go", "get", "github.com/russross/blackfriday"]
WORKDIR /go/src/app
ENV PORT=3001
ENV GOPATH=/go
CMD ["go", "run", "myMD.go"]
