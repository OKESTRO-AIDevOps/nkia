
build:

	cd ofront && go build -o ofront .

	cd osock && go build -o osock .

db:

	cd odb && sudo docker compose up --build



up: 

	sudo docker compose up --build



gen-okey: 

	go run ./_okeygen/keygen.go

	tar czf certs.tar.gz certs

	gpg -o certs.tar.gz.gpg --symmetric certs.tar.gz

	rm certs.tar.gz

	rm -r certs/*.crt certs/*.priv certs/*.pub

	/bin/cp -Rf certs_server/* osock/.npia/certs/

	rm -r certs_server/*.crt certs_server/*.priv certs_server/*.pub

