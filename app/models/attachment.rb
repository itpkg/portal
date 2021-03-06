class Attachment < ActiveRecord::Base
  mount_uploader :avatar, AvatarUploader

  belongs_to :user

  def image?
    self.content_type.start_with? 'image'
  end

  def read!(file)
    name = file.original_filename
    self.content_type = file.content_type
    self.title = name
    self.size = File.size file.tempfile
    self.avatar = file
  end

end
