# raru: run as random user
Simple and portable cli tool and golang package for run program as random user. Source code ported from C implementation [raru](https://github.com/teran-mckinney/raru).

## Installation

To install raru, please use `go get`.

### Command line tool

#### Exec
```
$ go get github.com/ArtemKulyabin/raru/cmd/raru
```

#### Spawn - fork and exec

```
$ go get github.com/ArtemKulyabin/raru/cmd/raru-spawn
```

#### Jail - run program as random user and changes the apparent root directory for the running process using a chroot(2) system call. If the program linked with dynamic libraries, they are copied into the root directory of the process.

```
$ go get github.com/ArtemKulyabin/raru/cmd/raru-jail
```

### Package

```
$ go get github.com/ArtemKulyabin/raru
```

## Usage

### Command line tool

```
$ raru ./fishy-app-youve-never-ran-before
...
$ raru bash # Whole shell as a random user.
...
$ raru curl https://fishysite
...
```

```
$ raru-jail run --chroot /dir /bin/ls
...
```

### Package

```
package main

import (
  "log"
  "github.com/ArtemKulyabin/raru"
)

func main() {
  log.Println(raru.Exec(os.Args[1], os.Args[2:]...))
}
```

## Cross compilation
For cross compilation you may use the [Gox](github.com/mitchellh/gox). Example:
```
$ gox github.com/ArtemKulyabin/raru/cmd/raru
```

## Operating system support
* Linux, FreeBSD, OpenBSD, NetBSD, Darwin(OS X)

### Tested on
* Ubuntu 14.04, FreeBSD 10.1, OpenBSD 5.6, NetBSD 6.1.5, OS X Yosemite 10.10.3
