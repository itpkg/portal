module ApplicationHelper

  def to_html(md)
    Rails.application.config.md2hm.render(md) if md
  end

  def to_json(o)
    JSON.pretty_generate o
  end

  def paginate objects, options = {}
    options.reverse_merge!(theme: 'twitter-bootstrap-3')

    super(objects, options)
  end
end
