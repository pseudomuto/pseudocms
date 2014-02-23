#= require jquery
#= require handlebars
#= require ember
#= require ember-data
#= require_self
#= require ./store
#= require_tree ./models
#= require_tree ./controllers
#= require_tree ./views
#= require_tree ./helpers
#= require_tree ./templates
#= require_tree ./routes
#= require ./router

window.Admin = Ember.Application.create()
  #LOG_TRANSITIONS: true
  #LOG_TRANSITIONS_INTERNAL: true
  #LOG_ACTIVE_GENERATION: true
  #LOG_VIEW_LOOKUPS: true
