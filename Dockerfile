#FROM cimg/go:1.23

FROM golang:1.23

WORKDIR /app/

COPY . .

RUN go mod download

#ENTRYPOINT ["go", "test", "./..."]
#CMD ["go", "test", "./...", "&&", "ls", "-al"]
CMD ["./run_test.sh"]
