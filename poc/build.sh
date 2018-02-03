CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/assets/bin/core.a ./core
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/assets/bin/modules.a ./modules
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o docker/assets/bin/autoc .
