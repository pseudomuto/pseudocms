grunt = require "grunt"

module.exports =
  pkg: grunt.file.readJSON("package.json")

  enableSass: true

  loadNpmTasks: [
    "grunt-batman-templates",
    "grunt-contrib-sass",
    "grunt-newer",
    "grunt-concat-sourcemap",
    "grunt-contrib-watch"
  ]

  appTasks:
    common: ["batman_templates", "newer:coffee", "sass", "concat_sourcemap"]
    dev: ["copy", "watch"]
    dist: ["copy"]

  newer:
    options:
      cache: "generated/timestamps"
