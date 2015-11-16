class Cms::Article < ActiveRecord::Base
  resourcify

  belongs_to :user
  has_and_belongs_to_many :tags
end
