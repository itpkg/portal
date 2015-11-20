class Cms::TagsController < ApplicationController
  before_action :must_be_admin!, except: [:show]
  layout 'personal'

  def show
    tag=Cms::Tag.find params[:id]
    tag.update_column :visits, tag.visits+1
    @articles = tag.articles.select(:id, :summary, :title).order(id: :desc).page params[:page]
    @title = t 'cms.tags.show.title', name: tag.name
    render 'cms/articles/index', layout: 'cms'
  end

  def new
    @tag = Cms::Tag.new
  end

  def create
    begin
      Cms::Tag.create _params
      flash[:notice] = t 'messages.success'
    rescue => e
      flash[:alert] = e.to_s
    end
    redirect_to cms_tags_path

  end

  def update
    begin
      tag = Cms::Tag.find params[:id]
      tag.update _params
      flash[:notice] = t 'messages.success'
    rescue => e
      flash[:alert] = e.to_s
    end
    redirect_to cms_tags_path
  end

  def edit
    @tag = Cms::Tag.find params[:id]
  end

  def index
    @tags = initialize_grid(Cms::Tag.select(:id, :name, :visits).order(id: :desc))
  end

  def destroy
    begin
      tag = Cms::Tag.find params[:id]
      if tag.articles.count == 0
        tag.destroy
        flash[:notice] = t 'messages.success'
      else
        raise t('messages.in_use')
      end

    rescue => e
      flash[:alert] = e.to_s
    end
    redirect_to cms_tags_path
  end

  private
  def _params
    params.require(:cms_tag).permit(:name)
  end

end
