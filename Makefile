all: api/google
	docker build --cache-from golang:1.22-alpine -t protos -f Dockerfile.proto
	docker run --pull=missing -v `pwd`:/app protos
	sudo chown $$USER:$$GROUP -R pkg
	go mod tidy
	
api/google:
	git clone --depth=1 https://github.com/googleapis/googleapis.git api/googleapis
	cp -r api/googleapis/google api
	rm -rf api/googleapis
	find api/google -type f -and -not -name '*.proto' -delete
