require 'tmpdir'
require 'fileutils'

RSpec.configure do |config|
  config.before(:type => :feature) do
    @dir = Dir.mktmpdir("project-lifecycle-smoketest")
  end
  config.after(:type => :feature) do
    FileUtils.rm_rf(@dir)
  end
end

step "a file named :name with:" do |name, contents|
  parts = [@dir]
  dir = File.dirname(name)
  if !dir.empty?
    parts << dir
    FileUtils.mkdir_p(File.join(*parts))
  end
  parts << File.basename(name)
  IO.write(File.join(*parts), contents)
end
