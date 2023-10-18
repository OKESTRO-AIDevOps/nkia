
build: 

	go build -o ./ofront/ofront ./ofront/.

	go build -o ./osock/osock ./osock/.

run: build

	go run ./_okeygen/keygen.go

	cp okey osock/okey

	rm okey

	sudo docker-compose up --build
