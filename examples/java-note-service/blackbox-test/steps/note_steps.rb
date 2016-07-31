step "the note service is healthy" do
  eventually(timeout: 60){
    resp = Excon.get('http://noteservice:8080/admin/healthcheck')
    expect(resp.status).to be(200), "note service is not healthy, found body: #{resp.body}"
  }
end

step "I POST a message to :path with content :content" do |path, content|
  @resp = Excon.post("http://noteservice:8080#{path}",
    :body => content,
    :headers => { "Content-Type" => "application/json" })
end

step "I execute a GET against :path" do |path|
  @resp = Excon.get("http://noteservice:8080#{path}")
end

step "the response code should be :code" do |code|
  expect(@resp.status).to be(code.to_i), "expected response code of '#{code}', but found '#{@resp.status}'"
end

step "the response should contain :count note with the contents :content" do |count, content|
  json = MultiJson.load(@resp.body)
  expect(json.length).to eq(count.to_i)
  expect(json.first['note']).to eq(content)
end
