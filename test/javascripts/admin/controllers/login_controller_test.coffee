controller = null
spy = null

module 'login controller',
  setup: ->
    controller = Admin.__container__.lookup('controller:login')
    spy = sinon.spy(jQuery, 'ajax')

  teardown: ->
    jQuery.ajax.restore()

test '#reset sets default property values', ->
  controller.setProperties
    identification: 'some@user.com'
    password: 'pAssword1'
    errorMessage: 'some message'

  controller.reset()
  equal(controller.get('errorMessage'), null)
  equal(controller.get('identification'), null)
  equal(controller.get('password'), null)

test '#authenticate posts to /token', ->
  stubRequest '/token',
    responseText:
      access_token: 'some_token'
      token_type: 'bearer'

  controller.setProperties
    identification: 'some@user.com'
    password: 'pAssword1'

  controller.send('authenticate')
  ok(spy.calledOnce)

test '#authenticate skips the call when credentials are not supplied', ->
  controller.send('authenticate')
  equal(spy.called, false)
