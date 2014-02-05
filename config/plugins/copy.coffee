module.exports = (lineman) ->
  config:
    copy:
      admin:
        files:
          "<%= files.js.admin.target %>": "<%= files.js.admin.concatenated %>"
          "<%= files.js.admin.sourcemapTarget %>": "<%= files.js.admin.sourcemap %>"
      vendor:
        files:
          "<%= files.js.vendor.target %>": "<%= files.js.vendor.concatenated %>"
          "<%= files.js.vendor.sourcemapTarget %>": "<%= files.js.vendor.sourcemap %>"
      css:
        files:
          "<%= files.sass.target %>": "<%= files.sass.generatedApp %>"
