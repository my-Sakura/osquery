#! /bin/expect

spawn brew install --cask osquery
expect "password"
send "123456"