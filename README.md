# gowit

`gowit` is the Go SDK for [Wit.ai](https://wit.ai/)

## Install
Using `go get`:

	go get github.com/dragonraid/gowit
	
From source:

	git clone https://github.com/dragonraid/gowit $GOPATH/src/github.com/dragonraid/
	
## Usage:

For minimal working example you need to specify `WIT_API_TOKEN` environment variable. Then you can do:
```
witAPI := wit.New()
body, _, err := witApi.Message("YOUR MESSAGE").Do()
```

See the also `examples` directory

## API

### Versioning
The default API version is `20180527`. You can target a specific version by setting the env variable `WIT_API_VERSION`. 

### Overview

`gowit` provides interaction with following endpoints:

* `Message` - [the Wit message API](https://wit.ai/docs/http/20170307#get--message-link)


## Optional settings
You have two options how to setup your client behaviour
* Environment variables
	* `WIT_API_URL` (defaults to `20180527`)
	* `WIT_API_VERSION` (defaults to `20180527`)
	* `WIT_API_VERBOSE` (defaults to `false`) 
	
	
* Wit struct 
```
witApi := &wit.Wit{
	Token:   "<WIT_API_TOKEN>",
	URL:     "<WIT_API_URL>",
	Version: "<WIT_API_VERSION>",
	Verbose: "<WIT_API_VERBOSE>",
}
```
