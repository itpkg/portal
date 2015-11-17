class SiteController < ApplicationController
  before_action :must_be_admin!

  layout false

  def info

  end

  def captcha

  end

  def seo

  end

  def status

  end

  def adverts

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

end
