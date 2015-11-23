# encoding: UTF-8
# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# Note that this schema.rb definition is the authoritative source for your
# database schema. If you need to create the application database on another
# system, you should be using db:schema:load, not running all the migrations
# from scratch. The latter is a flawed and unsustainable approach (the more migrations
# you'll amass, the slower it'll run and the greater likelihood for issues).
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 20151123181335) do

  # These are extensions that must be enabled in order to support this database
  enable_extension "plpgsql"

  create_table "attachments", force: :cascade do |t|
    t.integer  "user_id",      null: false
    t.string   "title",        null: false
    t.string   "content_type", null: false
    t.integer  "size",         null: false
    t.datetime "created_at",   null: false
    t.datetime "updated_at",   null: false
    t.string   "avatar",       null: false
  end

  add_index "attachments", ["content_type"], name: "index_attachments_on_content_type", using: :btree
  add_index "attachments", ["title"], name: "index_attachments_on_title", using: :btree
  add_index "attachments", ["user_id"], name: "index_attachments_on_user_id", using: :btree

  create_table "cms_articles", force: :cascade do |t|
    t.integer  "user_id",                                null: false
    t.string   "title",      limit: 255,                 null: false
    t.string   "summary",    limit: 800,                 null: false
    t.string   "logo",       limit: 255
    t.boolean  "top",                    default: false, null: false
    t.text     "body",                                   null: false
    t.string   "lang",       limit: 5,   default: "en",  null: false
    t.integer  "visits",                 default: 0,     null: false
    t.datetime "created_at",                             null: false
    t.datetime "updated_at",                             null: false
  end

  add_index "cms_articles", ["lang"], name: "index_cms_articles_on_lang", using: :btree
  add_index "cms_articles", ["user_id"], name: "index_cms_articles_on_user_id", using: :btree

  create_table "cms_articles_tags", id: false, force: :cascade do |t|
    t.integer "article_id", null: false
    t.integer "tag_id",     null: false
  end

  add_index "cms_articles_tags", ["article_id"], name: "index_cms_articles_tags_on_article_id", using: :btree
  add_index "cms_articles_tags", ["tag_id"], name: "index_cms_articles_tags_on_tag_id", using: :btree

  create_table "cms_comments", force: :cascade do |t|
    t.integer  "user_id"
    t.integer  "article_id", null: false
    t.text     "content",    null: false
    t.datetime "created_at", null: false
    t.datetime "updated_at", null: false
  end

  add_index "cms_comments", ["article_id"], name: "index_cms_comments_on_article_id", using: :btree
  add_index "cms_comments", ["user_id"], name: "index_cms_comments_on_user_id", using: :btree

  create_table "cms_tags", force: :cascade do |t|
    t.string   "name",                      null: false
    t.integer  "visits",     default: 0,    null: false
    t.datetime "created_at",                null: false
    t.datetime "updated_at",                null: false
    t.string   "lang",       default: "en", null: false
  end

  add_index "cms_tags", ["lang", "name"], name: "index_cms_tags_on_lang_and_name", unique: true, using: :btree
  add_index "cms_tags", ["lang"], name: "index_cms_tags_on_lang", using: :btree

  create_table "friend_links", force: :cascade do |t|
    t.string   "home",       limit: 255, null: false
    t.string   "name",       limit: 255, null: false
    t.string   "logo",       limit: 255
    t.datetime "created_at",             null: false
  end

  add_index "friend_links", ["home"], name: "index_friend_links_on_home", unique: true, using: :btree

  create_table "leave_words", force: :cascade do |t|
    t.string   "content",    limit: 800, null: false
    t.datetime "created_at",             null: false
  end

  create_table "logs", force: :cascade do |t|
    t.integer  "user_id",                            null: false
    t.string   "message",    limit: 255,             null: false
    t.integer  "flag",                   default: 0, null: false
    t.datetime "created_at",                         null: false
  end

  create_table "notices", force: :cascade do |t|
    t.string   "lang",       limit: 5,   default: "en", null: false
    t.string   "content",    limit: 500,                null: false
    t.datetime "created_at",                            null: false
  end

  add_index "notices", ["lang"], name: "index_notices_on_lang", using: :btree

  create_table "questionnaire_answers", force: :cascade do |t|
    t.integer  "question_id",            null: false
    t.text     "content"
    t.string   "uid",         limit: 36, null: false
    t.datetime "created_at",             null: false
    t.datetime "updated_at",             null: false
  end

  add_index "questionnaire_answers", ["question_id"], name: "index_questionnaire_answers_on_question_id", using: :btree
  add_index "questionnaire_answers", ["uid"], name: "index_questionnaire_answers_on_uid", using: :btree

  create_table "questionnaire_questions", force: :cascade do |t|
    t.integer  "report_id",                          null: false
    t.string   "name",       limit: 255,             null: false
    t.integer  "flag",                   default: 0, null: false
    t.string   "args",       limit: 800
    t.datetime "created_at",                         null: false
    t.datetime "updated_at",                         null: false
  end

  add_index "questionnaire_questions", ["report_id"], name: "index_questionnaire_questions_on_report_id", using: :btree

  create_table "questionnaire_reports", force: :cascade do |t|
    t.string   "title",      limit: 255, null: false
    t.string   "summary",    limit: 800, null: false
    t.datetime "created_at",             null: false
    t.datetime "updated_at",             null: false
  end

  create_table "roles", force: :cascade do |t|
    t.string   "name"
    t.integer  "resource_id"
    t.string   "resource_type"
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "roles", ["name", "resource_type", "resource_id"], name: "index_roles_on_name_and_resource_type_and_resource_id", using: :btree
  add_index "roles", ["name"], name: "index_roles_on_name", using: :btree

  create_table "settings", force: :cascade do |t|
    t.string   "var",                   null: false
    t.text     "value"
    t.integer  "thing_id"
    t.string   "thing_type", limit: 30
    t.datetime "created_at"
    t.datetime "updated_at"
  end

  add_index "settings", ["thing_type", "thing_id", "var"], name: "index_settings_on_thing_type_and_thing_id_and_var", unique: true, using: :btree

  create_table "users", force: :cascade do |t|
    t.string   "email",                              default: "",  null: false
    t.string   "encrypted_password",                 default: "",  null: false
    t.string   "reset_password_token"
    t.datetime "reset_password_sent_at"
    t.datetime "remember_created_at"
    t.integer  "sign_in_count",                      default: 0,   null: false
    t.datetime "current_sign_in_at"
    t.datetime "last_sign_in_at"
    t.inet     "current_sign_in_ip"
    t.inet     "last_sign_in_ip"
    t.string   "confirmation_token"
    t.datetime "confirmed_at"
    t.datetime "confirmation_sent_at"
    t.string   "unconfirmed_email"
    t.integer  "failed_attempts",                    default: 0,   null: false
    t.string   "unlock_token"
    t.datetime "locked_at"
    t.datetime "created_at",                                       null: false
    t.datetime "updated_at",                                       null: false
    t.string   "username",               limit: 255,               null: false
    t.string   "details",                limit: 255, default: " ", null: false
    t.string   "logo",                   limit: 255,               null: false
  end

  add_index "users", ["confirmation_token"], name: "index_users_on_confirmation_token", unique: true, using: :btree
  add_index "users", ["email"], name: "index_users_on_email", unique: true, using: :btree
  add_index "users", ["reset_password_token"], name: "index_users_on_reset_password_token", unique: true, using: :btree
  add_index "users", ["unlock_token"], name: "index_users_on_unlock_token", unique: true, using: :btree
  add_index "users", ["username"], name: "index_users_on_username", using: :btree

  create_table "users_roles", id: false, force: :cascade do |t|
    t.integer "user_id"
    t.integer "role_id"
  end

  add_index "users_roles", ["user_id", "role_id"], name: "index_users_roles_on_user_id_and_role_id", using: :btree

end
