Pseudocms::Application.routes.draw do
  root to: 'main#index'

  namespace :admin do
    root to: 'main#index'
  end
end
