##
## Build
##
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /demo-cloud-yncrea

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /demo-cloud-yncrea /demo-cloud-yncrea

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/demo-cloud-yncrea"]
