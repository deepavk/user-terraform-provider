terraform {
  required_providers {
    httpserver = {
      version = "0.1.0"
      source = "local.com/customprovider/httpserver"
    }
  }
}

resource "httpserver_user" "user1" {
  name = "user1"
  phone = "111"
}


resource "httpserver_user" "user2" {
  name = "user2"
  phone = "222"
}


