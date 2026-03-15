start:
    podman compose up -d

stop:
    podman compose stop

run:
    podman exec svg-guard go run .

add package:
    podman exec svg-guard go get {{package}}

clean:
    podman exec svg-guard go mod tidy

format:
    podman exec svg-guard go fmt ./...