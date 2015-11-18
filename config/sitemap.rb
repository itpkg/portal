SitemapGenerator::Sitemap.create(default_host: "https://#{ENV['PORTAL_DOMAIN']}") do
  Cms::Article.select(:id, :lang).order(id: :desc).each { |a| add cms_articles_path(id: a.id, locale: a.lang), changefreq: :monthly, priority: 0.7 }
  add root_path(locale: 'en'), changefreq: :daily, priority: 0.9
  add root_path(locale: 'zh-CN'), changefreq: :daily, priority: 0.9
end