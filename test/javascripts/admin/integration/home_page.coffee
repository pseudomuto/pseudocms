module 'home page', QUNIT_MODULE

test 'displays welcome message', ->
  visit('/').then ->
    ok exists('h1'), 'Welcome to the Admin!'
