module.exports = (lineman) ->
  config:
    coffee:
      admin:
        expand: true
        cwd: "<%= files.admin.src %>"
        src: "**/*.coffee"
        dest: "<%= files.admin.dest %>"
        ext: ".js"
      spec:
        expand: true
        cwd:  "<%= files.js.spec.src_dir %>"
        src:  "**/*.coffee"
        dest: "<%= files.js.spec.dest_dir %>"
        ext:  ".js"
