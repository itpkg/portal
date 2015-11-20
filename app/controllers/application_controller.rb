class ApplicationController < ActionController::Base
  # Prevent CSRF attacks by raising an exception.
  # For APIs, you may want to use :null_session instead.
  protect_from_forgery with: :exception

  before_action :set_locale
  before_action :configure_permitted_parameters, if: :devise_controller?

  def set_locale
    I18n.locale = params[:locale] || I18n.default_locale
  end

  def default_url_options(options = {})
    {locale: I18n.locale}.merge options
  end

  protected
  def configure_permitted_parameters
    devise_parameter_sanitizer.for(:sign_up) { |u| u.permit(:username, :email, :password, :password_confirmation, :remember_me) }
    devise_parameter_sanitizer.for(:account_update) { |u| u.permit(:username, :password, :password_confirmation, :current_password, :details, :logo) }
  end

  def must_be_admin!
    authenticate_user!
    head(:forbidden) unless current_user.is_admin?
  end

  def recaptcha?

    res = Net::HTTP.post_form URI('https://www.google.com/recaptcha/api/siteverify'),
                              secret: Setting.recaptcha_secret_key,
                              response: params.fetch('g-recaptcha-response'.to_sym)

    valid = false
    if res.is_a?(Net::HTTPSuccess)
      rs = JSON.parse res.body
      valid = rs['success']
    end

    valid
  end
end
