class Admin.LoginController extends Ember.ObjectController
  isProcessing: false
  loginFailed: false

  login: (email, password) ->
    @setProperties(
      isProcessing: true
    )

    request = $.post('/login', { email: email, password: password })
    request.then @_success.bind(this), @_failure.bind(this)
  
  _success: (result) ->
    @_reset()
    @transitionTo('/')

  _failure: (result) ->
    @_reset()
    @set('loginFailed', true)

  _reset: ->
    @setProperties(
      isProcessing: false
    )
