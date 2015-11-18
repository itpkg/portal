# RailsSettings Model
class Setting < RailsSettings::CachedSettings

  def self.set_site_info(k, v)
    self[self._site_key(k)] = v
  end

  def self.get_site_info(k)
    self[self._site_key(k)]
  end

  private
  def self._site_key(k)
    "#{I18n.locale}://site/#{k}"
  end
end
