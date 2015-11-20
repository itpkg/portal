class CreateCmsArticles < ActiveRecord::Migration
  def change
    create_table :cms_articles do |t|
      t.belongs_to :user, null: false, index: true
      t.string :title, null: false, limit: 255
      t.string :summary, null: false, limit: 800
      t.boolean :top, null:false, default:false
      t.text :body, null: false
      t.string :lang, null: false, limit: 5, default: 'en'
      t.integer :visits, null:false, default:0
      t.timestamps null: false
    end

    add_index :cms_articles, :lang

  end
end
