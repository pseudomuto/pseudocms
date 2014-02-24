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

  test '.validate returns the user when valid' do
    @user.save!
    assert_equal @user, User.validate(@user.email, @user.password)
  end

  test '.validate returns nil with invalid password' do
    @user.save!
    refute User.validate(@user.email, 'bad password')
  end

  test '.validate returns nil when user not found' do
    refute User.validate('who@ami.com', 'some password')
  end
end
