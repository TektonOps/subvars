FROM alpine:latest
ENTRYPOINT ["/usr/bin/subvars"]
COPY subvars /usr/bin/subvars