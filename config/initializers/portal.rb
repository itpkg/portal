Rails.application.config.redis = Redis.new
Rails.application.config.md2hm = Redcarpet::Markdown.new(Redcarpet::Render::HTML, autolink: true, tables: true)