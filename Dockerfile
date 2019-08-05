FROM golang:1.10.1
WORKDIR /go/src/app
ADD . /go/src/app/
CMD ["go", "get", "github.com/gin-contrib/static"]
CMD ["go", "get", "github.com/gin-gonic/gin"] 
CMD ["go", "get", "github.com/russross/blackfriday"]
RUN ls -altr /go/src/app
ENV PORT=8080
ENV GOPATH=/go
CMD ["go", "run", "myMD.go"]
