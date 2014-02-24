module 'admin login scenarios', QUNIT_MODULE

test '/login route', ->
  expect(2)

  visit('/login').then ->
    equal(currentPath(), 'login')
    equal(currentRouteName(), 'login')
