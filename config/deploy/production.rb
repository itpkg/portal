server '192.168.1.108', user: 'deploy', roles: %w{db app web}
set :deploy_to, '/var/www/demo'

