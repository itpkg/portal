class Log < ActiveRecord::Base
  belongs_to :user

  enum flag: [:info, :error, :notice]

  validates :message, presence: true
  validates :user_id, presence: true

  before_create :set_created_at


  private
  def set_created_at
    self.created_at = Time.now
  end
end
