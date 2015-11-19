class Notice < ActiveRecord::Base
  validates :content, presence: true

  before_create :set_created_at


  private
  def set_created_at
    self.created_at = Time.now
  end
end
