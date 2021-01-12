all: build/decode-cosmosbase64tx build/decode-cosmostx build/decode-cosmospubkey

build/decode-cosmosbase64tx:
	go build -o ./build/decode-cosmosbase64tx ./cmd/decode-cosmosbase64tx

build/decode-cosmostx:
	go build -o ./build/decode-cosmostx ./cmd/decode-cosmostx

build/decode-cosmospubkey:
	go build -o ./build/decode-cosmospubkey ./cmd/decode-cosmospubkey

clean:
	rm ./build/decode-cosmosbase64tx ./build/decode-cosmostx ./build/decode-cosmospubkey
