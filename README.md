# fetch-ssh-keys
`fetch-ssh-keys` is small shell command to get users public SSH keys from different cloud services like [GitHub](https://github.com).

> Note: Tested only on Linux and Mac. If you test on any other platform, please me know!

## Usage
```shell
fetch-ssh-keys <source name> <parameters> <output file>
```

For example fetch users public SSH keys from GitHub `my-lovely-team` team in `my-awesome-company` organization and output in SSH authorized_keys format
```shell
fetch-ssh-keys github --organization my-awesome-company --team my-lovely-team --token YOUR-TOKEN-HERE ./the-keys
```

Tool can be used for example to automatically update `.ssh/authorized_keys` file by giving path to `.ssh/authorized_keys` as last argument and adding the script to cron job.

## Installation
- Download binary from [releases](https://github.com/ernoaapa/fetch-ssh-keys/releases)
- Give execution rights (`chmod +x fetch-ssh-keys`) and add it into your $PATH

### Configuration
| Parameter      | Required          | Description                                                                                               |
|----------------|-------------------|-----------------------------------------------------------------------------------------------------------|
| --format       | No (default ssh)  | Output format. Only ssh authorized_keys format supported for now                                          |

#### GitHub
| Parameter      | Required | Description                                                                                               |
|----------------|----------|-----------------------------------------------------------------------------------------------------------|
| --organization | Yes      | Name of the organization which members keys to pick                                                       |
| --team         | No       | Name of the team which members keys to pick                                                               |
| --token        | No       | GitHub API token to use for communication. Without token you get only public members of the organization. |
| --public-only  | No       | Return only members what are publicly members of the given organization                                   |

## Development
### Get dependencies
```shell
go get ./...
```

### Run
```shell
go run main.go github --output ssh
```
