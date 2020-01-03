build:
	protoc -I. --go_out=plugins=micro:. proto/demo/demo.proto
	GOOS=linux GOARCH=amd64 go build
	docker build -t laracom-demo-server .

run:
	docker run -d -p 9091:9999 -p 8081:8000 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns  laracom-demo-server