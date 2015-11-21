CarrierWave.configure do |config|
  config.permissions = 0600
  config.directory_permissions = 0700
  config.storage = :fog

  config.fog_credentials=YAML.load(ERB.new(File.new("#{Rails.root}/config/fog.yml").read).result)[Rails.env]
  config.fog_directory = 'tmp'
  config.fog_public = true
  config.fog_attributes = {'Cache-Control' => "max-age=#{5.years.to_i}"}

end