Handlebars.registerHelper 'times', (n, block) ->
  block.fn(i) for i in [1..n]
  ''
