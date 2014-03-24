#= require jquery
#= require handlebars
#= require ember
#= require ember-data
#= require ember-simple-auth
#= require_self
#= require ./store
#= require_tree ./components
#= require_tree ./mixins
#= require_tree ./models
#= require_tree ./controllers
#= require_tree ./views
#= require_tree ./helpers
#= require_tree ./templates
#= require_tree ./routes
#= require ./router

Ember.Application.initializer
  name: 'authentication'
  initialize: (container, application) ->
    Ember.SimpleAuth.setup container, application, 
      authenticationRoute: 'login'
      routeAfterAuthentication: 'index'


window.Admin = Ember.Application.create
  LOG_TRANSITIONS: true
  #LOG_TRANSITIONS_INTERNAL: true
  #LOG_ACTIVE_GENERATION: true
  #LOG_VIEW_LOOKUPS: true
