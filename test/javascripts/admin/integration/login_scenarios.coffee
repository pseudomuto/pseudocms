module 'admin login scenarios', QUNIT_MODULE

test '/login route', ->
  expect(2)

  visit('/login').then ->
    equal(currentPath(), 'login')
    equal(currentRouteName(), 'login')

test 'displays error with invalid credentials', ->
  expects(2)
  visit('/login')
  fillIn('#email', 'some@user.com')
  fillIn('#password', 'pAssword1')
  click('.submit')
  andThen ->
    equal(currentURL(), '/login')
    ok exists('p'), 'Invalid username or password.'
