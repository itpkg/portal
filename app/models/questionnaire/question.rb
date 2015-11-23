class Questionnaire::Question < ActiveRecord::Base
  belongs_to :report
  validates :name, presence: true
  validates :flag, presence: true


  enum flag: [:single, :multi, :y_n]
end
