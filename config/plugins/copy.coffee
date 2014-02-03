module.exports = (lineman) ->
  config:
    copy:
      admin:
        files:
          "<%= files.js.admin.target %>": "<%= files.js.admin.concatenated %>"
      vendor:
        files:
          "<%= files.js.vendor.target %>": "<%= files.js.vendor.concatenated %>"
      css:
        files:
          "<%= files.sass.target %>": "<%= files.sass.generatedApp %>"
