all: build/decode-cosmosbase64tx build/decode-cosmostx build/pubkey-to-consaddress

build/decode-cosmosbase64tx:
	go build -o ./build/decode-cosmosbase64tx ./cmd/decode-cosmosbase64tx

build/decode-cosmostx:
	go build -o ./build/decode-cosmostx ./cmd/decode-cosmostx

build/pubkey-to-consaddress:
	go build -o ./build/pubkey-to-consaddress ./cmd/pubkey-to-consaddress

clean:
	rm ./build/decode-cosmosbase64tx ./build/decode-cosmostx ./build/pubkey-to-consaddress
