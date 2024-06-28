
build:

	cd ofront && go build -o ofront .

	cd osock && go build -o osock .

db:

	cd odb && sudo docker compose up --build



up: 

	sudo docker compose up --build



gen-okey: 

	go run ./_okeygen/keygen.go

	tar czf certs_client.tar.gz certs_client

	gpg -o certs_client.tar.gz.gpg --symmetric certs_client.tar.gz

	rm certs_client.tar.gz

	rm -r certs_client/*.crt certs_client/*.priv certs_client/*.pub

	/bin/cp -Rf certs_server/* osock/.npia/certs/

	rm -r certs_server/*.crt certs_server/*.priv certs_server/*.pub

