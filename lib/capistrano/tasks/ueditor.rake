namespace :ueditor do

  desc 'install ueditor library'
  task :install do

    on roles(:web) do |_|
      src = "#{shared_path}/build/ueditor"
      tgt = "#{shared_path}/public/3rd/ueditor"
      unless test["[ -d #{src}]"]
        execute :git, :clone, 'https://github.com/fex-team/ueditor.git', src
      end

      within src do
        execute :git, :pull
        execute :git, :checkout, fetch(:ueditor_version)
        execute :npm, :install
        execute :grunt
        execute :mkdir, '-p', tgt
        %w(dialogs themes lang third-party ueditor.all.min.js).each do |f|
          execute :cp, '-a', "dist/utf8-php/#{f}", "#{tgt}/#{f}"
        end
      end

    end
  end
end

