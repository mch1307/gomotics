# gomotics

[![GitHub release](https://img.shields.io/github/release/mch1307/gomotics.svg)](https://github.com/mch1307/gomotics/releases)
[![Travis branch](https://img.shields.io/travis/mch1307/gomotics/master.svg)](https://travis-ci.org/mch1307/gomotics)
[![Coverage Status](https://coveralls.io/repos/github/mch1307/gomotics/badge.svg?branch=master)](https://coveralls.io/github/mch1307/gomotics?branch=master)
[![Go Report Card](https://goreportcard.com/badge/mch1307/gomotics)](http://goreportcard.com/report/mch1307/gomotics)
[![license](https://img.shields.io/github/license/mch1307/gomotics.svg)](https://github.com/mch1307/gomotics/blob/master/LICENSE.md) [![](https://images.microbadger.com/badges/image/mch1307/gomotics.svg)](https://microbadger.com/images/mch1307/gomotics "Get your own image badge on microbadger.com")

Niko Home Control - Jeedom interface in Go
* expose API to NHC
* interface NHC with Jeedom
* automatically creates NHC items in Jeedom

More information/doc on https://blog.csnet.me/gomotics/

### See gomotics automatically creating NHC objects in Jeedom
[![gomotics v0.4](https://i.vimeocdn.com/video/666787552.webp)](https://player.vimeo.com/video/242817325 "gomotics v0.4")



## Installation:

### Docker

A docker image is automatically build with Travis-CI. It is available on [Docker Hub](https://hub.docker.com/r/mch1307/gomotics/)

### Binaries

Download your platform binary from the [release page](https://github.com/mch1307/gomotics/releases), extract the executable from the archive. 

See the [wiki](https://github.com/mch1307/gomotics/wiki) for an example of automating gomotics process startup with [supervisor](http://supervisord.org/)

## Running
### Config file
gomotics will run with default config if you do not provide a configuration file. If you want to link gomotics with Jeedom, provide the Jeedom URL and API key as follows

```
[jeedom]
url = "http://jeedom/core/api/jeeApi.php"
apikey = "abcdefgh1234"
```

To enable Gomotics to automatically create NHC items in Jeedom:
```
[jeedom]
url = "http://jeedom/core/api/jeeApi.php"
apikey = "abcdefgh1234"
autoCreateObjects= true

[server]
gmHostPort = "gomotics.local:8081"
```

A complete config file would like as follows:

``` 
[server]
listenPort = 8081
logLevel = "DEBUG"
logPath = "."
gmHostPort = "gomotics.local:8081"

[jeedom]
url = "http://jeedom-host/core/api/jeeApi.php"
apikey = "abcdefgh1234"
autoCreateObjects= false

[nhc]
host =          "x.x.x.x"
port =          8000

```
### env variables

Config can also be setup as env variable:

```
LISTEN_PORT         optional    default 8081
LOG_LEVEL           optional    default INFO
LOG_PATH            optional    default . (specify stdout for docker)
JEE_URL             mandatory if using Jeedom
JEE_APIKEY          mandatory if using Jeedom
NHC_HOST            optional    autodiscover
NHC_PORT            optional    autodiscover on port 8000
AUTO_CREATE_OBJECTS optional    default false
GM_HOSTPORT         mandatory if using AutoCreateObjects feature, default localhost
```
Then start gomotics as follows:

> gomotics -conf path/confg.toml

Or if using docker:

> docker run -d -P --net host --name gomotics -e JEE_URL=http://jeedom-host/core/api/jeeApi.php -e JEE_APIKEY=abcdegf1234 -e LOG_PATH=stdout mch1307/gomotics