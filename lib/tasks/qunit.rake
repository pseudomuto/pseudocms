namespace :qunit do
  desc "run headless qunit tests"
  task test: :environment do
    unless %x(which phantomjs > /dev/null 2>&1)
      abort "PhantomJS is not installed!"
    end

    def port_available?(port)
      server = TCPServer.open(port)
      server.close
      true
    rescue Errno::EADDRINUSE
      false
    end

    port = ENV['TEST_SERVER_PORT'] || 60099
    while not port_available?(port)
      port += 1
    end

    unless pid = fork
      Rack::Server.start(
        config: 'config.ru',
        AccessLog: [],
        Port: port
      )

      exit
    end

    begin
      success = true
      js_dir = 'vendor/assets/javascripts'
      command = "phantomjs #{js_dir}/run-qunit.js http://localhost:#{port}/qunit"

      attempts = 0
      begin
        sh(command)
      rescue Exception => ex
        puts ex.message
        sleep 2
        attempts += 1
        retry unless attempts == 10
      end

      success &&= $?.success?
    ensure
      Process.kill "KILL", pid
    end

    if success
      puts "Tests passed."
    else
      puts "Tests failed."
      exit(1)
    end
  end
end
