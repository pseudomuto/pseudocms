#= require admin/admin
#= require sinon
#= require sinon-qunit
#= require_self
#= require_tree .

document.write('<div id="ember-testing-container"><div id="ember-testing"></div></div>')
document.write('<style>#ember-testing-container { position: absolute; background: white; bottom: 0; right: 0; width: 640px; height: 384px; overflow: auto; z-index: 9999; border: 1px solid #ccc; } #ember-testing { zoom: 50%; }</style>')

Admin.rootElement = '#ember-testing'
Admin.setupForTesting()
Admin.injectTestHelpers()

@QUNIT_MODULE =
  setup: ->
    Ember.run(Admin, Admin.advanceReadiness)

  teardown: ->
    Admin.reset()

@exists = (selector) ->
  !!find(selector).length
