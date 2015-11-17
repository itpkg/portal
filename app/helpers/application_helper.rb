module ApplicationHelper

  def site_info(key)
    Setting["#{I18n.locale}://#{key}"]
  end

  def paginate objects, options = {}
    options.reverse_merge!(theme: 'twitter-bootstrap-3')

    super(objects, options)
  end
end
