class NoticesController < ApplicationController
  before_action :must_be_admin!

  layout false

  def index
    @notice = Notice.new
    @notices = Notice.where(lang: I18n.locale).order(id: :desc).limit(50)
  end

  def create
    n = Notice.new params.require(:notice).permit(:content)
    n.lang = I18n.locale
    n.save
    render json: {ok: true}
  end

  def destroy
    Notice.destroy params[:id]
    render json: {ok: true}
  end
end
