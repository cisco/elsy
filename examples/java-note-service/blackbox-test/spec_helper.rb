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

require 'turnip/rspec'
require 'excon'
require 'multi_json'

RSpec.configure do |config|
    # fail when steps are unimplemented
  config.raise_error_for_unimplemented_steps = true
  config.color = true

  # load base helpers
  Dir.glob("blackbox-test/helpers/**/*_helper.rb") { |f| load f }
  config.include EventuallyHelper, :type => :feature
  
  Dir.glob("blackbox-test/steps/**/*_steps.rb") { |f| load f, true }

  config.before do
    @mysql_host = 'mysql'
    @mysql_user = 'mysqluser'
    @mysql_pw = 'notsecurepw'
  end
end
