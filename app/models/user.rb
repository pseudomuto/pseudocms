class User < ActiveRecord::Base
  has_secure_password
  before_save { self.email = email.downcase }

  EMAIL_REGEX = /\A[\w+\-.]+@[a-z\d\-.]+\.[a-z]+\z/i

  validates :email,
    presence: true,
    format: { with: EMAIL_REGEX },
    uniqueness: { case_sensitive: false }

  def self.validate(email, password)
    user = User.find_by(email: email)
    return nil unless user

    user.authenticate(password) ? user : nil
  end
end
