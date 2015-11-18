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

  end

  def index
    @links = [
        {href: site_info_path, title: 'site.index.info'},
        {href: site_captcha_path, title: 'site.index.captcha'},
        {href: site_seo_path, title: 'site.index.seo'},
        {href: site_status_path, title: 'site.index.status'},
        {href: site_adverts_path, title: 'site.index.adverts'},
        {href: site_users_path, title: 'site.index.users'},
    ]
  end

  private
  def _s_key(k)
    "#{I18n.locale}://site/#{k}"
  end
end
