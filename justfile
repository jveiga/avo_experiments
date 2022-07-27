simple_sum:
	go run avo/sums/sums.go -out sums.s

test-all:
	go test -v ./...

bench-all:
	go test -v -bench=. ./...

fuzz-add-simple:
	go test -v -fuzz=. -run=/AddSimple/
