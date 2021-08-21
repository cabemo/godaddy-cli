# GoDaddy CLI

As an everyday maintainer of many domains I've found CLIs very useful (e.g AWS). I've created this simple CLI to manage your account domains.


## Setup

In order to get godaddy-cli to work you need to get your [GoDaddy keys](https://developer.godaddy.com/keys) and place them in `$HOME/.config/godaddy/credentials.json`

## Installation

```bash
$ git clone git@github.com/Cabemo/godaddy-cli
$ cd godaddy-cli
$ make
$ sudo make install
```

## Uninstall

```bash
$ sudo make uninstall
```

## Examples

```bash
$ godaddy domains list
awesomedomain.com
anotherone.org
```
