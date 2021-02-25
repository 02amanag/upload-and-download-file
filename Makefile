dep:
	go get github.com/gin-gonic/gin

build:
	go build

run :
	./upload-and-download-file

docker:
	docker build -t upload-and-download-file .

