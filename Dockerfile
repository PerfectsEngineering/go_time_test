FROM cimg/go:1.23

WORKDIR /app/

COPY . .

RUN go mod download

ENTRYPOINT ["go", "test", "./..."]
