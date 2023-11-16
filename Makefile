
all:
	@echo "specify option"
	@echo "build : build and push runnable(testable) environment"
	@echo "run   : run orch.io"

build:

	mkdir -p bin/nokubeadm

	mkdir -p bin/nokubectl

	mkdir -p bin/nokubelet

	cd ./nokubeadm && make build

	cd ./nokubectl && make build

	cd ./nokubelet && make build

	cd ./orch.io && make build

build-commit:

	mkdir -p bin/nokubeadm

	mkdir -p bin/nokubectl

	mkdir -p bin/nokubelet

	cd ./nokubeadm && make build

	cd ./nokubectl && make build

	cd ./nokubelet && make build

	cd ./orch.io && make build

	/bin/cp -Rf ./hack/libupdate.sh ./bin/

	/bin/cp -Rf ./hack/binupdate.sh ./bin/

	tar -czvf lib.tgz lib

	tar -czvf bin.tgz bin

	git pull

	git add .

	git commit 

	git fetch --all

	git rebase upstream/main

	git push

run:

	cd ./orch.io && make run


