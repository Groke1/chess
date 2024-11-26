FROM golang:1.23

WORKDIR src/engine/

RUN apt update && \
    apt-get install -y stockfish

COPY . .

EXPOSE 8080

ENTRYPOINT ["go", "run", "cmd/main.go"]