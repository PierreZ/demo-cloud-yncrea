# Demo Cloud for yncrea

## Description

An simple application exposing a REST API that can be deployed on a:

* IaaS
* PaaS
* KaaS

## Start local PG

```shell
docker run --rm --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres
```