class Admin.LoginRoute extends Ember.Route
  activate: ->
    Ember.$('body').addClass('login')

  deactivate: ->
    Ember.$('body').removeClass('login')

  setupController: (controller, model) ->
    controller.reset()

  actions:
    sessionAuthenticationFailed: (response) ->
      @controller.set('errorMessage', response.message)
