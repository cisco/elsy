step "it should pull :image" do |image|
  expect(@output).to match(%r{(Pulling from (library/)?#{image})|(Downloaded newer image for #{image})}i)
end

step "it :expectation fail pulling :image" do |expectation, image|
  meth = expectation ? :to : :to_not

  ## starting with docker 1.10 the error message changed to remove the 'latest' from the image name
  ## so this regex tests for both
  expect(@output).send(meth, match(%r{image library/#{image}(:latest|) not found}))
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
