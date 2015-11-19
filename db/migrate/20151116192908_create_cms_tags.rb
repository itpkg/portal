class CreateCmsTags < ActiveRecord::Migration
  def change
    create_table :cms_tags do |t|
      t.string :name, null: false
      t.integer :visits, null:false, default:0
      t.timestamps null: false
    end


    add_index :cms_tags, :name, unique: true


    create_table :cms_articles_tags, id: false do |t|
      t.belongs_to :article, index: true, null: false
      t.belongs_to :tag, index: true, null: false
    end

  end
end
