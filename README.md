# blazer - CLI tool for uploading/downloading files

- Control thread count
- File verification with SHA
- Rich meta data
- Robust
- Resume from interruption

# usage
blazer -url=example.com/1.pdf -thread=10 -out=file.pdf

## Bugs:
- this didn't work
- enable multiple download in the same directory by using a unique ID for each downloads
- maintain history of downloads, if there is a flag -h = true
- Get "https://logic.pdmi.ras.ru/~gravin/storage/Concrete_Mathematics_2e.pdf": dial tcp 83.149.197.121:443: i/o timeout
- you should be able to run this from anywhere..
- Handle improper URL
- Get "https://mirror.yandex.ru/fedora/linux/releases/34/Workstation/aarch64/iso/Fedora-Workstation-Live-aarch64-34-1.2.iso": read tcp 192.168.0.113:38734->213.180.204.183:443: read: connection reset by peer
- pass multiple URLs or json of URLs
- Handle the case where there is no internet
- while downloading file if there is interruption and if the downloaded file is incomplete, what will you do,
you can find out incomplete file download by checking expected bytes to be written and bytes written, and then 
you can try to retry that attempt
- Get "https://sample-videos.com/img/Sample-jpg-image-30mb.jpg": read tcp 192.168.43.62:41606->103.145.51.95:443: read: connection reset by peer
- TCP timeout incase of large file
- Try to download a large file like 10 or 20 gb and 
- implement retry if a segment is corrupted
- Get "https://releases.ubuntu.com/20.04.2.0/ubuntu-20.04.2.0-desktop-amd64.iso": EOF
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x40 pc=0x64366a]
- check and handle http status code
- even if the program crashes, use exception handling and restart the download with same thread count  
- move all constants in a single file, for more clarity
- write basic tests and add it to pipeline
- move it to top system level, find out where linux apps store temroary files ?

## Calcualating ETA
- ETA depends on the number of threads, download speed, speed of server

## Experiments and benchmarks
- Create a ruby server and try to serve a 10 gb file and try to do the same in go and other servers and
estimate performance and also node js

# Todo:
- use interactive terminal support to print messages
- show time elapsed
- find optimal concrrency to realize parallel downloads
- support for verbose logging
- print the performance graph as well to see perfromance difference, using different number of threads and 
- storing data in current drive
- show progress
- also need support for google-drive, dropbox, a simple youtube version too.
- use interactive terminal to show progress across each thread, so it more cooler
- upload it to free file sharing service upto some MB, you can build one for yourself and share
// it with others, write in go, or node js and also use it here, a simple service
// the life of the service can be just 1 hour, there you can learn so many interesting things
- download files after giving username and password(with authentication)
- export as package too, later somehow
- show current download speed and ETA
large file sizes

# Roadmap:
- Write a compatable, high speed server to compliment this client
- Support torrent based
- GUI maybe in electron or QTCreate
