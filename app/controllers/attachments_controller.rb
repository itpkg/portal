class AttachmentsController < ApplicationController
  before_action :authenticate_user!, except:[:show]
  layout nil

  def show

  end


  def create
    a = Attachment.new user:current_user
    a.read! params[:file]

    if a.save
      flash[:notice] = t 'messages.success'
    else
      flash[:alert] = t 'messages.failed'
    end
    redirect_to attachments_path
  end

  def destroy
    a = Attachment.find params[:id]
    if a && a.user_id == current_user.id && a.by_use == 0
      a.destroy
      flash[:notice] = t 'messages.success'
    else
      flash[:alert] = t 'messages.failed'
    end
    redirect_to attachments_path

  end

  def index
    if current_user.is_admin?
      @items = initialize_grid Attachment.order(id: :desc)
    else
      @items = initialize_grid current_user.attachments
    end
    render layout: 'personal'
  end
end