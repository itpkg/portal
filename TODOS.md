DEVEL
---

## Todos
 * cms

## Notes
### Rails
 * rails g controller home index --no-assets --helper
 * <%= f.input :flag, as: :radio_buttons, collection: Cms::Article.flags.map { |k, _| k.to_sym }, label_method: lambda { |f| t "simple_form.labels.cms_article.flag_#{f}" }, item_wrapper_class: 'radio-inline' %>
