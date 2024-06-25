
build:

	cd ofront && go build -o ofront.out .

	cd osock && go build -o osock.out .


up: 

	sudo docker compose up --build


gen-okey: 

	go run ./_okeygen/keygen.go

	cp okey osock/okey

