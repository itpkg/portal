class Questionnaire::Report < ActiveRecord::Base
  validates :title, presence: true
  validates :summary, presence: true

  has_many :questions, class_name: 'Questionnaire::Question'
end
