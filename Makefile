blazer:
	rm -rf build
	mkdir build
	go build -o build/blazer main.go
	echo "Binary is build successfully"

package:
	rm -rf build
	mkdir build
	go build -o build/blazer main.go
	sudo cp build/blazer /usr/local/bin/
	echo "binary is now globally available"

