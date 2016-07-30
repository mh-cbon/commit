# commit

Commit files on the underlying VCS.

## Install

Pick an msi package [here](https://github.com/mh-cbon/commit/releases)!

__deb/ubuntu/rpm repositories__

```sh
wget -O - https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/commit sh -xe
# or
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/source.sh \
| GH=mh-cbon/commit sh -xe
```

__deb/ubuntu/rpm packages__

```sh
curl -L https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/commit sh -xe
# or
wget -q -O - --no-check-certificate \
https://raw.githubusercontent.com/mh-cbon/latest/master/install.sh \
| GH=mh-cbon/commit sh -xe
```

__chocolatey__

```sh
choco install commit -y
```

__go__

```sh
mkdir -p $GOPATH/src/github.com/mh-cbon
cd $GOPATH/src/github.com/mh-cbon
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
   0.0.0

COMMANDS:
GLOBAL OPTIONS:
   --file value, -f value        File to add and commit
   --message value, -m value     Message of the commit
   --quiet, -q                   Silently fail
   --help, -h                    show help
   --version, -v                 print the version
```
