namespace :generate do

  desc 'generate error files'
  task :errors do
    on roles(:web) do |_|
      %w(404 422 500).each do |code|
        obj = OpenStruct.new code: code
        template 'error.html.erb', "#{release_path}/public/#{code}.html", obj.instance_eval { binding }
      end
    end
  end
end

def template(from, to, obj)
  template_path = File.expand_path("../../../../config/deploy/templates/#{from}", __FILE__)
  template = ERB.new(File.new(template_path).read).result obj
  upload! StringIO.new(template), to
end


after 'deploy', 'generate:errors'

