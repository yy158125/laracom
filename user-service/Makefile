build:
	protoc -I. --go_out=plugins=micro:. proto/user/user.proto
run:
	#docker run -d -p 9092:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns -e DB_NAME=laracom_user -e DB_USER=root -e DB_PASSWORD=root --name laracom-user-service laracom-user-server


