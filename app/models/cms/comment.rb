class Cms::Comment < ActiveRecord::Base
  resourcify

  validates :content, presence: true

  belongs_to :user
  belongs_to :article


  def can_edit?(u)
    u && (self.user_id == u.id || u.is_admin?)
  end
end
