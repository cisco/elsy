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

step "the following folders should not be empty:" do |table|
  table.raw.flatten.each do |relativePath|
    absPath = File.join(@dir, relativePath)
    expect(Dir.exists?(absPath)).to be(true), "expected #{relativePath} to exist, but it was not found"
    expect(Dir["#{absPath}/*"].empty?).to be(false), "expected #{relativePath} to non empty, but it was empty"
  end
end

step "the following folders should be empty:" do |table|
  table.raw.flatten.each do |relativePath|
    absPath = File.join(@dir, relativePath)
    expect(Dir["#{absPath}/*"].empty?).to be(true), "expected #{relativePath} to be empty, but it had content"
  end
end
