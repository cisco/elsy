step "it should report :url" do |url|
  expect(@output).to match(%r{.*#{url}.*}im)
end
