integration('home page')

test 'displays welcome message', ->
  expect(1)
  visit('/')
  andThen ->
    ok exists('h1:contains("Welcome to the Admin")')
