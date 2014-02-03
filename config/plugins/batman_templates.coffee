module.exports = (lineman) ->
  config:
    batman_templates:
      admin:
        options:
          templateFolder: "app/assets/batman/admin/html"
        src: "<%= files.batman.views %>"
        dest: "<%= files.batman.generated %>"
