FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN CGO_ENABLED=0 go build -o /out/helloworld .

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/helloworld /helloworld
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/helloworld"]
