module.exports = (lineman) ->
  lineman.config.application.watch = {}

  config:
    watch:
      admin:
        files: ["<%= files.coffee.app %>"]
        tasks: ["newer:coffee:admin", "concat_sourcemap:admin", "copy:admin"]
      batman:
        files: ["<%= files.batman.watch %>"]
        tasks: ["batman_templates:admin", "concat_sourcemap:admin", "copy:admin"]
      vendor:
        files: ["<%= files.js.vendor.watch %>"]
        tasks: ["concat_sourcemap:vendor", "copy:vendor"]
      spec:
        files: ["<%= files.coffee.spec %>"]
        tasks: ["newer:coffee:spec", "concat_sourcemap:spec"]
