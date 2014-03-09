integration('login form')

test 'redirects to root on successful login', ->
  stubRequest '/token',
    status: 200,
    responseText:
      access_token: 'some_token'
      token_type: 'bearer'

  visit('/login')
  fillIn('#email', 'some@userguy.com')
  fillIn('#password', 'pAssword1')
  click('button.submit')
  andThen ->
    equal(currentURL(), '/')
    ok(exists('nav a:contains("logout")'), 'adds logout option')

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
