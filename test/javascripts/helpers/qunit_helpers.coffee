@integration = (name) ->
  module "Integration: #{name}",
    setup: ->
      Admin.reset()
      Ember.run(Admin, Admin.advanceReadiness)

@controller = (type, name) ->
  name = type unless name

  module "Controller: #{name}",
    controller: null
    ajaxSpy: null
    setup: ->
      @controller = Admin.__container__.lookup("controller:#{type}")
      @ajaxSpy = sinon.spy(jQuery, 'ajax')

    teardown: ->
      jQuery.ajax.restore()
