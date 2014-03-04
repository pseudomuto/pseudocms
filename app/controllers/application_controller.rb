class ApplicationController < ActionController::Base
  protect_from_forgery

  def authenticate
    head :unauthorized and return unless current_user
  end

  def current_user
    @current_user ||= access_token && User.find_by(login_token: User.encrypt(access_token))
  end

  def access_token
    @access_token ||= request.authorization && request.authorization.split(' ').last
  end
end
