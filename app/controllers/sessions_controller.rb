class SessionsController < ApplicationController

  def create
    values = signin_params
    user = User.find_by(email: values[:username])

    if user && user.authenticate(values[:password])
      token = user.generate_token
      render json: { access_token: token, token_type: 'bearer' }
    else
      render json: { message: 'Invalid email or password.' }, status: :unauthorized
    end
  end

  private

  def signin_params
    params.permit(:username, :password)
  end
end
