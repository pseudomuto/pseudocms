module 'login form view'

test 'submit event passes username and password to controller#login', ->
  controller =
    login: (email, password) ->

  mock = @mock(controller)
  mock.expects('login').withExactArgs('some@user.com', 'pAssword1')

  view = Admin.LoginFormView.create(
    email: 'some@user.com',
    password: 'pAssword1',
    controller: controller
  )

  view.submit($.Event('click'))
