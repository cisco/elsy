step "it should report a correct address" do
  ip = if RUBY_PLATFORM =~ /linux/
    "127.0.0.1"
  else
    "172.17.8.101"
  end

  expect(@output).to match(%r{.*#{ip}.*}im)
end

step "it should report :text" do |text|
  expect(@output).to match(%r{.*#{text}.*}im)
end
