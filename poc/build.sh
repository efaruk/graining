CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/core.a ./core
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/modules.a ./modules
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/autoc .
docker build -t gopoc .
