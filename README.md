# Blazer CLI - concurrent file downloader

- Control thread count
- Resume from interruption
- File integrity check - SHA256

## Usage
``` blazer -url=example.com/1.pdf -t=10  ```

## Flags 
```
blazer -h
Usage of blazer:
  -checksum string
    	checksum SHA256(currently supported) to verify file
  -out string
    	output path to store the downloaded file
  -t int
    	Thread count - Number of concurrent downloads (default 10)
  -url string
    	Valid URL to download
  -v	prints current version of blazer

```

## Install

- Download 64bit binary from releases: https://github.com/arvpyrna/blazer/releases

or

- Build from source: 

```
git clone git@github.com:arvpyrna/blazer.git
make package
```