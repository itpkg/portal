class AddLangToCmsTags < ActiveRecord::Migration
  def change
    remove_index :cms_tags, :name

    add_column :cms_tags, :lang, :string, null:false, default: 'en'
    add_index :cms_tags, :lang
    add_index :cms_tags, [:lang, :name], unique: true
  end
end
