class AddLogoToUsers < ActiveRecord::Migration
  def change
    add_column :users, :logo, :string, null: false, limit: 255
  end
end
