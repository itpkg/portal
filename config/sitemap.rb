SitemapGenerator::Sitemap.create(default_host: "https://#{ENV['PORTAL_DOMAIN']}") do
  Cms::Article.select(:id, :lang).order(id: :desc).each { |a| add cms_articles_path(a, locale: a.lang), changefreq: :monthly, priority: 0.7 }
  Cms::Tag.select(:id, :lang).each { |t| add cms_tags_path(t, locale: t.lang), changefreq: :daily, priority: 0.7 }

  %w(en zh-CN).each do |l|
    add root_path(locale: l), changefreq: :daily, priority: 0.9
    add about_us_path(locale: 1), changefreq: :daily, priority: 0.9


    User.select(:id).each { |u| add cms_users_path(u, locale: l), changefreq: :daily, priority: 0.2 }
  end


end