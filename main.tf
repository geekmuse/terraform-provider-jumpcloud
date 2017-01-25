### 
#   This is a demo file, but it demonstrates the use
#   of the provider in its current state.
#   
#   Your JumpCloud API key can be found within the web console
#   of your JumpCloud account.  The properties for the 
#   "jumpcloud_user" resource should be fairly intuitive.


provider "jumpcloud" {
    api_key = "abcde12345"
}

resource "jumpcloud_user" "jcuser_test" {
    "user_name" = "jc-username"
    "first_name" = "Brad"
    "last_name" = "Campbell"
    "password" = "pAssw0rD"
    "email" = "me@example.net"
    "sudo" = true
}