class Cms::ArticlesController < ApplicationController
  layout 'cms'

  before_action :authenticate_user!, except: [:index, :show]

  def index
    @articles = Cms::Article.select(:id, :summary, :title).where(lang: I18n.locale).order(id: :desc).page params[:page]
    @title = t 'cms.articles.index.title'
  end

  def show
    @article = Cms::Article.find params[:id]
    @article.update_column :visits, @article.visits+1
    @comment = Cms::Comment.new article: @article
  end

  def new
    @article = Cms::Article.new
  end

  def create
    a = Cms::Article.new _params

    a.user = current_user
    a.lang = I18n.locale
    if a.save
      flash[:notice] = t 'messages.success'
      redirect_to(cms_article_path(a))
    else
      flash[:alert] = a.errors.full_messages
      @article = a
      render 'new'
    end

  end

  def edit
    @article = Cms::Article.find params[:id]
    unless @article.can_edit?(current_user)
      head :forbidden
    end
  end

  def update
    a = Cms::Article.find params[:id]

    unless a.can_edit?(current_user)
      head :forbidden
    end
    begin
      if a.update _params
        flash[:notice] = t 'messages.success'
        redirect_to(cms_article_path(a)) and return
      end
    rescue => e
      flash[:alert] = e.to_s
    end
    @article = a
    render 'edit'
  end

  def destroy
    a = Cms::Article.find params[:id]
    if a.can_edit?(current_user)
      a.destroy
      redirect_to cms_articles_path
    else
      head :forbidden
    end
  end

  private
  def _params
    params.require(:cms_article).permit(:title, :summary, :body, tag_ids: [])
  end
end
