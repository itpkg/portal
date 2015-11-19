class Cms::ArticlesController < ApplicationController
  layout 'cms'

  def index
    @articles = Cms::Article.select(:id, :summary, :flag, :title).order(id: :desc).page params[:page]
    @title = t 'cms.articles.index.title'
  end

  def show
    @article = Cms::Article.find params[:id]
    @article.update_column :visits, @article.visits+1
  end
end
