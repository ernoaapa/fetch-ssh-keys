# fetch-ssh-keys

[![Build Status](https://travis-ci.org/ernoaapa/fetch-ssh-keys.svg?branch=master)](https://travis-ci.org/ernoaapa/fetch-ssh-keys)
[![Go Report Card](https://goreportcard.com/badge/github.com/ernoaapa/fetch-ssh-keys)](https://goreportcard.com/report/github.com/ernoaapa/fetch-ssh-keys)

`fetch-ssh-keys` is small shell command to get users public SSH keys from different cloud services like [GitHub](https://github.com).

> Note: Tested only on Linux and Mac. If you test on any other platform, please me know!

## Usage
```shell
fetch-ssh-keys <source name> <parameters> <output file>
```

For example fetch users public SSH keys from GitHub `my-lovely-team` team in `my-awesome-company` organization and output in SSH authorized_keys format
```shell
# Fetch 'my-lovely-team' keys in 'my-awesome-company' organisation
fetch-ssh-keys github --organization my-awesome-company --team my-lovely-team --token YOUR-TOKEN-HERE ./the-keys

# Fetch 'ernoaapa' and 'arnested' public keys
fetch-ssh-keys github --user ernoaapa --user arnested  --token YOUR-TOKEN-HERE ./the-keys

# Fetch 'ernoaapa/fetch-ssh-keys' deploy keys (requires a Github token with the `repo` or `public_repo` scope)
fetch-ssh-keys github --deploy-key ernoaapa/fetch-ssh-keys  --token YOUR-TOKEN-HERE ./the-keys
```

Tool can be used for example to automatically update `.ssh/authorized_keys` file by giving path to `.ssh/authorized_keys` as last argument and adding the script to cron job.

## Installation
- Download binary from [releases](https://github.com/ernoaapa/fetch-ssh-keys/releases)
- Give execution rights (`chmod +x fetch-ssh-keys`) and add it into your $PATH

### Configuration
| Parameter      | Required          | Description                                                                                               |
|----------------|-------------------|-----------------------------------------------------------------------------------------------------------|
| --format       | No (default ssh)  | Output format. Only ssh authorized_keys format supported for now                                          |
| --file-mode    | No (default 0600) | File permissions when writing to a file                                                                   |
| --comment      | No (default none) | Include COMMENT at top and bottom                                                                         |

#### GitHub
| Parameter      | Description                                                                                               |
|----------------|-----------------------------------------------------------------------------------------------------------|
| --organization | Name of the organization which members keys to pick                                                       |
| --team         | Name or slug of the team which members keys to pick                                                               |
| --user         | Name of the user which keys to pick                                                                       |
| --deploy-key   | Name of the owner/repo which deploy keys to pick                                                          |
| --token        | GitHub API token to use for communication. Without token you get only public members of the organization. |
| --public-only  | Return only members what are publicly members of the given organization                                   |

You can give `--organisation` (optionally combined with `--team` flag) and/or one or more `--user` or `--deploy-key` flags.

The `--deploy-key` parameter requires a Github token with the `repo` or `public_repo` scope.

## Development
### Get dependencies
```shell
go get ./...
```

### Run
```shell
go run main.go github --output ssh
```
