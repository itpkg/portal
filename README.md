it-package portal(golang version)
---
## Development
### Install gvm
    bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
    gvm install go1.4
    GOROOT_BOOTSTRAP=/home/flamen/.gvm/gos/go1.4 gvm install go1.5.2
    gvm use go1.5.2 --default
    go get github.com/itpkg/portal

### Clone code
    cd $GOPATH/src/github.com/itpkg/portal
    git checkout go
    npm install

### Start
    go run app.go
    npm run start

## Deployment

### Build
    make
    ls release


