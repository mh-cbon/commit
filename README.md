# commit

Commit files on the underlying VCS __without failing__.

# Install

```sh
mkdir -p $GOPATH/github.com/mh-cbon
cd $GOPATH/github.com/mh-cbon
git clone https://github.com/mh-cbon/commit.git
cd commit
glide install
go install
```

# Usage

```sh
NAME:
   commit - Commit file

USAGE:
   commit -m <message> -f <file>

VERSION:
   0.0.1

COMMANDS:
GLOBAL OPTIONS:
   --file value, -f value	File to add and commit
   --message value, -m value	Message of the commit
   --help, -h			show help
   --version, -v		print the version
```
