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

step ":host is listening on :port" do |host, port|
  eventually(timeout: 120) {
    `nc -z #{host} #{port}`
    expect($?).to be_success
  }
end

step "I go to the :target page" do |target|
  path = case target
  when "notes" then "/notes"
  else
    raise "unknown page: '#{target}'"
  end
  visit(path)
end

step "I fill in the notes field with :note" do |note|
  @note = note
  fill_in 'note', with: @note
end

step "when I click submit, the note should appear in the list" do
  click_on 'submit'
  expect(page).to have_content @note
end
