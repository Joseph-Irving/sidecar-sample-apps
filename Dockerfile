FROM scratch

ADD bin/sample-linux-amd64 sample

ENTRYPOINT ["/sample"]
