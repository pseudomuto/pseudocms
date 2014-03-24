integration('home page')

test 'redirects to login screen when not authenticated', ->
  expect(1)

  Ember.run ->
    visit('/')
    andThen ->
      equal(currentRouteName(), 'login')
