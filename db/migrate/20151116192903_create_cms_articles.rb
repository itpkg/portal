class CreateCmsArticles < ActiveRecord::Migration
  def change
    create_table :cms_articles do |t|
      t.belongs_to :user, null: false, index: true
      t.string :title, null: false, limit: 255
      t.string :summary, null: false, limit: 800
      t.text :body, null: false
      t.timestamps null: false
    end

  end
end
