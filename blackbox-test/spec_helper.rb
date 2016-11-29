# Copyright 2016 Cisco Systems, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
# http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

require 'tmpdir'
require 'turnip/rspec'
require 'turnip_formatter'
require 'rspec_junit_formatter'
require 'rspec/instafail'

RSpec.configure do |config|
  # configure formatters
  config.add_formatter 'documentation'
  config.color = true
  config.add_formatter RSpecTurnipFormatter, "/opt/project/test-reports/servicetest-report.html"
  config.add_formatter RspecJunitFormatter, "/opt/project/test-reports/servicetest-rspec-junit.xml"
  config.add_formatter RSpec::Instafail

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
