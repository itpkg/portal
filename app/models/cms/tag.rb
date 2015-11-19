class Cms::Tag < ActiveRecord::Base
  resourcify

  validates :name, presence: true

  has_and_belongs_to_many :articles
end
