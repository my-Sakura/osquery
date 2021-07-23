#! /bin/expect

set timeout 30
spawn brew install --cask osquery
expect "password"
send "123456"