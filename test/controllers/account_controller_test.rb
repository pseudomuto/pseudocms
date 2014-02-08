require 'test_helper'

class AccountControllerTest < ActionController::TestCase

  test 'GET #index should present the sign up form' do
    get :index
    assert assigns(:user)
  end

  test 'POST #create should create a new user' do
    assert_difference('User.count') do
      post :create, user: {
        email: 'pseudomuto@pseudocms.com',
        password: 'my password',
        password_confirmation: 'my password'
      }
    end

    assert_redirected_to admin_path
  end

  test 'POST #create with bad data renders sign up form' do
    assert_no_difference('User.count') do
      post :create, user: {
        email: 'pseudomuto@pseudocms',
        password: 'my password',
        password_confirmation: 'my password'
      }
    end

    assert_template :index
    assert_not_nil flash[:error]
  end
end
