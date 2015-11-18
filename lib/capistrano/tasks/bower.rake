namespace :bower do
  desc 'install bower resources'
  task :install do
    on roles(:web) do |_|
      within release_path do | |
        execute :bower, :install
      end
    end
  end
end

before 'deploy:compile_assets', 'bower:install'
before 'bower:install', 'deploy:migrate'