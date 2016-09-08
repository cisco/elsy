require 'tmpdir'
require 'turnip/rspec'
require 'turnip_formatter'
require 'rspec_junit_formatter'

RSpec.configure do |config|
  # configure formatters
  config.add_formatter 'documentation'
  config.color = true
  config.add_formatter RSpecTurnipFormatter, "/opt/project/test-reports/servicetest-report.html"
  config.add_formatter RspecJunitFormatter, "/opt/project/test-reports/servicetest-rspec-junit.xml"

  # fail when steps are unimplemented
  config.raise_error_for_unimplemented_steps = true

  # load base helpers
  Dir.glob("blackbox-test/helpers/**/*_helper.rb") { |f| load f }
  config.include EventuallyHelper, :type => :feature
  Dir.glob("blackbox-test/steps/**/*_steps.rb") { |f| load f }

  config.before do
    @dir = Dir.mktmpdir("project-lifecycle-blackbox")
  end
  config.after do
    FileUtils.rm_rf(@dir)
  end
  config.after(:teardown => true) do
    Dir.chdir(@dir) do
      %x{lc teardown -f 2>&1}
    end
  end
end

# add `lc` to our path
require 'fileutils'
FileUtils.ln_s("/opt/project/target/lc-blackbox", "/usr/local/bin/lc")
