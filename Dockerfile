FROM golang:1.8

RUN go get github.com/baykovr/zgrab && \ 
go build github.com/baykovr/zgrab

VOLUME ["/var/lib/zgrab"]

ENTRYPOINT ["/go/bin/zgrab"]
CMD ["--help"]
