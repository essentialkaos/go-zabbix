<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-zabbix.svg"/></a></p>

<p align="center">
  <a href="https://kaos.sh/g/go-zabbix"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev" /></a>
  <a href="https://kaos.sh/r/go-zabbix"><img src="https://kaos.sh/r/go-zabbix.svg" alt="GoReportCard" /></a>
  <a href="https://kaos.sh/w/go-zabbix/ci"><img src="https://kaos.sh/w/go-zabbix/ci.svg" alt="GitHub Actions CI Status" /></a>
  <a href="https://kaos.sh/w/go-zabbix/codeql"><img src="https://kaos.sh/w/go-zabbix/codeql.svg" alt="GitHub Actions CodeQL Status" /></a>
  <a href="https://kaos.sh/c/go-zabbix"><img src="https://kaos.sh/c/go-zabbix.svg" alt="Coverage Status" /></a>
  <a href="https://kaos.sh/b/go-zabbix"><img src="https://kaos.sh/b/e3257f5f-8f63-4d80-92d0-e083713efbed.svg" alt="Codebeat badge" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#zabbix-version-support">Zabbix version support</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`zabbix` is a Go package for sending metrics data to Zabbix Server 3+.

### Installation

Make sure you have a working Go 1.17+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get -d github.com/essentialkaos/go-zabbix
```

For update to the latest stable release, do:

```
go get -d -u github.com/essentialkaos/go-zabbix
```

### Zabbix version support

| Zabbix Version | Support Status |
|----------------|----------------|
| `1.x`          | No             |
| `2.x`          | No             |
| `3.x`          | Yes (_Full_)   |
| `4.x`          | Yes (_Full_)   |
| `5.x`          | Yes (_Full_)   |
| `6.x`          | Yes (_Full_)   |

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://kaos.sh/w/go-zabbix/ci.svg?branch=master)](https://kaos.sh/w/go-zabbix/ci?query=branch:master) |
| `develop` | [![CI](https://kaos.sh/w/go-zabbix/ci.svg?branch=develop)](https://kaos.sh/w/go-zabbix/ci?query=branch:develop) |

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
