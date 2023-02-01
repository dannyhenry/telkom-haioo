# kode sebelum diperbaiki
# FROM golang
# ADD . /go/src/github.com/telkomdev/indihome/backend
# WORKDIR /go/src/github.com/telkomdev/indihome
# RUN go get github.com/tools/godep
# RUN godep restore
# RUN go install github.com/telkomdev/indihome
# ENTRYPOINT /go/bin/indihome
# LISTEN 80

FROM golang:alpine
ADD . /go/src/github.com/telkomdev/indihome/backend
WORKDIR /go/src/github.com/telkomdev/indihome
RUN go install github.com/tools/godep@latest
# RUN godep restore
RUN go mod init
RUN go mod tidy
RUN go get github.com/telkomdev/indihome
ENTRYPOINT /go/bin/indihome
EXPOSE 80

# error godep restore tidak ada di dalam direktory github.com/telkomdev/indihome
# dan repository github.com/telkomdev/indihome/backend tersebut tidak ada digithub