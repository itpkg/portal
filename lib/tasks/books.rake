
namespace :books do
  desc 'convert books to html'
  task :convert do
    root ="#{Rails.root}/tmp/books"
    Dir["#{root}/*.tex"].each do |f|
      puts "find file #{f}"
      tf = f[(root.size+1)..-1]
      `cd #{root} && pandoc #{tf} --toc -o #{tf[0..-4]}html`
    end
  end
end