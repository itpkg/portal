class CreateQuestionnaireQuestions < ActiveRecord::Migration
  def change
    create_table :questionnaire_questions do |t|
      t.belongs_to :report, index: true, null: false
      t.string :name, limit:255, null:false
      t.integer :flag, null:false, default:0
      t.string :args, limit:800
      t.timestamps null: false
    end
  end
end
