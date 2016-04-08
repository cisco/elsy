placeholder :expectation do
  match /should not/ do
    false
  end

  match /should/ do
    true
  end
end
