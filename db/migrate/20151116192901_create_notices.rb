class CreateNotices < ActiveRecord::Migration
  def change
    create_table :notices do |t|
      t.string :message, limit:255, null:false
      t.timestamps null: false
    end
  end
end
