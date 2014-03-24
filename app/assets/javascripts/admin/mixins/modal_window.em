mixin Admin.ModalWindow
  actions:
    close: ->
      @send('closeModal')
