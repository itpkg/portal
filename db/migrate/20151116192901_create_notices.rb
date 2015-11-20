class CreateNotices < ActiveRecord::Migration
  def change
    create_table :notices do |t|
      t.string :lang, null: false, limit: 5, default: 'en'
      t.string :content, limit: 500, null: false
      t.datetime :created_at, null: false
    end
    add_index :notices, :lang
  end
end
