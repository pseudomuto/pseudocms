class Admin.LoginController extends Ember.Controller with Ember.SimpleAuth.LoginControllerMixin
  actions:
    sessionAuthenticationFailed: (error) ->
      msg = JSON.parse(error).error
      @set('errorMessage', msg)
    closeMessage: ->
      @set('errorMessage', null)
 
  reset: ->
    @setProperties
      errorMessage: null
      identification: null
      password: null
