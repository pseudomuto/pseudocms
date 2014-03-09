#= require admin/admin
#= require sinon
#= require sinon-qunit
#= require jquery.mockjax
#= require_self
#= require_tree ./helpers
#= require_tree .

document.write('<div id="ember-testing-container"><div id="ember-testing"></div></div>')
document.write('<style>#ember-testing-container { position: absolute; background: white; bottom: 0; right: 0; width: 640px; height: 384px; overflow: auto; z-index: 9999; border: 1px solid #ccc; } #ember-testing { zoom: 50%; }</style>')

sinon.config =
  injectIntoThis: true
  injectInto: null
  properties: ["spy", "stub", "mock", "clock", "sandbox"]
  useFakeTimers: false
  useFakeServer: false

Admin.rootElement = '#ember-testing'
Admin.setupForTesting()
Admin.injectTestHelpers()

$.mockjaxSettings.logging = false

@exists = (selector) ->
  !!find(selector).length

@stubRequest = (url, options) ->
  defaults =
    url: url
    dataType: 'json'
    status: 200
    responseTime: 0

  $.mockjax $.extend({}, defaults, options)

QUnit.testDone (testDetails) ->
  $.mockjaxClear()
