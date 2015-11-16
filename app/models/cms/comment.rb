class Cms::Comment < ActiveRecord::Base
  resourcify

  belongs_to :user
end
