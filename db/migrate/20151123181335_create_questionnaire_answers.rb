class CreateQuestionnaireAnswers < ActiveRecord::Migration
  def change
    create_table :questionnaire_answers do |t|
      t.belongs_to :question, index: true, null: false
      t.text :content
      t.string :uid, null:false, limit:36
      t.timestamps null: false
    end
    add_index :questionnaire_answers, :uid
  end
end
