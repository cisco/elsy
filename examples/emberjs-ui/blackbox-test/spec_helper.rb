require 'turnip/rspec'
require 'turnip/capybara'
require 'capybara/poltergeist'

RSpec.configure do |config|
    # fail when steps are unimplemented
  config.raise_error_for_unimplemented_steps = true
  config.color = true

  # load base helpers
  Dir.glob("blackbox-test/helpers/**/*_helper.rb") { |f| load f }
  config.include EventuallyHelper, :type => :feature

  Dir.glob("blackbox-test/steps/**/*_steps.rb") { |f| load f, true }

  Capybara.register_driver :poltergeist do |app|
  Capybara::Poltergeist::Driver.new(app, {
    phantomjs_logger: File.open('/dev/null', 'w'),
    window_size: [1280,1024],
    js_errors: false
  })
  end
  Capybara.default_driver = :poltergeist
  Capybara.javascript_driver = :poltergeist

  require 'gnawrnip'
  Gnawrnip.configure do |c|
    c.max_frame_size = 1280
  end
  Gnawrnip.ready!

  Capybara.app_host = "http://prodserver"
end
