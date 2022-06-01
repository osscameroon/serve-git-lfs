build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sglfs

run:
	go run main.go

run-prod:
	./sglfs

docker-build:
	docker build -t sanix-darker/sglfs:latest -f Dockerfile .

docker-build-no-cache:
	docker build --no-cache -t sanix-darker/sglfs:latest -f Dockerfile .

docker-run:
	docker run -it --rm -p 3000:3000 -v ${PWD}/shared:/shared sanix-darker/sglfs:latest
