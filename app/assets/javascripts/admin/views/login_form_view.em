class Admin.LoginFormView extends Ember.View
  tagName: 'form'
  email: ''
  password: ''

  submit: (event) ->
    event.preventDefault()
    @get('controller').login(@get('email'), @get('password'))
