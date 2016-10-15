# fetch-ssh-keys
`fetch-ssh-keys` is small shell command to get users public SSH keys from different cloud services like [GitHub](https://github.com).

> Note: Tested only on Linux and Mac. If you test on any other platform, please me know!

## Usage
```shell
# Fetch users from GitHub and output in SSH authorized_keys format
fetch-ssh-keys github --output ssh
```

Tool can be used for example to automatically update `.ssh/authorized_keys` file by adding the script to cron job.

## Installation
- Download binary from [releases](https://github.com/ernoaapa/fetch-ssh-keys/releases)
- Give execution rights (`chmod +x fetch-ssh-keys`) and add it into your $PATH

### Configuration
TODO

## Development
### Get dependencies
```shell
go get ./...
```

### Run stastd-exec
```shell
go run main.go github --output ssh
```
