step "it should pull :image" do |image|
  expect(@output).to match(%r{(Pulling from (library/)?#{image})|(Downloaded newer image for #{image})}i)
end

step "it :expectation fail pulling :image" do |expectation, image|
  meth = expectation ? :to : :to_not

  ## starting with docker 1.10 the error message changed to remove the 'latest' from the image name
  ## so this regex tests for both
  expect(@output).send(meth, match(%r{image library/#{image}(:latest|) not found}))
end
