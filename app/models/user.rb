class User < ActiveRecord::Base
  rolify

  # Include default devise modules. Others available are:
  # :confirmable, :lockable, :timeoutable and :omniauthable
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :trackable, :validatable,
         :confirmable, :lockable, :timeoutable, :omniauthable


  validates :username, presence: true
  validates :email, presence: true
  validates :logo, presence: true

  has_many :logs
  has_many :attachments
  has_many :articles, class_name: 'Cms::Article'

  before_create :set_logo



  def send_devise_notification(notification, *args)
    devise_mailer.send(notification, self, *args).deliver_later
  end


  private
  def set_logo
    unless self.logo
      self.logo = "http://gravatar.com/avatar/#{Digest::MD5.hexdigest(self.email.downcase.strip)}.png"
    end
  end
end
