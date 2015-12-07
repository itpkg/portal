namespace :cache do
  desc 'clear cache'
  task :clear do
    on roles(:web) do |_|
      within release_path do | |
        execute :rake, 'cache:clear'
      end
    end
  end
end
