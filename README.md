Portal system (by rails)
---

## Deploy

### On server

#### Create deploy user (as root)
    useradd -s /bin/bash -m deploy
    passwd -l deploy    
    echo "deploy ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/deploy
    chmod 0400 /etc/sudoers.d/deploy

#### Enable deploy user ssh login (as root)
    su - deploy
    mkdir .ssh
    chmod 700 .ssh    
    cat /tmp/id_rsa.pub > ~/.ssh/authorized_keys
    
#### Install packages (as root)
    apt-get update
    apt-get install build-essential
    apt-get install nginx postgresql redis-server 
    apt-get install nodejs git npm imagemagick  
    apt-get install libreadline-dev libpq-dev
    apt-get clean
    
    npm install -g bower grunt-cli
    ln -s /usr/bin/nodejs /usr/bin/node
    
    mkdir /var/www
    chown deploy:deploy /var/www


#### Install ruby (as deploy)
    git clone https://github.com/sstephenson/rbenv.git ~/.rbenv
    git clone https://github.com/sstephenson/ruby-build.git ~/.rbenv/plugins/ruby-build
    git clone https://github.com/sstephenson/rbenv-vars.git ~/.rbenv/plugins/rbenv-vars
    echo 'export PATH="$HOME/.rbenv/bin:$PATH"' >> ~/.bashrc
    echo 'eval "$(rbenv init -)"' >> ~/.bashrc
    
    ###### RE-LOGIN BEFORE CONTINUE #####
    rbenv install 2.2.3
    rbenv global 2.2.3
    gem install bundler
    
    mkdir -p /var/www/portal/production/shared/config/  
   
    
#### Setup database (as postgres)

    psql
    CREATE USER portal WITH PASSWORD 'changeme';
    CREATE DATABASE portal_production WITH ENCODING='UTF8';
    GRANT ALL PRIVILEGES ON DATABASE portal_production TO portal;
    \q
    
    ### test database connection
    psql -U portal -d portal_production    

* if report 'FATAL:  Peer authentication failed for user "portal"', open file "/etc/postgresql/9.3/main/pg_hba.conf" change line "local   all             all                                     peer" to "local   all             all                                     md5" and then run: 

    service postgresql restart
    
#### Create config files (as deploy)    
    
    
### On local 

#### Clone source
    git clone https://github.com/itpkg/portal.git
    cd portal
    
#### Copy setting files(change 192.168.1.108 to your hostname)
    scp config/*.yml deploy@192.168.1.108:/var/www/portal/production/shared/config/
    scp .rbenv-vars deploy@192.168.1.108:/var/www/portal/production/shared/
    
    ##### REMEMBER CHANGE  YOUR SETTING FILES ON SERVER BEFORE CONTINUE #####    
    cap production deploy:check 
    cap production puma:config
    cap production puma:nginx_config
