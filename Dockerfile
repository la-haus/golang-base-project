FROM 521664337176.dkr.ecr.us-east-1.amazonaws.com/go-golden:1.0 AS builder
WORKDIR /go/src/github.com/la-haus/sample
COPY ./go.* ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -a -tags musl -installsuffix cgo -o app .

FROM 521664337176.dkr.ecr.us-east-1.amazonaws.com/alpine-golden:1.0
WORKDIR /usr/
COPY --from=builder /go/src/github.com/la-haus/sample/app .
EXPOSE 3000

CMD /usr/app
