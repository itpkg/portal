namespace :books do
  desc 'convert book files'
  task :convert do
    on roles(:web) do |_|
      within release_path do | |
        execute :rake, 'books:convert'
      end
    end
  end
end