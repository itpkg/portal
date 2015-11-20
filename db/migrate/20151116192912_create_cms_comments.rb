class CreateCmsComments < ActiveRecord::Migration
  def change
    create_table :cms_comments do |t|
      t.belongs_to :user, index: true
      t.belongs_to :article, null: false, index: true
      t.text :content, null: false
      t.timestamps null: false
    end
  end
end
