source 'https://rubygems.org'
ruby '2.1.1'

gem 'rails', '4.0.4'
gem 'unicorn-rails'
gem 'foreman'

gem 'ember-rails'
gem 'ember-source'
gem 'ember_script-rails', github: 'ghempton/ember-script-rails'
gem 'jquery-rails'
gem 'font-awesome-rails'
gem 'bcrypt-ruby'
gem 'sprockets_better_errors'

group :assets do
  gem 'sass-rails', '~> 4.0.0'
  gem 'uglifier', '>= 1.3.0'
  gem 'coffee-rails', '~> 4.0.0'
end

group :development, :test do
  gem 'mysql2'
  gem 'qunit-rails'
  gem 'mocha', require: false
  gem 'pry'
end

group :production do
  gem 'rails_12factor'
  gem 'pg'
end
