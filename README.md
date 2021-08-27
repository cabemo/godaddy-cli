<img src="https://storage.googleapis.com/cabemo/godaddy-cli.svg" width="300px" />

# GoDaddy CLI
As an everyday maintainer of many domains I've found CLIs very useful (e.g AWS). I've created this simple CLI to manage your GoDaddy domains.


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

### List your domains 

```bash
$ godaddy domains list
awesomedomain.com
anotherone.org
```

### Availability search

```bash
$ godaddy domains search example.com
Domain       Available Price     Period
example.com  true      11.99 USD 1 year(s)
```

### List the DNS records for one of your domains
```bash
$ godaddy records list --domain example.com
A    	 600	             @	134.233.0.41
NS   	3600	             @	ns23.domaincontrol.com
NS   	3600	             @	ns24.domaincontrol.com
CNAME	3600	           www	@
CNAME	3600	_domainconnect	_domainconnect.gd.domaincontrol.com
```

### Add DNS records (A, AAAA, CNAME, MX, TXT)

Type and ttl have default values 'A' and 600 respectively. If you are setting an MX record you must pass the --priority flag too.

```bash
$ godaddy records add --domain example.com --type txt --name some --value test --ttl 600
Added: TXT 600 some test
```

### Remove DNS records

In here you only need the domain, type and name

```bash
$ godaddy records remove --domain example.com --type txt --name some
```
