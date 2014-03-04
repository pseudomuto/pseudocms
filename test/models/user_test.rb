require 'test_helper'

class UserTest < ActiveSupport::TestCase
  setup do
    @user = User.new(
      email: 'valid@email.com',
      password: 'pAssword1',
      password_confirmation: 'pAssword1'
    )
  end

  test '#email is required' do
    @user.email = ''
    refute @user.valid?
  end

  test '#email must be a valid email address' do
    %w(some_thing close_but@notanemail).each do |bad_value|
      @user.email = bad_value
      refute @user.valid?
    end
  end

  test '#email must be unique' do
    existing_user = @user.dup
    existing_user.save!

    refute @user.save
  end

  test '#password is required' do
    @user.password = ''
    @user.password_confirmation = ''
    refute @user.valid?
  end

  test '#password_confirmation must match' do
    @user.password_confirmation = 'mismatch'
    refute @user.valid?
  end

  test '.generate_token returns and stores encrypted token' do
    @user.save!
    User.stubs(:encrypt).returns('').once
    assert_not_nil @user.generate_token
  end
end
