module.exports = (lineman) ->
  # don't want these...
  delete lineman.config.application.concat_sourcemap.js
  delete lineman.config.application.concat_sourcemap.spec
  delete lineman.config.application.concat_sourcemap.css

  config:
    concat_sourcemap:
      admin:
        src: ["<%= files.js.admin.src %>", "<%= files.batman.generated %>"]
        dest: "<%= files.js.admin.concatenated %>"
      vendor:
        src: ["<%= files.js.vendor.src %>"]
        dest: "<%= files.js.vendor.concatenated %>"
      spec:
        src: ["<%= files.js.spec.src %>"]
        dest: "<%= files.js.spec.concatenated %>"
