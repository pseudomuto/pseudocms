#= require admin/admin
#= require sinon
#= require sinon-qunit
#= require jquery.mockjax
#= require_self
#= require_tree ./helpers
#= require_tree .

sinon.config =
  injectIntoThis: true
  injectInto: null
  properties: ["spy", "stub", "mock", "clock", "sandbox"]
  useFakeTimers: false
  useFakeServer: false

Admin.rootElement = '#qunit-fixture'
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
