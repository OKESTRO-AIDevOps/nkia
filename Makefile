
all:
	@echo "specify option"
	@echo "build   : build and push runnable(testable) environment"
	@echo "commit  : commit"
	@echo "release : build, commit and generate release binary"
	@echo "run     : run orch.io"
	@echo "stage   : stage to all downstream repos including docs"

build:

	mkdir -p bin/nokubeadm

	mkdir -p bin/nokubectl

	mkdir -p bin/nokubelet

	cd ./nokubeadm && make build

	cd ./nokubectl && make build

	cd ./nokubelet && make build

	cd ./orch.io && make build


commit:

	git pull

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

	cd ./orch.io && make build

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

	cd ./orch.io && make build

	cd hack && ./libgen.sh

	/bin/cp -Rf ./hack/binupdate.sh ./nkia/

	tar -czvf lib.tgz lib

	tar -czvf nkia.tgz nkia

	git pull

	git add .

	git commit 

	git fetch --all

	git rebase upstream/main

	git push

run:

	cd ./orch.io && make run


stage:


	@echo "not implemented"

