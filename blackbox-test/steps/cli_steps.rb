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

step "the output should not contain :expected" do |expected|
  expect(@output).not_to include(expected)
end

step "the output should contain one of the following:" do |table|
  found = false

  table.raw.flatten.each do |expected|
    if @output.include? expected then
      found = true
      break
    end
  end

  expect(found).to be(true), "the output should contain one of '#{table.raw.flatten}', but it did not. found: \n#{@output}"
end

step "the output should contain all of these:" do |table|
  table.raw.flatten.each do |string|
    expect(@output).to include(string)
  end
end

step "the output should not contain any of these:" do |table|
  table.raw.flatten.each do |string|
    expect(@output).not_to include(string)
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

step "it should succeed with no output" do
  send "it should succeed"
  expect(@output.length == 0)
end
