.PHONY: all binary1 binary2 clean

all: binary1 binary2

binary1:
	go build -o ./bin/binary1 ./cmd/binary1

binary2:
	go build -o ./bin/binary2 ./cmd/binary2

clean:
	rm -rf ./bin