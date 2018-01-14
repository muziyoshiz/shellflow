# shellflow

Build a straight workflow from shell commands

## Description

shellflow provides the following features for building a straight workflow:

- shellflow receives shell commands from stdin
- shellflow executes the commands in sequence
- If one of the commands fails, shellflow stops
- shellflow can continue the execution from the specified order of the commands

## Usage

Execute all commands:

```
$ shellflow < workflow.sh
```

Execute third command only:

```
$ shellflow 3 < workflow.sh
```

Execute third command and its trailing commands:

```
$ shellflow 3+ < workflow.sh
```

## Install

If you are OSX user, you can use [Homebrew](https://brew.sh/):

```
$ brew tap muziyoshiz/shellflow
$ brew install shellflow
```

If you are in another platform, you can download binary from [release page](https://github.com/muziyoshiz/shellflow/releases) and place it in `$PATH` directory.

Or you can use go get (you need to use go1.7 or later),

```
$ go get -u github.com/muziyoshiz/shellflow
```

## Develop

### Setup

```
$ make bootstrap
$ export GITHUB_USER="..."
$ export GITHUB_TOKEN="..."
```

### Build

```
$ make build
```

### Test

```
$ make test-all
```

### Release 

```
$ make package
$ make upload
```

### Update Homebrew formula in ../homebrew-shellflow

```
$ make brew
```

## License

[MIT](https://github.com/muziyoshiz/shellflow/blob/master/LICENCE)

## Author

[Masahiro Yoshizawa](https://github.com/muziyoshiz)
