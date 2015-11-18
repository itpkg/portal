class HomeController < ApplicationController
  def index
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
