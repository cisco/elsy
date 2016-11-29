# Copyright 2016 Cisco Systems, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
# 
# http://www.apache.org/licenses/LICENSE-2.0
# 
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

step "the :db database should exist" do |db|
  @db = db
  dbs = `mysql -h #{@mysql_host} -u #{@mysql_user} -p#{@mysql_pw} -s -N -e 'SHOW DATABASES'`
  expect($?).to eq(0), "failed running mysql command, response was:\n #{dbs}"
  expect(dbs.include? @db).to eq(true), "did not find db '#{@db}', found:\n #{dbs}"
end

step "it should contain the following empty tables:" do |table|
  tables = `mysql -h #{@mysql_host} -u #{@mysql_user} -p#{@mysql_pw} -s -N -e 'SHOW TABLES in #{@db}'`
  expect($?).to eq(0), "failed running mysql command, response was:\n #{tables}"
  table.raw.flatten.each do |expected|
    expect(tables.include? expected).to eq(true), "did not find table '#{@expected}', found:\n #{tables}"
    count = `mysql -h #{@mysql_host} -u #{@mysql_user} -p#{@mysql_pw} -s -N -e 'SELECT count(*) FROM #{@db}.#{expected}'`
    expect(count.to_i).to eq(0), "expected table #{expected} to be empty, but found #{count.to_i} row(s)"
  end
end

step "the :db database should contain :expected_count row in the :table table" do |db, expected_count, table|
  count = `mysql -h #{@mysql_host} -u #{@mysql_user} -p#{@mysql_pw} -s -N -e 'SELECT count(*) FROM #{db}.#{table}'`
  expect($?).to eq(0), "failed running mysql command, response was:\n #{count}"
  expect(count.to_i).to eq(expected_count.to_i), "expected table #{table} to contain #{expected_count} row(s), but found #{count.to_i} row(s)"
end
