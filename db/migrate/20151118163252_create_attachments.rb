class CreateAttachments < ActiveRecord::Migration
  def change
    create_table :attachments do |t|
      t.belongs_to :user, null: false, index: true

      t.string :title, null: false
      t.string :content_type, null: false
      t.string :ext, limit: 5, null: false
      t.integer :size, null: false
      t.integer :by_use, null:false, default: 0

      t.timestamps null: false
    end

    add_index :attachments, :title
    add_index :attachments, :content_type
    add_index :attachments, :ext
  end
end
