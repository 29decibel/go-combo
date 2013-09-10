## YUI3 Combo Handler by [Go](http://golang.org)

> **Development in progress, not ready for production**

### Basic usage
#### YUI3 configuration
```
....
combine:true,
comboBase: 'http://localhost:8123/combo?',
....
```

#### Start gocombo server
```
// directly use build dir:  /build/node-module/....
$ ./server/server --base="/Users/your-name/projects/yui3/build/" --port=4444 --ignore-version=true

// cdn dir structure: /cdn/0.0.1/build/node-module/....
$ ./server/server --base="/Users/your-name/projects/yui3/cdn/" --port=4444
```

### TODO
* √ ~~version number support~~
* √ ~~custom port~~
* able to resovle relative base dir(gocombo ../yui3/build/ )
* gh-pages, logo
* clone YUI instruction
* get OptionValue performance issue
* better command line help(must provide resource base, available options)
* compiled executable program for Linux platform
* https support
* performance test
* more tests

