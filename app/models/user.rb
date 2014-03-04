class User < ActiveRecord::Base
  has_secure_password
  before_save { self.email = email.downcase }

  EMAIL_REGEX = /\A[\w+\-.]+@[a-z\d\-.]+\.[a-z]+\z/i

  validates :email,
    presence: true,
    format: { with: EMAIL_REGEX },
    uniqueness: { case_sensitive: false }

  def self.new_login_token
    SecureRandom.urlsafe_base64
  end

  def self.encrypt(token)
    Digest::SHA1.hexdigest(token.to_s)
  end

  def generate_token
    token = User.new_login_token
    update_attribute(:login_token, User.encrypt(token))
    token
  end
end
