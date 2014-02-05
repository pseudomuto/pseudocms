APP_PATH = "app/assets/batman"
ADMIN_PATH = "#{APP_PATH}/admin"
VENDOR_PATH = "#{APP_PATH}/vendor"

module.exports =
  admin:
    src: ADMIN_PATH
    dest: "generated/js/admin"


  coffee:
    app: [
      "#{ADMIN_PATH}/pseudocms.coffee",
      "#{ADMIN_PATH}/lib/**/*.coffee",
      "#{ADMIN_PATH}/controllers/**/*.coffee",
      "#{ADMIN_PATH}/models/**/*.coffee",
      "#{ADMIN_PATH}/views/**/*.coffee"
    ]
    spec: "test/batman/**/*.coffee"
    generated: "generated/js/admin.coffee.js"
    generatedSpec: "generated/js/spec.coffee.js"

  js:
    admin:
      watch: "#{ADMIN_PATH}/**/*.coffee"
      src: [
        "generated/js/admin/pseudocms.js",
        # "generated/js/admin/lib/**/*.js",
        "generated/js/admin/controllers/**/*.js",
        # "generated/js/admin/models/**/*.js",
        "generated/js/admin/views/**/*.js"
      ]
      concatenated: "generated/js/admin-bundle.js"
      minified: "generated/js/admin-bundle.min.js"
      target: "public/assets/admin-bundle.js"
      sourcemap: 'generated/js/admin-bundle.js.map'
      sourcemapTarget: 'public/assets/admin-bundle.js.map'

    vendor:
      watch: "#{VENDOR_PATH}/**/*.js"
      src: [
        "#{VENDOR_PATH}/es5-shim.js",
        "#{VENDOR_PATH}/batman.js",
        "#{VENDOR_PATH}/batman.rails.js",
        "#{VENDOR_PATH}/jquery.js",
        "#{VENDOR_PATH}/batman.jquery.js"
      ],
      concatenated: "generated/js/vendor-bundle.js"
      minified: "generated/js/vendor-bundle.min.js"
      target: "public/assets/vendor-bundle.js"
      sourcemap: 'generated/js/vendor-bundle.js.map'
      sourcemapTarget: 'public/assets/vendor-bundle.js.map'

    spec:
      src_dir: "test/batman"
      dest_dir: "generated/js/spec"
      src: [
        "test/batman/vendor/sinon.js",
        "test/batman/vendor/sinon-qunit.js",
        "test/batman/vendor/batman.testing.js",
        "test/batman/vendor/batman.test_case.js",
        "test/batman/vendor/reqwest.js",
        "generated/js/spec/**/*.js"
      ]
      concatenated: "generated/js/spec-bundle.js"

  sass:
    main: ["#{ADMIN_PATH}/resources/css/style.scss"]
    generatedApp: "generated/css/admin-bundle.css"
    target: "public/assets/admin-bundle.css"

  batman:
    watch: ["#{ADMIN_PATH}/html"]
    views: ["#{ADMIN_PATH}/html/**/*.html"]
    generated: "generated/batman/view-store.js"
