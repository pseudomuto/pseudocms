integration('login form')

test '/login route', ->
  expect(2)

  visit('/login')
  andThen ->
    equal(currentPath(), 'login')
    equal(currentRouteName(), 'login')

test 'displays error with invalid credentials', ->
  expect(2)

  stubRequest '/token',
    status: 401
    responseText:
      error: 'Bad Creds'

  visit('/login')
  fillIn('#email', 'some@userguy.com')
  fillIn('#password', 'pAssword1')
  click('button.submit')
  andThen ->
    equal(currentURL(), '/login')
    ok(exists('p:contains("Bad Creds")'), 'error message displayed')
