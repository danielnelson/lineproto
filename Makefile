
lineproto: parser.go .FORCE
	go build -o lineproto ./cmd/parser/main.go

parser.go: parser/grammar.peg Makefile
	pigeon -cache -optimize-basic-latin -optimize-parser -optimize-grammar -o parser.go parser/grammar.peg
	@# pigeon -cache -optimize-basic-latin  -optimize-grammar -o parser.go parser/grammar.peg

lineproto.%: parser.go .FORCE
	go test -bench=Benchmark_500lines -benchmem -memprofile=lineproto.mem ./. -run none

.PHONY: pprof-alloc
pprof-alloc: lineproto.%
	go tool pprof --alloc_objects lineproto.test lineproto.mem

.PHONY: run
run: lineproto
	./lineproto

.PHONY: clean
clean:
	rm -f lineproto lineproto.test lineproto.mem

.FORCE:
