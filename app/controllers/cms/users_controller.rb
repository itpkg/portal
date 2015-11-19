class Cms::UsersController < ApplicationController
  layout 'cms'

  def show
    user=User.select(:id, :username).find params[:id]
    @articles = user.articles.select(:id, :summary, :flag, :title).order(id: :desc).page params[:page]
    @title = t 'cms.users.index.title', name: user.username
    render 'cms/articles/index', layout: 'cms'
  end
end