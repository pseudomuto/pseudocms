@signout = ->
  Admin.__container__.registry.get('ember-simple-auth:session').clear(false)
