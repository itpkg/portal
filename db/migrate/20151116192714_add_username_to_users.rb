class AddUsernameToUsers < ActiveRecord::Migration
  def change
    add_column :users, :username, :string, null: false, limit: 255
    add_index :users, :username
  end
end
