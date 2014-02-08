require 'test_helper'

class UserTest < ActiveSupport::TestCase

  def setup
    @user = User.new(
      email: 'pseudomuto@pseudocms.com',
      password: 'simplePassWD',
      password_confirmation: 'simplePassWD'
    )
  end

  test 'email is required' do
    @user.email = ''
    refute @user.valid?
  end

  test 'email must be a valid email address' do
    ['some thing', 'close_bit@notanemail'].each do |email|
      @user.email = email
      refute @user.valid?
    end
  end

  test 'email must be unique' do
    other_user = @user.dup
    other_user.save

    refute @user.save
  end

  test 'email uniqueness ignores case' do
    other_user = @user.dup
    other_user.email = other_user.email.upcase
    other_user.save

    refute @user.save
  end

  test 'email must be less that 120 characters' do
    @user.email = "#{('a' * 100)}@{'b' * 48}.com"
    refute @user.valid?
  end

  test 'password is required' do
    @user.password = ''
    @user.password_confirmation = ''
    refute @user.valid?
  end

  test 'password confirmation much match' do
    @user.password_confirmation = 'mismatch'
    refute @user.valid?
  end
end
