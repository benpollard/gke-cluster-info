## BUILD
FROM golang:1.19.2-alpine3.16 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o gke-cluster-info .

## DEPLOY
FROM alpine:3.16

COPY --from=build /app/gke-cluster-info /gke-cluster-info

EXPOSE 8080

ENTRYPOINT [ "/gke-cluster-info" ]
