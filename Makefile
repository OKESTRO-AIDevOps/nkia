
all:
	@echo "specify option"
	@echo "build : build and push runnable(testable) environment"
	@echo "run   : run orch.io"

build:

	mkdir -p _out/nokubeadm

	mkdir -p _out/nokubectl

	mkdir -p _out/nokubelet

	cd ./nokubeadm && make build

	cd ./nokubectl && make build

	cd ./nokubelet && make build

	cd ./orch.io && make build

build-commit:

	mkdir -p _out/nokubeadm

	mkdir -p _out/nokubectl

	mkdir -p _out/nokubelet

	cd ./nokubeadm && make build

	cd ./nokubectl && make build

	cd ./nokubelet && make build

	cd ./orch.io && make build

	git pull

	git add .

	git commit 

	git fetch --all

	git rebase upstream/main

	git push

run:

	cd ./orch.io && make run


