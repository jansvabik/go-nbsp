# Non-breakable space auto-adder for Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/jansvabik/go-nbsp)](https://goreportcard.com/report/github.com/jansvabik/go-nbsp)
[![Coverage tests](http://gocover.io/_badge/github.com/jansvabik/go-nbsp)](https://gocover.io/github.com/jansvabik/go-nbsp)
[![Build Status](https://travis-ci.org/jansvabik/go-nbsp.svg?branch=master)](https://travis-ci.org/jansvabik/go-nbsp)
[![Maintainability](https://api.codeclimate.com/v1/badges/06a33ed30e237fa413ee/maintainability)](https://codeclimate.com/github/jansvabik/go-nbsp/maintainability)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/jansvabik/go-nbsp/issues)

Go package for replacing spaces by unbreakable spaces where they should be according to the typographical rules of the specified language. Currently, I still work on this project. This is a new version of my very first Go program ([gonbspcz](https://github.com/jansvabik/gonbspcz)) and it shouldn't work just for Czech language, but also for any other language that will be configured by corresponding lang_CODE.go file.

## Installation
You can add this package to your workspace by installing it and then importing it.

1. Run command `go get github.com/jansvabik/go-nbsp`
2. Import package `import "github.com/jansvabik/go-nbsp"`
