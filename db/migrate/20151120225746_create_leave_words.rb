class CreateLeaveWords < ActiveRecord::Migration
  def change
    create_table :leave_words do |t|
      t.string :content, null: false, limit: 800
      t.datetime :created_at, null: false
    end
  end
end
