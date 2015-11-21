server '192.168.1.108', user: 'deploy', roles: %w{db app web}
set :server_name, '192.168.1.108'
set :deploy_to, '/var/www/demo'

