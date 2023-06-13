build:
	go build -o bin/cnmotd *.go

run: build
	./bin/cnmotd --debug

run-validate: build
	./bin/cnmotd --debug --validate

image:
	docker build -t jeefy/cnmotd .

image-push: image
	docker push jeefy/cnmotd

image-run: image
	docker rm -f cnmotd || true
	docker run --name=cnmotd -d -p 8080:8080 jeefy/cnmotd:latest