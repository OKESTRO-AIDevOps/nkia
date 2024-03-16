
build: 

	go build -o ./ofront/ofront ./ofront/.

	go build -o ./osock/osock ./osock/.

run: 

	sudo docker compose up --build


gen-okey: 

	go run ./_okeygen/keygen.go

	cp okey osock/okey

