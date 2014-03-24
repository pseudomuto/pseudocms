class Admin.ApplicationRoute extends Ember.Route with Ember.SimpleAuth.ApplicationRouteMixin
  actions:
    openModal: (name) ->
      @render name,
        into: 'application'
        outlet: 'modal'

    closeModal: ->
      @disconnectOutlet
        outlet: 'modal'
        parentView: 'application'
