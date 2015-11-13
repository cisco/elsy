require 'turnip/rspec'

RSpec.configure do |config|
  Dir.glob("smoketest/steps/**/*_steps.rb") { |f| load f, true }
end

# add `lc` to our path
require 'fileutils'
FileUtils.ln_s("/opt/project/target/lc-linux-amd64", "/usr/local/bin/lc")
