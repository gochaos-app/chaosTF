.PHONY : dev prod install compile

dev:
	echo "Compiling..."
	go build -o chaosTF . 

prod:
	echo "Compiling..."
	go build -ldflags="-s -w" -o chaosTF .

move:
	mv chaosTF ~/bin/chaosTF

install: prod move

compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/chaosTF-freebsd-amd64 .
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/chaosTF-linux-amd64 .
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o bin/chaosTF-linux-arm64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/chaosTF-darwin-amd64 .
	GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o bin/chaosTF-darwin-m1 .
