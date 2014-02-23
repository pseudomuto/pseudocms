require 'test_helper'

class MainControllerTest < ActionController::TestCase
  test 'should GET index' do
    get :index
    assert_response :success
    assert_template :index
    assert_template layout: 'layouts/application'
  end
end
