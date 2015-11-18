class AddAvatarToAttachments < ActiveRecord::Migration
  def change
    add_column :attachments, :avatar, :string, null: false
  end
end
