all:

	@echo "nkia-ci"
	@echo "specify option"
	@echo "build        : build cibot"
	@echo "build-commit : build and push runnable(testable) environment"
	@echo "test         : build from var"



build: 

	go build -v -o cibot ./robot


build-commit:

	go build -v -o cibot ./robot



var:

	go build -v -o test ./var



clean :

	rm -r cibot test
