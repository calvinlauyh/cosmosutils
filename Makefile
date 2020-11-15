all: build/decode-cosmosbase64tx build/decode-cosmostx

build/decode-cosmosbase64tx:
	go build -o ./build/decode-cosmosbase64tx ./cmd/decode-cosmosbase64tx

build/decode-cosmostx:
	go build -o ./build/decode-cosmostx ./cmd/decode-cosmostx

clean:
	rm ./build/decode-cosmosbase64tx ./build/decode-cosmostx
