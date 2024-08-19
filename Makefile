
all:
	@echo "specify option"
	@echo "build   : build and push runnable(testable) environment"
	@echo "commit  : commit"
	@echo "release : build, commit and generate release binary"
	@echo "run     : run orch.io"
	@echo "stage   : stage to all downstream repos including docs"



.PHONY: orch.io
orch.io:

	make -C orch.io gen-okey

	make -C orch.io build


#orch.io-db:

#	make -C orch.io db


orch.io-up:

	make -C orch.io gen-okey

	make -C orch.io up 




build:

	make -C nokubeadm build 

	make -C nokubelet build 

	make -C nokubectl build 

	cd hack && ./libgen.sh && mv lib ..

	/bin/cp -Rf lib nokubeadm/

	/bin/cp -Rf lib nokubelet/

	sudo rm -rf nokubeadm/.usr nokubeadm/.etc nokubeadm/.npia/.init

	sudo rm -rf nokubelet/.usr nokubelet/.etc nokubelet/.npia/.init

	cp orch.io/certs.tar.gz.gpg nokubectl/.npia/

	gpg --output nokubectl/.npia/certs.tar.gz --decrypt nokubectl/.npia/certs.tar.gz.gpg

	tar -xzf nokubectl/.npia/certs.tar.gz -C nokubectl/.npia/

	rm -r lib

build-noctl:

	make -C nokubeadm build 

	make -C nokubelet build 

	cd hack && ./libgen.sh && mv lib ..

	/bin/cp -Rf lib nokubeadm/

	/bin/cp -Rf lib nokubelet/

	sudo rm -rf nokubeadm/.usr nokubeadm/.etc nokubeadm/.npia/.init
	
	sudo rm -rf nokubelet/.usr nokubelet/.etc nokubelet/.npia/.init


	rm -r lib

.PHONY: infra
infra:

	make -C infra build

infra-ci:

	cd ./infra && /bin/cp -Rf infractl ../ && /bin/cp -Rf ./.npia.infra ../


	sudo ./infractl 	--repo https://github.com/OKESTRO-AIDevOps/nkia.git \
			   	        --id seantywork \
			   	        --token - \
			            --name nkia \
				        --plan ci \


	sudo rm -rf ./infractl ./.npia.infra

clean:

	rm -rf *.out lib

	make -C nokubeadm clean 

	make -C nokubectl clean 

	make -C nokubelet clean