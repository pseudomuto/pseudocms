integration('login form')

test 'redirects to root on successful login', ->
  stubRequest '/token',
    status: 200,
    responseText:
      access_token: 'some_token'
      token_type: 'bearer'

  Ember.run ->
    visit('/login')
    fillIn('#email', 'some@userguy.com')
    fillIn('#password', 'pAssword1')
    click('button.submit')
    andThen ->
      equal(currentRouteName(), 'index')
      ok(exists('#profile a:contains("logout")'), 'adds logout option')

test 'displays error with invalid credentials', ->
  expect(2)

  stubRequest '/token',
    status: 401
    responseText:
      message: 'Bad Creds'

  Ember.run ->
    visit('/login')
    fillIn('#email', 'some@userguy.com')
    fillIn('#password', 'pAssword1')
    click('button.submit')
    andThen ->
      equal(currentRouteName(), 'login')
      ok(exists('p:contains("Bad Creds")'), 'error message displayed')
