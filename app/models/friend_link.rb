class FriendLink < ActiveRecord::Base
  validates :home, presence: true, uniqueness: true
  validates :name, presence: true
end
