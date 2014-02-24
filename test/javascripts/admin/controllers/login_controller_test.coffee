controller = null

module 'login controller',
  setup: ->
    Ember.run(Admin, Admin.advanceReadiness)
    controller = Admin.__container__.lookup('controller:login')

  teardown: ->
    Admin.reset()

test '#login posts to /login', ->
  expect(2)
  @spy(jQuery, 'ajax')

  controller.login('me@somesite.com', 'password')
  ok(jQuery.ajax.calledOnce)
  equal(jQuery.ajax.getCall(0).args[0].url, '/login')
