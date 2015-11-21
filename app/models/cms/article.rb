class Cms::Article < ActiveRecord::Base
  resourcify

  belongs_to :user
  has_and_belongs_to_many :tags
  has_many :comments


  validates :title, presence: true
  validates :summary, presence: true
  validates :body, presence: true

  before_create :set_logo

  def can_edit?(u)
    u &&(self.user_id == u.id || u.is_admin?)
  end

  private
  def set_logo
    doc = Nokogiri::HTML(self.body)
    img = doc.xpath('//img').first
    if img
      self.logo = img.attr 'src'
    end
  end
end
