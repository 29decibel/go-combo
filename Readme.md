## YUI3 Combo Handler by [Go](http://golang.org)

> **Development in progress, not ready for production, need performance test.**


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
# choose your platform executable file in need(gocombo.linux or gocombo.exe)
$ ./bin/gocombo.mac

# That's it! Now your go-combo server is running. If you want to more customization,
# you can provide ```--base```, ```--with-version``` or ```--port``` flags.

# provide custom yui3 build direcotry
# choose your platform executable file in need(gocombo.linux or gocombo.exe)
$ ./bin/gocombo.mac --base="./some-where-else/yui3/build/"


# you can also put different yui3 versions into different directories like this
# ./yui3/3.10.1/build/.
# ./yui3/3.11.0/build/.
# choose your platform executable file in need(gocombo.linux or gocombo.exe)
$ ./bin/gocombo.mac --base="./yui3/"  --with-version=true
```

#### Compile for different platforms
```bash
# make sure install your go with --cross-compile-common option
$ brew install go --cross-compile-common

# build for mac
$ go build -o gocombo.mac ./server/main.go

# build for linux
$ GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o gocombo.linux ./server/main.go

# build for windows
$ GOOS=windows GOARCH=386 go build -o gocombo.exe ./server/main.go
```

### TODO
* ~~version number support~~
* ~~custom port~~
* ~~able to resovle relative base dir(gocombo ../yui3/build/ )~~
* ~~clone YUI instruction~~
* ~~get OptionValue performance issue~~
* ~~compiled executable program for Linux platform~~
* Makefile to create three platform executables
* gh-pages, logo
* https support
* performance test
* more tests

