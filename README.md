# gomotics

[![GitHub release](https://img.shields.io/github/release/mch1307/gomotics.svg)](https://github.com/mch1307/gomotics/releases)
[![Travis branch](https://img.shields.io/travis/mch1307/gomotics/master.svg)](https://travis-ci.org/mch1307/gomotics)
[![Coverage Status](https://coveralls.io/repos/github/mch1307/gomotics/badge.svg?branch=master)](https://coveralls.io/github/mch1307/gomotics?branch=master)
[![Go Report Card](https://goreportcard.com/badge/mch1307/gomotics)](http://goreportcard.com/report/mch1307/gomotics)
[![license](https://img.shields.io/github/license/mch1307/gomotics.svg)](https://github.com/mch1307/gomotics/blob/master/LICENSE.md) [![](https://images.microbadger.com/badges/image/mch1307/gomotics.svg)](https://microbadger.com/images/mch1307/gomotics "Get your own image badge on microbadger.com")

Go API server for Niko Home Control

"Roadmap": 

- [X] link with Jeedom
- [ ] GUI?


More information/doc on https://blog.csnet.me/gomotics/

## Installation:

### Docker

A docker image is automatically build with Travis-CI. It is available on [Docker Hub](https://hub.docker.com/r/mch1307/gomotics/)

> docker run -d -P --net host --name gomotics mch1307/gomotics

### Binaries

Download your platform binary from the release page, extract the executable from the archive. 

## Running
gomotics will run with default config if you do not provide a configuration file. If you want to link gomotics with Jeedom, provide the Jeedom URL and API key as follows

```
[jeedom]
url = "http://jeedom/core/api/jeeApi.php"
apikey = "abcdefgh1234"
```
A complete config file would like as follows:

``` 
[server]
[server]
ListenPort = 8081
LogLevel = "DEBUG"
LogPath = "."

[nhc]
host =          "x.x.x.x"
port =          8000

[jeedom]
url = "http://jeedom-host/core/api/jeeApi.php"
apikey = "abcdefgh1234"
```

Then start gomotics as follows:

> gomotics -conf path/confg.toml