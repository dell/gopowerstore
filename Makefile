
integration_tests_path=./inttests
unit_test_paths= ./ ./api

all: unit-test int-test mock-test check gosec

unit-test:
	go clean -cache
	go test -v -coverprofile=c.out $(unit_test_paths)

int-test:
		 go test -v -coverprofile=c.out -coverpkg github.com/dell/gopowerstore \
		 $(integration_tests_path)

mock-test:
	go clean -cache
	go test -v ./mock

gocover:
	go tool cover -html=c.out

check: mock-test gosec
	gofmt -w ./.
	golint ./...
	go vet

mock-gen:
	mockgen -source client.go -package mock -destination mock/client_mock.go

gosec:
	gosec -quiet -log gosec.log -out=gosecresults.csv -fmt=csv ./...
