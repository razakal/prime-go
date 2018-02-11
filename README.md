# prime-go

## Build app on Linux/Mac
```bash
env GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -i -o=bin/prime/prime src/prime/main.go
```

## Build app on Windows
```bash
$env:GOOS="linux" 
$env:GOARCH="amd64" 
go build -ldflags "-s -w" -i -o=bin/prime/prime src/prime/main.go
```

## Build Docker image
```bash
docker build -f Dockerfile.prime -t razakal/prime:prime-v1 .
```

## Test Docker image
```bash
docker run -p 8080:8080 razakal/prime:prime-v1
docker kill <container id>
```

## Upload Docker image to repository
```bash
docker login
docker push razakal/prime:prime-v1
```