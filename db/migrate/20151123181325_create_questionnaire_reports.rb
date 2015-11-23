class CreateQuestionnaireReports < ActiveRecord::Migration
  def change
    create_table :questionnaire_reports do |t|
      t.string :title, null:false, limit:255
      t.string :summary, null:false, limit:800
      t.timestamps null: false
    end
    
  end
end
