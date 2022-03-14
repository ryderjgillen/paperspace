# syntax=docker/dockerfile:labs
ARG GOLANG_VERSION
FROM golang:${GOLANG_VERSION}-alpine3.15

RUN apk add git g++ docker make

COPY <<EOF /usr/local/bin/update.sh
#!/bin/sh
echo "Updated Tools: `date`" >> /usr/local/log.txt
go install golang.org/x/tools/gopls@latest
go install github.com/cweill/gotests/gotests@latest
go install github.com/fatih/gomodifytags@latest 
go install github.com/haya14busa/goplay/cmd/goplay@latest
go install github.com/josharian/impl@latest
go install github.com/go-delve/delve/cmd/dlv@2f13672765fe && mv /go/bin/dlv /go/bin/dlv-dap
go install github.com/go-delve/delve/cmd/dlv@latest
go install github.com/ramya-rao-a/go-outline@latest
go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
EOF

RUN chmod +x /usr/local/bin/update.sh

RUN \
    go install golang.org/x/tools/gopls@latest &&\
    go install github.com/cweill/gotests/gotests@latest &&\
    go install github.com/fatih/gomodifytags@latest &&\
    go install github.com/haya14busa/goplay/cmd/goplay@latest &&\
    go install github.com/josharian/impl@latest &&\
    go install github.com/go-delve/delve/cmd/dlv@2f13672765fe && mv /go/bin/dlv /go/bin/dlv-dap &&\
    go install github.com/go-delve/delve/cmd/dlv@latest &&\
    go install github.com/ramya-rao-a/go-outline@latest &&\
    go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest &&\
    go install honnef.co/go/tools/cmd/staticcheck@latest