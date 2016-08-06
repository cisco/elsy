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
