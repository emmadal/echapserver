# echapserver API

##  Introduction
This document describes the RESTful HTTP interface for the echapserver
application. The echapserver application is a server that provides an endpoint to receive
and store data from IPMI sensors and other sources of information. It can be used as
a backend for projects like BMC Web UI or similar frontends, which need access to
data provided by hardware or software running on the system being monitored.

The documentation assumes knowledge of JSON format and basic understanding of how to use
HTTP methods (GET, POST). If you are not familiar with these concepts, please refer to
the [HTTP/1.1 specification](http://www.w3.org/Protocols/rfc2616/rfc2616.txt) before reading this guide.

**Note: This API is currently under development and may change in future releases without notice.**

## Base URL
All requests should be made against `/<hostname>:8080/` where `<hostname>` is the hostname or IP address of your echapserver instance.

## Caveats

Server built using :
- [Gin](https://gin-gonic.com/) web framework written in Golang
- [MySQL](https://www.mysql.com/) database
- [Docker](https://docker.com) for containerization
