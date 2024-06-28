
all:
	@echo "specify option"
	@echo "build   : build and push runnable(testable) environment"
	@echo "commit  : commit"
	@echo "release : build, commit and generate release binary"
	@echo "run     : run orch.io"
	@echo "stage   : stage to all downstream repos including docs"


build:

	make -C nokubeadm build 

	make -C nokubectl build 

	make -C nokubelet build 

	cd hack && ./libgen.sh && mv lib ..

	/bin/cp -Rf lib nokubeadm/

	/bin/cp -Rf lib nokubelet/

	echo ""  > nokubeadm/.npia/.init

	echo ""  > nokubectl/.npia/.init

	echo ""  > nokubelet/.npia/.init

	rm -r lib

release:

	make -C nokubeadm build 

	make -C nokubectl build

	make -C nokubelet build

	mkdir -p nkia/nokubeadm

	mkdir -p nkia/nokubectl

	mkdir -p nkia/nokubelet

	/bin/cp -Rf nokubeadm/.npia nkia/nokubeadm/

	rm nkia/nokubeadm/.npia/.init

	/bin/cp -Rf nokubectl/.npia nkia/nokubectl/

	rm nkia/nokubectl/.npia/.init

	rm nkia/nokubectl/.npia/.priv

	/bin/cp -Rf nokubelet/.npia nkia/nokubelet/

	rm nkia/nokubelet/.npia/.init

	/bin/cp -Rf nokubelet/nkletd nkia/nokubelet/nkletd

	mv nokubeadm/nokubeadm nkia/nokubeadm/

	mv nokubectl/nokubectl nkia/nokubectl/

	mv nokubelet/nokubelet nkia/nokubelet/

	cd hack && ./libgen.sh && mv lib ..

	/bin/cp -Rf ./hack/binupdate.sh ./nkia/

	tar -czvf lib.tgz lib

	tar -czvf nkia.tgz nkia

	rm -r lib

	rm -r nkia


.PHONY: hack/release
hack/release:

	cd hack/release/x86_64-ubuntu-20 && docker compose up --build && cp -Rf _output ../../../_x86_64-ubuntu-20.out

	cd hack/release/x86_64-ubuntu-22 && docker compose up --build && cp -Rf _output ../../../_x86_64-ubuntu-22.out


.PHONY: orch.io
orch.io:

	make -C orch.io gen-okey

	make -C orch.io build


orch.io-db:

	make -C orch.io db


orch.io-up:

	make -C orch.io gen-okey

	make -C orch.io up 

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

	rm -rf *.out

	make -C nokubeadm clean 

	make -C nokubectl clean 

	make -C nokubelet clean