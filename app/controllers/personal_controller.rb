class PersonalController < ApplicationController
  before_action :authenticate_user!
  layout 'personal'


  def logs
    @logs = initialize_grid(current_user.logs.select(:flag, :message, :created_at).order(id: :desc))
  end
end
