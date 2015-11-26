class Cms::UsersController < ApplicationController
  layout 'cms'

  def show
    user=User.select(:id, :username).find params[:id]
    if user.is_admin?
      head :not_found
    else
      @articles = user.articles.select(:id, :summary, :title, :logo).order(id: :desc).page params[:page]
      @title = t 'cms.users.show.title', name: user.username
      render 'cms/articles/index'
    end
  end

  def index
    @users = User.select(:id, :username, :details, :email, :logo).order(sign_in_count: :desc).page(params[:page]).per 30
  end
end