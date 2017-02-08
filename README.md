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
*  Sudo/PasswordlessSudo
*  AllowPublicKey/PublicKey

## Usage

There is a demonstration Terraform implementation in `main.tf`.

## Roadmap

I'd like to dig into the `ExternallyManaged` feature of SystemUsers next, though I will probably resort to tests first, now that public keys are implemented; then onto systems (which it appears will leverage Terraform's `Import` functionality as a means of creation, then Update/Delete), and SystemUser/System associations.