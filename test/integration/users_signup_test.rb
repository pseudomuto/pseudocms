require 'test_helper'

class UsersSignupTest < ActionDispatch::IntegrationTest
  attr_reader :params

  setup do
    @params = {
      user: {
        email: 'pseudomuto@pseudocms.com',
        password: 'some_password',
        password_confirmation: 'some_password'
      }
    }

    get new_user_path
    assert_equal 200, status
  end

  test "sign up with valid account details" do
    post_via_redirect sign_up_path, params
    assert_equal admin_path, path
  end

  test "sign up without entering email address" do
    params[:user][:email] = ''
    post_via_redirect sign_up_path, params

    assert_equal new_user_path, path
  end

  test "sign up without entering a password" do
    params[:user][:password] = ''
    post_via_redirect sign_up_path, params

    assert_equal new_user_path, path
  end

  test "sign up without confirming the password" do
    params[:user][:password_confirmation] = ''
    post_via_redirect sign_up_path, params

    assert_equal new_user_path, path
  end

  test "sign up with invalid password confirmation" do
    params[:user][:password_confirmation] = 'something_else'
    post_via_redirect sign_up_path, params

    assert_equal new_user_path, path
  end
end
