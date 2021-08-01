# v0.3-beta
### Features
- Checksum verification SHA256
- Guarentee per segment download
- Handle network errors in segment download, validations
- Indicate download status if download is resumed

# v0.2-alpha
### Features
- Resume file from interrupted download
- Use session_ID to store different downloads at same folder
- performance improvements 
- Show download time, more logs
- fix bug with request creation
- control thread count

### Known issues:
* Unstable, did not handle edge cases, network error codes etc.,
* Can't perform multiple downloads in the same folder

# v0.1-alpha

### Features
* Can concurrently download file of any size
* Control thread count

### Known issues:
* unstable, did not handle edge cases, network error codes etc.,
* Can't perform multiple downloads in the same folder
