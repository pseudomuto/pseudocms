integration('home page')

test 'redirects to login screen when not authenticated', ->
  expect(1)

  visit('/')
  andThen ->
    equal(currentRouteName(), 'login')
