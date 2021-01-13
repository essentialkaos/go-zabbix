<p align="center"><a href="#readme"><img src="https://gh.kaos.st/go-zabbix.svg"/></a></p>

<p align="center">
  <a href="https://pkg.re/essentialkaos/go-zabbix.v1?docs"><img src="https://gh.kaos.st/godoc.svg" alt="PkgGoDev"></a>
  <a href="https://goreportcard.com/report/github.com/essentialkaos/go-zabbix"><img src="https://goreportcard.com/badge/github.com/essentialkaos/go-zabbix"></a>
  <a href="https://github.com/essentialkaos/go-zabbix/actions"><img src="https://github.com/essentialkaos/go-zabbix/workflows/CI/badge.svg" alt="GitHub Actions Status" /></a>
  <a href="https://github.com/essentialkaos/go-zabbix/actions?query=workflow%3ACodeQL"><img src="https://github.com/essentialkaos/go-zabbix/workflows/CodeQL/badge.svg" /></a>
  <a href='https://coveralls.io/github/essentialkaos/go-zabbix?branch=master'><img src='https://coveralls.io/repos/github/essentialkaos/go-zabbix/badge.svg?branch=master' alt='Coverage Status' /></a>
  <a href="https://codebeat.co/projects/github-com-essentialkaos-go-zabbix-master"><img alt="codebeat badge" src="https://codebeat.co/badges/e3257f5f-8f63-4d80-92d0-e083713efbed" /></a>
  <a href="#license"><img src="https://gh.kaos.st/apache2.svg"></a>
</p>

<p align="center"><a href="#installation">Installation</a> • <a href="#zabbix-version-support">Zabbix version support</a> • <a href="#build-status">Build Status</a> • <a href="#license">License</a></p>

<br/>

`zabbix` is a Go package for sending metrics data to Zabbix Server 3+.

### Installation

Make sure you have a working Go 1.14+ workspace (_[instructions](https://golang.org/doc/install)_), then:

```
go get pkg.re/essentialkaos/go-zabbix.v1
```

For update to the latest stable release, do:

```
go get -u pkg.re/essentialkaos/go-zabbix.v1
```

### Zabbix version support

| Zabbix Version | Support Status |
|----------------|----------------|
| `1.x`          | No             |
| `2.x`          | No             |
| `3.x`          | Yes (_Full_)   |
| `4.x`          | Yes (_Full_)   |
| `5.x`          | Yes (_Full_)   |

### Build Status

| Branch | Status |
|--------|--------|
| `master` | [![CI](https://github.com/essentialkaos/go-zabbix/workflows/CI/badge.svg?branch=master)](https://github.com/essentialkaos/go-zabbix/actions) |
| `develop` | [![CI](https://github.com/essentialkaos/go-zabbix/workflows/CI/badge.svg?branch=develop)](https://github.com/essentialkaos/go-zabbix/actions) |

### License

[Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)

<p align="center"><a href="https://essentialkaos.com"><img src="https://gh.kaos.st/ekgh.svg"/></a></p>
