class LeaveWordsController < ApplicationController

  before_action :must_be_admin!, except: [:create]

  layout false

  def create
    lw = LeaveWord.new params.require(:leave_word).permit :content
    if lw.save
      flash[:notice] = t 'messages.success'
    else
      flash[:alert] = lw.errors.full_messages
    end
    redirect_to about_us_path
  end

  def index
    @leave_words = LeaveWord.order(id: :desc)
  end

  def destroy
    LeaveWord.destroy params[:id]
    render json: {ok: true}
  end

end