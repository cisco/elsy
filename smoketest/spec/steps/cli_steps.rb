placeholder :command do
  match /`([^`]+)`/ do |command|
    command
  end
end

step "I run :command" do |command|
  Dir.chdir(@dir) do
    @output = %x{#{command} 2>&1}
    @rtn_status = $?
  end
end

step "it should succeed" do
  expect(@rtn_status).to be_success, "expected process to succeed. output was:\n#{@output}"
end

step "it should fail" do
  expect(@rtn_status).not_to be_success, "expected process to fail. output was:\n#{@output}"
end

step "the output should contain :expected" do |expected|
  expect(@output).to include(expected)
end

step "the output should contain all of these:" do |table|
  table.raw.flatten.each do |string|
    expect(@output).to include(string)
  end
end

step "it should succeed with :expected" do |expected|
  send "it should succeed"
  send "the output should contain :expected", expected
end

step "it should fail with :expected" do |expected|
  send "it should fail"
  send "the output should contain :expected", expected
end
