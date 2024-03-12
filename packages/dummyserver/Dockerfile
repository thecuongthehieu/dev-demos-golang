FROM golang:1.21
WORKDIR /dev-demos-golang
COPY . .
RUN pwd
RUN ls -l
# RUN go mod tidy
RUN go build -v -o /usr/local/bin/dummyserver ./packages/dummyserver
ENTRYPOINT ["dummyserver"]
