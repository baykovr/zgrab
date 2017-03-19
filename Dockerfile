FROM golang 

#ADD . /go/src/github.com/baykovr/zgrab 
RUN go get github.com/baykovr/zgrab && \ 
go build github.com/baykovr/zgrab

VOLUME ["/var/lib/zgrab"]

ENTRYPOINT ["/go/bin/zgrab"]
CMD ["--help"]
