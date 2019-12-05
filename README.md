<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-zabbix.svg"/></a></p>

<p align="center">
  <a href="https://godoc.org/pkg.re/essentialkaos/zabbix.v1"><img src="https://godoc.org/pkg.re/essentialkaos/zabbix.v1?status.svg"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/zabbix"><img src="https://goreportcard.com/badge/github.com/essentialkaos/zabbix"></a>
  <a href="https://travis-ci.com/essentialkaos/zabbix"><img src="https://travis-ci.com/essentialkaos/zabbix.svg"></a>
  <a href='https://coveralls.io/github/essentialkaos/zabbix?branch=master'><img src='https://coveralls.io/repos/github/essentialkaos/zabbix/badge.svg?branch=master' alt='Coverage Status' /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-zabbix-master"><img alt="codebeat badge" src="https://codebeat.co/badges/a8a976b8-8fdc-4a65-8a4b-754c284db842" /></a>
  <a href="https://essentialkaos.com/ekol"><img src="https://gh.kaos.st/ekol.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#zabbix-version-support">Zabbix version support</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`zabbix` is a Go package for sending metrics data to Zabbix Server 3+.

### Installation

Before the initial install allows git to use redirects for [pkg.re](https://github.com/essentialkaos/pkgre) service (_reason why you should do this described [here](https://github.com/essentialkaos/pkgre#git-support)_):

```
git config --global http.https://pkg.re.followRedirects true
```

Make sure you have a working Go 1.11+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/zabbix.v1
```

For update to the latest stable release, do:

```
go get -u pkg.re/essentialkaos/zabbix.v1
```

### Zabbix version support

| Zabbix Version | Support Status |
|----------------|----------------|
| `1.x`          | No             |
| `2.x`          | No             |
| `3.x`          | Yes (_Full_)   |
| `4.x`          | Yes (_Full_)   |

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![Build Status](https://travis-ci.com/essentialkaos/zabbix.svg?branch=master)](https://travis-ci.com/essentialkaos/zabbix) |
| `develop` | [![Build Status](https://travis-ci.com/essentialkaos/zabbix.svg?branch=develop)](https://travis-ci.com/essentialkaos/zabbix) |

### License

[EKOL](https://essentialkaos.com/ekol)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
