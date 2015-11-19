class Cms::Article < ActiveRecord::Base
  resourcify

  belongs_to :user
  has_and_belongs_to_many :tags
  has_many :comments


  validates :title, presence: true
  validates :summary, presence: true
  validates :body, presence: true
  validates :flag, presence: true

  enum flag: [:blog, :pictures, :video]

  def can_edit?(u)
    u &&(self.user_id == u.id || u.is_admin?)
  end
end
