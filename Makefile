build:
	go build -o bin/cnmotd *.go

run: build
	./bin/cnmotd --debug

image:
	docker build -t jeefy/cnmotd .

image-push:
	docker push jeefy/cnmotd