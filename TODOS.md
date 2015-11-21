DEVEL
---

## Todos
 * facebook/qq/weibo oauth
 * cache page

## Notes
###
 * 格式化cap模板代码会引入puma.rb bug
### Rails
 * rails g controller home index --no-assets --helper
 * <%= f.input :flag, as: :radio_buttons, collection: Cms::Article.flags.map { |k, _| k.to_sym }, label_method: lambda { |f| t "simple_form.labels.cms_article.flag_#{f}" }, item_wrapper_class: 'radio-inline' %>
