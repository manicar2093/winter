# Winter

An opinionated way to create APIs fast. I can't say it is a framework, just a gather of common used tools to create servers

This is powered by:

- [echo](https://echo.labstack.com/)
- [gorm](https://gorm.io/)
- [caarlos0/env](github.com/caarlos0/env/v10)
- [coditory/go-error](github.com/coditory/go-errors)
- [gookit/validate](github.com/gookit/validate)
- [manicar2093/echoroutesview](github.com/manicar2093/echoroutesview)
- [manicar2093/goption](github.com/manicar2093/goption)
- [manicar2093/gormpager](github.com/manicar2093/gormpager)
- and many others (all in go.mod)

## Packages

### apperrors

It contains a way to handle errors from our API. It has a MessageError which is an identified error. This is handled to return the needed status and message and avoid handle an error from controller to send needed response

### connections

All supported connections supported by now: Gorm with Postgres and Redis.

This connections implements winter.httphealthcheck Checkable interface making them available to check their state.

### converters

Some functions that might be consider as utils (I know, utils package is dicourage to be used, but I could not find better name for it hahaha)

### echoer

An echo handler meant to be use on dev or test. It returns all data send to the `/echo` endpoint

### env

It has only one function to get env variables or panic

### httphealthcheck

It contains a way to register health check from services as need. You just need to implement Checkable interface to be able to register and a checked service

### logger

A log implementation configurated for supported `stages`

### stages

Contains supported stages by this framework

### validator

It contains the implementation used by echo to validate http requests.

