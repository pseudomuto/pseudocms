class Admin.LoginController extends Ember.Controller with Ember.SimpleAuth.LoginControllerMixin
  reset: ->
    @setProperties
      errorMessage: null
      identification: null
      password: null
