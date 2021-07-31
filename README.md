# Blazer - CLI tool for concurrent file downloads

- Control thread count
- Resume from interruption

## Usage
``` blazer -url=example.com/1.pdf -t=10 -out=file.pdf ```

## Flags 
```
blazer -h
Usage of blazer:
  -checksum string
    	checksum SHA to verify file
  -out string
    	output path to store the downloaded file (default ".")
  -t int
    	Thread count - Number of concurrent downloads (default 10)
  -url string
    	Valid URL to download
  -v	prints current version of blazer
```

## Install

- Get latest build here: https://github.com/arvpyrna/blazer/releases

- or build from source: 

```
git clone git@github.com:arvpyrna/blazer.git
make package
```