class CreateAttachments < ActiveRecord::Migration
  def change
    create_table :attachments do |t|
      t.belongs_to :user, null: false, index: true
      t.string :title, null: false
      t.string :content_type, null: false
      t.integer :size, null: false

      t.timestamps null: false
    end

    add_index :attachments, :title
    add_index :attachments, :content_type
  end
end
