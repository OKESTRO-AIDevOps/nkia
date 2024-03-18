
all:
	@echo "specify option"
	@echo "build   : build and push runnable(testable) environment"
	@echo "commit  : commit"
	@echo "release : build, commit and generate release binary"
	@echo "run     : run orch.io"
	@echo "stage   : stage to all downstream repos including docs"

build:

	cd ./nokubeadm && make build

	cd ./nokubectl && make build

	cd ./nokubelet && make build

	cd ./orch.io && make build

	cd ./infra && make build


commit:


	git add .

	git commit 

	git fetch --all

	git rebase upstream/main

	git push


release:

	mkdir -p nkia/nokubeadm

	mkdir -p nkia/nokubectl

	mkdir -p nkia/nokubelet

	cd ./nokubeadm && make release

	cd ./nokubectl && make release

	cd ./nokubelet && make release

	cd hack && ./libgen.sh

	/bin/cp -Rf ./hack/binupdate.sh ./nkia/

	tar -czvf lib.tgz lib

	tar -czvf nkia.tgz nkia


release-commit:

	mkdir -p nkia/nokubeadm

	mkdir -p nkia/nokubectl

	mkdir -p nkia/nokubelet

	cd ./nokubeadm && make release

	cd ./nokubectl && make release

	cd ./nokubelet && make release

	cd hack && ./libgen.sh

	/bin/cp -Rf ./hack/binupdate.sh ./nkia/

	tar -czvf lib.tgz lib

	tar -czvf nkia.tgz nkia

	git add .

	git commit 

	git fetch --all

	git rebase upstream/main

	git push

run:

	cd ./orch.io && make run


stage:

	cd ./infra && /bin/cp -Rf infractl ../ && /bin/cp -Rf ./.npia.infra ../


	sudo ./infractl 	--repo https://github.com/OKESTRO-AIDevOps/nkia.git \
			   	        --id seantywork \
			   	        --token - \
			            --name nkia \
				        --plan ci \


	sudo rm -rf ./infractl ./.npia.infra

