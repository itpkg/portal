class Cms::UsersController < ApplicationController
  layout 'cms'

  def show
    user=User.select(:id, :username).find params[:id]
    @articles = user.articles.select(:id, :summary, :flag, :title).order(id: :desc).page params[:page]
    @title = t 'cms.users.show.title', name: user.username
    render 'cms/articles/index'
  end

  def index
    @users = User.select(:id, :username, :details, :email).order(sign_in_count: :desc).page(params[:page]).per 30
  end
end