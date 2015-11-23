DEVEL
---

## Todos
 * facebook/qq/weibo oauth
 * cache page

## Notes
### bootstrap

 * xs：extra small 特别窄屏幕，默认指浏览器像素宽度小于768px
 * sm：small 窄屏幕，默认指浏览器像素宽度大于等于768px
 * md：middle 中等宽度屏幕，默认值指浏览器像素宽度大于等于992px
 * lg：large 大屏幕，默认值指浏览器像素宽度大于等于1200px

###
 * 格式化cap模板代码会引入puma.rb bug
### Rails
 * rails g controller home index --no-assets --helper
 * <%= f.input :flag, as: :radio_buttons, collection: Cms::Article.flags.map { |k, _| k.to_sym }, label_method: lambda { |f| t "simple_form.labels.cms_article.flag_#{f}" }, item_wrapper_class: 'radio-inline' %>
