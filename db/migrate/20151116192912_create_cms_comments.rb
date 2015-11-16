class CreateCmsComments < ActiveRecord::Migration
  def change
    create_table :cms_comments do |t|
      t.belongs_to :user, null:false, index:true
      t.text :content, null:false
      t.timestamps null: false
    end
  end
end
