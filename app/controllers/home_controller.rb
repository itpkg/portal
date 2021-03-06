require 'rss'
require 'portal/book'

class HomeController < ApplicationController
  def index
    @hot = Cms::Article.select(:id, :title, :summary, :logo).where(lang: I18n.locale).order(visits: :desc).limit(8).select { |a| a.logo }
    @top = Cms::Article.select(:id, :title, :summary, :logo).where(top: true)
    @users = User.select(:id, :username, :details, :logo).order(sign_in_count: :desc).limit(4).select { |u| !u.is_admin? }
    @books = Portal::Book.names.sample 6
    @notices = Notice.select(:content, :created_at).where(lang: I18n.locale).order(id: :desc).limit(6)
    render layout: 'cms'
  end

  def about_us
    @key = Setting.google_map_browser_key
    @loc = Setting.site_geometry
    @place_id = Setting.site_place_id
    @leave_word = LeaveWord.new
    render layout: 'cms'
  end

  def search
    @keyword = params[:@keyword]
    @id = Setting.google_search_id
    render layout: 'cms'
  end

  def rss
    locale = I18n.locale
    xml = Rails.cache.fetch("cache://#{locale}/rss.atom", expires_in: 8.hours) do
      rss = RSS::Maker.make('atom') do |maker|
        maker.channel.author = Setting.site_author
        maker.channel.updated = Time.now.to_s
        maker.channel.about = rss_url
        maker.channel.title = Setting.get_site_info 'title'

        Cms::Article.select(:id, :title, :updated_at).where(lang: locale).order(id: :desc).limit(120).each do |a|
          maker.items.new_item do |item|
            item.link = cms_articles_url(id: a.id)

            item.title = a.title
            item.updated = a.updated_at.utc
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
