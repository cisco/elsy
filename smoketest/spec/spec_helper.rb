require 'turnip/rspec'

RSpec.configure do |config|
  Dir.glob("smoketest/spec/steps/**/*_steps.rb") { |f| load f, true }
  config.pattern = "smoketest/spec/**/*.feature"
end

# add `lc` to our path
require 'fileutils'
FileUtils.ln_s("/opt/project/target/lc-linux-amd64", "/usr/local/bin/lc")
