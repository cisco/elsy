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
