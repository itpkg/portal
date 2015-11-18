require 'rss'

class HomeController < ApplicationController
  def index
  end

  def sitemap

  end

  def rss
    xml = Rails.cache.fetch("cache://#{I18n.locale}/rss.atom", expires_in: 8.hours) do
      rss = RSS::Maker.make('atom') do |maker|
        maker.channel.author = Setting.site_author
        maker.channel.updated = Time.now.to_s
        maker.channel.about = rss_url
        maker.channel.title = Setting.get_site_info 'title'

        Cms::Article.order(id: :desc).limit(120).each do |a|
          maker.items.new_item do |item|
            item.link = cms_articles_url(id: a.id)
            item.title = a.title
            item.updated = a.updated_at
          end
        end
      end
      rss.to_s
    end

    render xml: xml
  end

  def baidu
    code = Setting.baidu_site_id
    if code == params.fetch(:id)
      render text: code
    else
      head :not_found
    end
  end

  def google
    code = Setting.google_site_id
    if code == params.fetch(:id)
      render text: "google-site-verification: google#{code}.html"
    else
      head :not_found
    end

  end

  def robots
    render plain: Setting.robots_txt
  end
end
