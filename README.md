zgrab
==================

A TLS Banner Grabber, in Go

## Docker 

### Building
    docker build . -t baykovr/zgrab:latest
    docker-compose build

### Usage

This dockerized version relies on an input file rather than piping over stdin. We expose /var/lib/zgrab as a possible volume mount point for input file storage, a sample file is included in zinput/host.sample ; example invocations follow below.

    docker run -v $GOPATH/src/github.com/baykovr/zgrab/zinput:/var/lib/zgrab baykovr/zgrab:latest --tls --port 443 --http="/" -input-file /var/lib/zgrab/host.sample
        
Alternative, modify the command section of docker-compose, invoke
     
     docker-compose up
    
    
## Building Local

You will need to have a valid `$GOPATH` set up, for more information about `$GOPATH`, see https://golang.org/doc/code.html.

Once you have a working `$GOPATH`, run:

```
go get github.com/zmap/zgrab
```

This will install zgrab under `$GOPATH/src/github.com/zmap/zgrab`

```
$ cd $GOPATH/src/github.com/zmap/zgrab
$ go build
```

## Usage

```
./zgrab --help
```

## Example

```
$ zmap -p 443 --output-fields=* | ztee results.csv | zgrab --port 443 --tls --http="/" --output-file=banners.json
```

## Requirements

zgrab requires go version of at least 1.6. Please note that this is newer than the version included in Ubuntu 14.04 apt repository. You can install ztee from ZMap Github repository at https://github.com/zmap/zmap.


## ZGrab as a library / dependency

If you are using ZGrab code in another Go program, import ZGrab using [gopkg.in](http://gopkg.in). ZGrab tends to be very unstable, API's may break at any time, so be sure to vendor ZGrab.

## License

ZGrab is licensed under Apache 2.0 and ISC. For more information, see the LICENSE file.
