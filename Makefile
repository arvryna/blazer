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
	echo "checkout to main"
	echo "download latest code from origin/main"
	echo "update changelog"
	echo "Update gloabal constant's variable"
	echo "Create tag with current commit"
	echo "build latest file and upload to release"
	echo "build binary for mac and windows - 64 bit"
	gox -parallel=3  -os="darwin windows linux" -arch="amd64"
	
