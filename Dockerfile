## compile css from sass file
FROM jbergknoff/sass:latest as sasscompiler
WORKDIR /tmp/
COPY static/main.scss main.scss
RUN sass main.scss main.css


## build binary withh go
FROM golang:latest as builder
WORKDIR /go/src/app
COPY *.go .
# remove debug information
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build ./


## building minimal image from scratch containing only the binary and static files
FROM scratch
WORKDIR /bin/

# copy static files
COPY static/ static/
COPY --from=sasscompiler /tmp/main.css static/main.css

# copy binary
COPY --from=builder /go/src/app/app app
CMD ["/bin/app"]