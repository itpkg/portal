class Cms::Tag < ActiveRecord::Base
  resourcify

  validates :name, presence: true
  validates :lang, presence: true
  validates :name, uniqueness: {scope: :lang}

  has_and_belongs_to_many :articles
end
