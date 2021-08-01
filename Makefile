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

release:
	echo "continue in dev branch"
	echo "update changelog, global constant's version"
	echo "commit current changes"
	echo "checkout to main and pull latest code"
	echo "Create tag"
	echo "build binaries [binary for mac and windows - 64 bit]"
	gox -parallel=3  -os="darwin windows linux" -arch="amd64"
	
