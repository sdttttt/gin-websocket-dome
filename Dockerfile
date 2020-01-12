FROM golang:1.13

LABEL author='sdttttt' email='760159537@qq.com' description='This is image of project'

WORKDIR $GOPATH/src/gin-web
ADD . $GOPATH/src/gin-web

RUN go build .
EXPOSE 10086

ENTRYPOINT ["./gin-web"]