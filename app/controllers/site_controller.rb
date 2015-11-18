require 'etc'

class SiteController < ApplicationController
  before_action :must_be_admin!

  layout false

  def info

    case request.method
      when 'POST'
        si = params.require('site').permit(:title, :keywords, :description, :author, :copyright)
        Setting[_s_key('title')] = si.fetch :title
        Setting[_s_key('keywords')] = si.fetch :keywords
        Setting[_s_key('description')] = si.fetch :description
        Setting[_s_key('copyright')] = si.fetch :copyright
        Setting.site_author = si.fetch :author
        render json: {ok: true}
      else
    end
  end

  def captcha
    case request.method
      when 'POST'
        si = params.require('captcha').permit(:site_key, :secret_key)
        Setting.recaptcha_site_key = si.fetch :site_key
        Setting.recaptcha_secret_key = si.fetch :secret_key
        render json: {ok: true}
      else
    end
  end

  def seo
    case request.method
      when 'POST'
        si = params.require('seo').permit(:google_site_id, :baidu_site_id, :robots_txt)
        Setting.google_site_id = si.fetch :google_site_id
        Setting.baidu_site_id = si.fetch :baidu_site_id
        Setting.robots_txt = si.fetch :robots_txt
        render json: {ok: true}
      else
    end
  end

  def status
    @system = {
        user: ENV['USER'],
        time: Time.now,
        os: Etc.uname,
        ruby: {
            version: "#{RUBY_VERSION}(#{RUBY_RELEASE_DATE})",
            bin: `which ruby`.chomp,
        },
        rails: {
            env: Rails.env,
            version: Rails.version,
            root: Rails.root,
        }
    }

    @redis = Rails.application.config.redis.info

    conn = ActiveRecord::Base.connection
    cfg = Rails.configuration.database_configuration[Rails.env]
    @database = {
        size: conn.exec_query("select pg_size_pretty(pg_database_size('#{cfg['database']}'))").first['pg_size_pretty'],
        activity: conn.exec_query('SELECT pid,waiting,current_timestamp - least(query_start,xact_start) AS runtime, substr(query,1,32) AS current_query FROM pg_stat_activity WHERE NOT pid=pg_backend_pid()').map { |row|
          {pid: row['pid'], waiting: row['waiting'], runtime: row['runtime'], current_query: row['current_query']}
        },
    }
  end

  def adverts
    case request.method
      when 'POST'
        si = params.require('adverts').permit(:vertical, :horizontal)
        Setting.advert_horizontal = si.fetch :horizontal
        Setting.advert_vertical = si.fetch :vertical
        render json: {ok: true}
      else
    end
  end

  def users
    @users = initialize_grid(User.select(:id, :username, :email, :last_sign_in_at, :sign_in_count, :created_at).order(sign_in_count: :desc, id: :desc))
    render layout: 'personal'
  end

  def index
    @links = [
        {href: site_info_path, title: 'site.index.info'},
        {href: site_captcha_path, title: 'site.index.captcha'},
        {href: site_seo_path, title: 'site.index.seo'},
        {href: site_status_path, title: 'site.index.status'},
        {href: site_adverts_path, title: 'site.index.adverts'},
    ]
    render layout: 'personal'
  end

  private
  def _s_key(k)
    "#{I18n.locale}://site/#{k}"
  end
end
