class CreateFriendLinks < ActiveRecord::Migration
  def change
    create_table :friend_links do |t|
      t.string :home, null:false, limit:255
      t.string :name, null:false, limit:255
      t.string :logo, limit:255
      t.datetime :created_at, null:false
    end
    add_index :friend_links, :home, unique: true
  end
end
