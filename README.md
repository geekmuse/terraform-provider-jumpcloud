# jumpcloud-terraform-provider

A Terraform provider for [JumpCloud](https://jumpcloud.com).

As time goes on, this will implement JumpCloud's API as exposed through their [Go SDK](https://github.com/TheJumpCloud/jcapi).

## Where It Stands

Currently implemented is the creation, update, and deletion of users using the [SystemUsers interface](https://github.com/TheJumpCloud/JumpCloudAPI#system-users).

The following properties are currently implemented on a SystemUser:

*  UserName
*  FirstName
*  LastName
*  Email
*  Password
*  Sudo

## Usage

There is a demonstration Terraform implementation in `main.tf`.

## Roadmap

More properties for SystemUser will be quickly supported, and tests will be implemented thereafter.  A download of the compiled binary plugin reflective of the latest tag in the repo is currently available [here](https://bradcod.es/down/terraform-provider-jumpcloud).