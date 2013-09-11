## YUI3 Combo Handler by [Go](http://golang.org)

> **Development in progress, not ready for production**

> **Mac only for now, other platform need to be compiled from source**

### Basic usage
#### YUI3 configuration
```javascript
....
combine:true,
comboBase: 'http://localhost:8123/combo?',
....
```

#### Start gocombo server
```bash
# clone YUI3 into dir
$ git clone git@github.com:yui/yui3.git

# start gocombo (custom port like --port=:4321)
$ ./bin/darwin/gocombo

# That's it! Now your go-combo server is running. If you want to more customization,
# you can provide ```--base```, ```--with-version``` or ```--port``` flags.

# provide custom yui3 build direcotry
$ ./bin/darwin/gocombo --base="./some-where-else/yui3/build/"


# you can also put different yui3 versions into different directories like this
# ./yui3/3.10.1/build/.
# ./yui3/3.11.0/build/.
$ ./bin/darwin/gocombo --base="./yui3/"  --with-version=true
```

### TODO
* ~~version number support~~
* ~~custom port~~
* ~~able to resovle relative base dir(gocombo ../yui3/build/ )~~
* ~~clone YUI instruction~~
* ~~get OptionValue performance issue~~
* gh-pages, logo
* compiled executable program for Linux platform
* https support
* performance test
* more tests

