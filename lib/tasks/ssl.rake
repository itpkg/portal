namespace :ssl do
  desc 'generate openssl certs'
  task :new do
    puts `openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tmp/nginx.key -out tmp/nginx.crt`
  end
end
