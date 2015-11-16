class Cms::Tag < ActiveRecord::Base
  resourcify

  has_and_belongs_to_many :articles
end
