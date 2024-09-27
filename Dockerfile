FROM golang:1.23 AS build

WORKDIR /knockbox/src
COPY . .

ARG GITHUB_TOKEN
ENV CGO_ENABLED=0 GOPRIVATE=github.com/knockbox
RUN git config --global url."https://oauth2:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"

RUN go mod download
RUN go build -o /knockbox/bin/main

FROM gcr.io/distroless/static-debian12

COPY --from=build /knockbox/bin /

EXPOSE 9090

ENTRYPOINT ["/main"]