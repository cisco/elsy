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

step "it :expectation fail pulling :image" do |expectation, image|
  meth = expectation ? :to : :to_not

  ## starting with docker 1.10 the error message changed to remove the 'latest' from the image name
  ## so this regex tests for both
  expect(@output).send(meth, match(%r{image library/#{image}(:latest|) not found}))

  if meth == :to
    send "it should fail"
  else
    send "it should succeed"
  end
end

step "the image :image should exist" do |image_name|
  @image_name = image_name
  image_id = %x(docker images -q #{@image_name} 2>/dev/null)

  expect($?).to eq(0), "error running `docker images`: err:\n #{image_id}"
  expect(image_id).to_not be_nil
end

step "it should have the following labels:" do |table|
  labels = %x(docker inspect -f '{{ index .Config.Labels  }}' #{@image_name})
  expect($?).to eq(0), "error running `docker inspect`: err:\n #{labels}"

  table.raw.flatten.each do |label|
    expect(labels).to include(label)
  end
end

step "the following containers :expectation still exist:" do |should, table|
  container_names = table.raw.flatten.join(" ")

  containers = %x{docker-compose ps -q #{container_names} 2>/dev/null}

  count_to_look_for = if should then table.raw.flatten.length else 0 end

  expect(containers.split(/\n/).length).to be(count_to_look_for)
end

step "kill the following containers:" do |table|
  container_names = table.raw.flatten.join(" ")

  %x{docker-compose kill #{container_names}}
end
