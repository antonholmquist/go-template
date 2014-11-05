# Project template for Go

### Export GOPATH

Should be set to your workspace directory

```bash
$ export GOPATH=~/Projects
$ export PATH=$PATH:~/Projects/bin
```

### Make sure this project is in the right folder

```bash
$ ~/Projects/src/go-template
```

### Install mercurial

Needed to install godep

```bash
$ brew install mercurial
```

### Install godep

```bash
$ go get github.com/tools/godep
```

### Copy dependencies from local project into workspace

```bash
$ godep restore
```

### Run

```bash
$ go install $GOPATH/bin/go-template
```

### Heroku

1. Make sure we have buildback (requires heroku toolbelt)

        $ heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git

2. Procfile
The name of the binary that heroku should launch. This will be name of the folder that the app is in.
web: go-template

3. Make sure that Godeps/Godeps.json "ImportPath" is set to : "go-template"
