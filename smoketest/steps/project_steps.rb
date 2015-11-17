require 'fileutils'

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
