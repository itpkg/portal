namespace :deploy do
  desc 'reload the database with seed data'
  task :seed do
    on roles(:db) do |_|
      within current_path do | |
        execute :rake, 'db:seed'
      end
    end
  end
end