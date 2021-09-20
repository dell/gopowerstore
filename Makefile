
integration_tests_path=./inttests
unit_test_paths= ./ ./api

.PHONY: mocks

all: unit-test int-test check gosec mocks

unit-test:
	go clean -cache
	go test -v -coverprofile=c.out $(unit_test_paths)

int-test:
		 go test -v -coverprofile=c.out -coverpkg github.com/dell/gopowerstore \
		 $(integration_tests_path)

gocover:
	go tool cover -html=c.out

check: mock-test gosec
	gofmt -w ./.
	golint ./...
	go vet

mocks:
	mockery --all --disable-version-string

gosec:
	gosec -quiet -log gosec.log -out=gosecresults.csv -fmt=csv ./...
