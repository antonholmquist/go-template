
# Export GOPATH

Should be set to your workspace directory  
`export GOPATH=~/Projects PATH=$PATH:~/Projects/bin`

# Make sure this project is in the right folder

`~/Projects/src/go-template`

# Install mercurial

Needed to install godep

`brew install mercurial`

# Install godep

`go get github.com/tools/godep`

# Copy dependencies from local project into workspace

`godep restore`

# Run

`go install`  
`$GOPATH/bin/go-template`

# Heroku

1. Make sure we have buildback (requires heroku toolbelt)
heroku config:set BUILDPACK_URL=https://github.com/kr/heroku-buildpack-go.git

2. Procfile
The name of the binary that heroku should launch. This will be name of the folder that the app is in.
web: go-template
