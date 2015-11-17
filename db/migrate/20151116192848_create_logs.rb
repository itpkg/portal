class CreateLogs < ActiveRecord::Migration
  def change
    create_table :logs do |t|
      t.belongs_to :user, null: false
      t.string :message, null: false, limit: 255
      t.integer :flag, null: false, default: 0
      t.datetime :created_at, null: false
    end
  end
end
