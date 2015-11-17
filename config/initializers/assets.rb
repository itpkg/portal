# Be sure to restart your server when you modify this file.

# Version of your assets, change this if you want to expire all your assets.
Rails.application.config.assets.version = '1.0'

# Add additional assets to the asset load path
# Rails.application.config.assets.paths << Emoji.images_path

# Precompile additional assets.
# application.js, application.css, and all non-JS/CSS in app/assets folder are already added.
# Rails.application.config.assets.precompile += %w( search.js )


# Bower asset paths
Rails.root.join('vendor', 'assets', 'bower_components').to_s.tap do |bower_path|
  Rails.application.config.sass.load_paths << bower_path
  Rails.application.config.assets.paths << bower_path
end

# Precompile Bootstrap fonts
Rails.application.config.assets.precompile << 'bootstrap/dist/css/bootstrap.css'
Rails.application.config.assets.precompile << 'bootstrap/dist/js/bootstrap.js'

# famfamfam
Rails.application.config.assets.precompile << %r(famfamfam-flags/dist/png/[\w-]+\.png$)
Rails.application.config.assets.precompile << %r(famfamfam-silk/dist/png/[\w-]+\.png$)

# riot
Rails.application.config.assets.precompile << 'riot/riot.js'