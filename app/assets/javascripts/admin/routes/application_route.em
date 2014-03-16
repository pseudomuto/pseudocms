class Admin.ApplicationRoute extends Ember.Route with Ember.SimpleAuth.ApplicationRouteMixin
  actions:
    sessionInvalidationSucceeded: ->
      window.location.replace('/admin')
