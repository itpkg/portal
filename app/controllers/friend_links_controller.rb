class FriendLinksController < ApplicationController
  before_action :must_be_admin!
  layout false

  def index
    @friend_link = FriendLink.new
  end

  def create
    fl = FriendLink.new params.require(:friend_link).permit(:name, :home, :logo)
    if fl.save
      render json: {ok: true}
    else
      render json: {ok: false, data: fl.errors.full_messages}
    end
  end

  def destroy
    FriendLink.destroy params[:id]
    render json: {ok: true}
  end
end