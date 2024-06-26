
build:

	cd ofront && go build -o ofront.out .

	cd osock && go build -o osock.out .

db:

	cd odb && sudo docker compose up --build



up: 

	sudo docker compose up --build



gen-okey: 

	go run ./_okeygen/keygen.go

	cp okey osock/okey

