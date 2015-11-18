class PersonalController < ApplicationController
  before_action :authenticate_user!
  layout 'personal'


  def logs
    @logs = initialize_grid(Log.select(:flag, :message, :created_at).order(id: :desc))
  end
end
