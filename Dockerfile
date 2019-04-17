FROM golang:1.12 as build

ENV GO111MODULE=on

WORKDIR /go/src/vida

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go install -v ./...

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/bin/vida /
CMD ["/vida"]
