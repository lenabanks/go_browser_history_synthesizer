build:
	rm -f bin/*

	env GOOS=linux go build -ldflags="-s -w" -o bin/activity_synthesizer src/vouchers/claim/claim.go

	chmod -R 777 bin