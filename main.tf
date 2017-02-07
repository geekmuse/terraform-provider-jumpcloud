### 
#   This is a demo file, but it demonstrates the use
#   of the provider in its current state.
#   
#   Your JumpCloud API key can be found within the web console
#   of your JumpCloud account.  The properties for the 
#   "jumpcloud_user" resource should be fairly intuitive.


provider "jumpcloud" {
    api_key = "your-api-key-here"
}

resource "jumpcloud_user" "jcuser_test" {
    "user_name" = "my_system_user"
    "first_name" = "Me"
    "last_name" = "SystemUser"
    "password" = "snArfblAt123!"
    "email" = "you@your-domain.net"
    "sudo" = true
    "passwordless_sudo" = true
    "allow_public_key" = true
    "public_key" = "${file("/abs/path/to/your/.ssh/public-key.pub")}"
}

resource "jumpcloud_user" "jcuser_test2" {
    "user_name" = "user_no_pub_key"
    "first_name" = "User"
    "last_name" = "NoPubKey"
    "password" = "f00f00!"
    "email" = "user@your-domain.net"
    "sudo" = false
    "passwordless_sudo" = false
}