FROM golang:1.20.5-alpine3.18

RUN apk update && apk add git

RUN go install -v golang.org/x/tools/gopls@lates \
                  golang.org/x/tools/cmd/goimports@latest