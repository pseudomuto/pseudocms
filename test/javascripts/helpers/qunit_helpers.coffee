@integration = (name) ->
  module "Integration: #{name}",
    setup: ->
      Admin.reset()
      #Ember.run(Admin, Admin.advanceReadiness)

    teardown: ->
      #Admin.reset()
