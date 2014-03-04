require 'test_helper'

class SessionsControllerTest < ActionController::TestCase
  setup do
    @user = User.create(
      email: 'some@user.com',
      password: 'pAssword1',
      password_confirmation: 'pAssword1'
    )
  end

  test 'returns access_token  when given valid credentials' do
    params = {
      username: @user.email,
      password: @user.password
    }

    User.stubs(:new_login_token).returns('myvalue')

    post :create, params
    assert_response :success
    assert_equal 'myvalue', JSON.parse(response.body)['access_token']
  end

  test 'returns a 401 when given invalid credentials' do
    params = {
      username: @user.email,
      password: 'wrong password'
    }

    post :create, params
    assert_response :unauthorized
  end

  test 'returns a 401 when user not found' do
    params = {
      username: 'some_other@user.com',
      password: 'somePassword'
    }

    post :create, params
    assert_response :unauthorized
  end
end
