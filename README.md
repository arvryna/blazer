# Blazer - Concurrent file downloader

<p align="left">
  <a href="https://goreportcard.com/report/github.com/arvryna/blazer">
    <img src="https://goreportcard.com/badge/github.com/arvryna/blazer" />
  </a>
   <a href="http://makeapullrequest.com">
    <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" />
  </a>
</p>

- Control thread count
- Resume from interruption
- File integrity check - SHA256

## Install
3 different ways:
- ``` go get -u github.com/arvryna/blazer ```
- Download specific version from releases: https://github.com/arvryna/blazer/releases
- Build from source: 
```
git clone git@github.com:arvryna/blazer.git
make package
```

## Usage
``` blazer -url=example.com/1.mp4 -t=10  ```

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

## Benchmarks
| Name       |Size    | Blazer                  | cURL          | Wget         |
| -----------|--------| -----------             | ----          | -----        |
| Debian ISO | 300 MB | 1min 10 sec (25 threads)| 2min 40 sec   | 3 min 10 sec |
| Windows-10 | 5.4 GB | 20min 52sec (25 threads)| 46min 52 sec  | 40 min 25 sec|

## Demo
[![asciicast](https://asciinema.org/a/DInboSaUY2Ik9JIOcY4vZHRY9.svg)](https://asciinema.org/a/DInboSaUY2Ik9JIOcY4vZHRY9)

# Usage examples:
- Integration with node.js server [here](https://github.com/teyalite/client-sever-multithreading-downloader)

# FAQ

## How blazer works
Blazer makes use of [RFC7233](https://datatracker.ietf.org/doc/html/rfc7233) to identify the size of the resource to download, and then initiates partial downloads concurrently with the help of multiple go-routines, once all segements are downloaded, we finally merge all the individual segments into the final file. Number of segments = thread count passed as param to blazer CLI

## How file resumption work ?
* If download is either interrupted manually or because of any network errors (Timeout RCP(connection reset by peer) - may happen if "-t" is a large number than server can handle) then those segments may fail and you may have to restart the download with same URL and thread count for retry to work because, temp folder name is a hash of URL and thread count. 

* When download is re-initiated, blazer downloads only the segments that were not successful during the previous attempt. Because a temproary download cache is stored in the current directory of download, which will be used to restart from where it was left off.

* File resumption won't work if the number of threads used is different from the previous attempt.

# Roadmap:
* Possiblity of adding a config file for ~/.blazerc in home directory, to add new features like storing history of downloads
* Ability to perform concurrent file uploads
* Figuring out optimal thread count at runtime by considering various factors like network speed, server bandwidth, etc.,
* Showing an interactive progress bar with higher accuracy
* Add more algorithms for file integrity check
