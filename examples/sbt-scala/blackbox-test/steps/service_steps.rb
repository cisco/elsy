step ":host is listening on :port" do |host, port|
  eventually(timeout: 120) {
    `nc -z #{host} #{port}`
    expect($?).to be_success
  }
end

step "the homepage should contain :content" do |content|
  visit("/")
  eventually(timeout: 120) {
    expect(page).to have_content(content)
  }
end
