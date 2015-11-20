class AddDetailsToUsers < ActiveRecord::Migration
  def change
    add_column :users, :details, :string, null: false, limit: 255, default: ' '
  end
end
