namespace :cache do
  desc 'clear cache'
  task clear: :environment do
    Rails.cache.clear
  end
end