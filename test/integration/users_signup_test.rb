require 'test_helper'

class UsersSignupTest < ActionDispatch::IntegrationTest
  # fixtures :users

  test "sign up for an account" do
    get '/'
    assert_equal 200, status

    post '/sign_up', email: 'pseudomuto@pseudocms.com', password: 'some_password'

  end
end
