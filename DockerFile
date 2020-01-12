FROM golang:1.13

LABEL author='sdttttt' email='760159537@qq.com' description='This is image of project'

# TODO: Haven't Test
RUN go get

RUN go run main.go