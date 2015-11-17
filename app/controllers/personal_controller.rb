class PersonalController < ApplicationController
  before_action :authenticate_user!

  layout false


  def logs

  end

  def index
    @links = []

    if current_user.is_admin?
      @links << {href:site_path, title:'links.personal.site'}
    end

    @links << {href:personal_logs_path, title:'links.personal.logs'}
    render layout: 'dashboard'
  end
end
