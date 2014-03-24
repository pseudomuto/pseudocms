@integration = (name) ->
  module "Integration: #{name}",
    setup: ->
      Ember.run(Admin, Admin.advanceReadiness)

    teardown: ->
      Ember.run ->
        signout()

      Admin.reset()

@controller = (type, name) ->
  name = type unless name

  module "Controller: #{name}",
    controller: null,
    ajaxSpy: null,

    setup: ->
      Admin.__container__.registry.get('ember-simple-auth:session').clear()
      @controller = Admin.__container__.lookup("controller:#{type}")
      @ajaxSpy = sinon.spy(jQuery, 'ajax')

    teardown: ->
      jQuery.ajax.restore()
      Admin.reset()
