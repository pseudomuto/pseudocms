Batman.extend Batman.config,
  pathToApp: '/admin/'
  pathToHTML: '/admin/html'

class PseudoCMS extends Batman.App
  @root 'main#index'

(global ? window).PseudoCMS = PseudoCMS
