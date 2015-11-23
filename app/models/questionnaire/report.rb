class Questionnaire::Report < ActiveRecord::Base
  validates :title, presence: true
  validates :summary, presence: true
end
