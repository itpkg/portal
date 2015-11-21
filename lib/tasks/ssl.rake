namespace :ssl do
  desc 'generate openssl certs'
  task :new do
    #todo
    echo 'openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx.key -out nginx.crt'
  end
end
