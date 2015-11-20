namespace :ueditor do

  desc 'build ueditor library'
  task :build do
    on roles(:web) do |_|
      target = "#{release_path}/public/ueditor"
      within "#{release_path}/vendor/assets/bower_components/ueditor" do | |

        execute :npm, :install
        execute :grunt
        execute :mkdir, '-pv', target
        %w(dialogs themes lang third-party ueditor.all.min.js).each do |f|
          execute :cp, '-a', "dist/utf8-php/#{f}", "#{target}/#{f}"
        end
      end
    end
  end
end

after 'deploy', 'ueditor:build'