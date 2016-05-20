# LDAP SERVICE

LDAP service is microservice based on go-micro for LDAP/AD binding and authentication. [![CircleCI](https://circleci.com/gh/kazoup/ldap-srv/tree/master.svg?style=svg)](https://circleci.com/gh/kazoup/ldap-srv/tree/master) [![Docker Repository on Quay](https://quay.io/repository/kazoup/ldap-srv/status "Docker Repository on Quay")](https://quay.io/repository/kazoup/ldap-srv)

## API endpoints

RPC  | Description
-----|------------
Bind | used to bind to LDAP/AD
Login | used to authenticate
Search | search user and return user info and groups

## Getting started

1. Install Consul

	Consul is the default registry/discovery for go-micro apps. It's however pluggable.
	[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

2. Run Consul
	```
	$ consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul
	```


3. Download and start the service

	```shell
	go get github.com/kazoup/ldap-srv
	ldap-srv
	```

	OR as a docker container

	```shell
	docker run kazoup/ldap-srv --registry_address=YOUR_REGISTRY_ADDRESS
	```
