class Cms::CommentsController < ApplicationController
  before_action :authenticate_user!, except: [:create]
  layout 'cms'

  def index
    must_be_admin!
    @comments = initialize_grid(Cms::Comment.select(:id, :article_id, :content).order(id: :desc))
    render layout: 'personal'
  end

  def edit
    @comment = Cms::Comment.find params[:id]
    unless @comment.can_edit? current_user
      head :forbidden
    end

  end


  def destroy
    c = Cms::Comment.find params[:id]
    if c.can_edit? current_user
      flash[:notice] = t 'messages.success'
      c.destroy
    end

    redirect_to cms_article_path(c.article_id)
  end

  def update
    c = Cms::Comment.find params[:id]
    unless c.can_edit? current_user
      head :forbidden
    end

    if c.update params.require(:cms_comment).permit(:content)
      flash[:notice] = t 'messages.success'
      redirect_to cms_article_path(c.article_id, anchor: "c-#{c.id}")
    else
      flash[:alert] = c.errors.full_messages
      @comment = c
      render 'edit'
    end

  end

  def create
    c = Cms::Comment.new params.require(:cms_comment).permit(:article_id, :content)
    if recaptcha?
      c.user = current_user
      if c.save
        flash[:notice] = t 'messages.success'
      else
        flash[:alert] = c.errors.full_messages
      end
    else
      flash[:alert] = t 'messages.bad_captcha_code'
    end

    redirect_to cms_article_path(c.article_id)
  end
end
