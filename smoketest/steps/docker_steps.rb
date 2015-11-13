step "it should pull :image" do |image|
  expect(@output).to match(%r{(Pulling from (library/)?#{image})|(Downloaded newer image for #{image})}i)
end

step "it should fail pulling :image" do |image|
  expect(@output).to match(%r{image library/#{image}:latest not found})
end
