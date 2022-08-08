FROM golang:1.18 as build-env
WORKDIR /app/
COPY . ./
RUN go mod download
RUN go get -d -v ./... 
RUN go vet -v ./...
RUN go test -v ./...
RUN CGO_ENABLED=0 go build -o devicemanager app/main.go
FROM gcr.io/distroless/static
LABEL "microservice.name"="Device Management"
COPY --from=build-env /app/devicemanager /
COPY --from=build-env /app/config.json /
CMD ["/devicemanager"]