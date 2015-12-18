it-package portal(golang version)
---
## Development

### Editor
 * [Intellij Idea](https://github.com/go-lang-plugin-org/go-lang-idea-plugin/wiki/v1.0.0-Setup-initial-project)
 * [Emacs](.emacs)
 
 
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

### Setup database (as postgres)

    psql
    CREATE USER portal WITH PASSWORD 'changeme';
    CREATE DATABASE portal_production WITH ENCODING='UTF8';
    GRANT ALL PRIVILEGES ON DATABASE portal_production TO portal;
    \q
    
#### test database connection


    psql -U portal -d portal_production    

* if report 'FATAL:  Peer authentication failed for user "portal"', open file "/etc/postgresql/9.3/main/pg_hba.conf" change line "local   all             all                                     peer" to "local   all             all                                     md5" and then run: 

    service postgresql restart

### Start
    go run app.go
    npm run start

## Deployment

### Build
    make
    ls release

### Run
    export ITPKG_ENV=production
    ls config # edit config files
    ./itpkg -h
    ./itpkg db migrate
    ./itpkg db seed
    ./itpkg nginx
    ./itpkg build
    ./itpkg server
     


