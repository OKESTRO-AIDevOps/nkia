
build: 

	go build -o ./ofront/ofront ./ofront/.

	go build -o ./osock/osock ./osock/.

run: 

	go run ./_okeygen/keygen.go

	cp okey osock/okey

	rm okey

	sudo docker-compose up --build
