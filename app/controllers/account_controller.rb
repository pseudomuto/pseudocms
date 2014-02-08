class AccountController < ApplicationController

  def index
    @user = User.new
  end

  def create
    @user = User.new(signup_params)
    if @user.save
      redirect_to admin_path
    else
      flash.now[:error] = "There was an error creating your account"
      render :index
    end
  end

  private

  def signup_params
    params.require(:user).permit(:email, :password, :password_confirmation)
  end
end
