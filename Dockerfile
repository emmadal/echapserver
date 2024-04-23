FROM go:1.22.2-bookworm

LABEL Author="Emmanuel Dalougou"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 4000

CMD ["/docker-gs-ping"]