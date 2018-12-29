FROM jbergknoff/sass:latest
WORKDIR /tmp/
COPY static/main.scss main.scss
RUN sass main.scss main.css

FROM golang:latest

WORKDIR /go/src/app
COPY static/ static/
COPY --from=0 /tmp/main.css static/main.css
COPY *.go .
RUN go install ./

CMD ["app"]