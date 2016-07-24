step ":host is listening on port :port" do |host, port|
  eventually(timeout: 15) {
    `nc -z #{host} #{port}`
    expect($?).to be_success
  }
end
