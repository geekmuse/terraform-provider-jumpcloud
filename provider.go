
package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/geekmuse/jcapi"
)

const (
    apiUrl string = "https://console.jumpcloud.com/api"
)

func Provider() *schema.Provider {
    return &schema.Provider {
        ResourcesMap:   map[string]*schema.Resource{
            "jumpcloud_user": &schema.Resource{
                Schema: map[string]*schema.Schema{
                    "user_name": &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   true,
                        ForceNew:   true,
                    },
                    "first_name":  &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   false,
                        Optional:   true,
                    },
                    "last_name":  &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   false,
                        Optional:   true,
                    },
                    "email":  &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   true,
                    },
                    "password":  &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   false,
                        Optional:   true,
                    },
                    "sudo":  &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   false,
                        Optional:   true,
                    },
                    "passwordless_sudo":  &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   false,
                        Optional:   true,
                    },
                    "allow_public_key":  &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   false,
                        Optional:   true,
                    },
                    "public_key":  &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   false,
                        Optional:   true,
                    },
                },
                SchemaVersion:  1,
                Create:     CreateSystemUser,
                Read:       ReadSystemUser,
                Update:     UpdateSystemUser,
                Delete:     DeleteSystemUser,
                Importer:   &schema.ResourceImporter{
                    State: ImportSystemUser,
                },
            },
            "jumpcloud_system": &schema.Resource{
                Schema: map[string]*schema.Schema{
                    "display_name": &schema.Schema{
                        Type:       schema.TypeString,
                        Required:   true,
                        ForceNew:   false,
                    },
                    "allow_ssh_password_auth": &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   true,
                        ForceNew:   false,
                    },
                    "allow_ssh_root_login": &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   true,
                        ForceNew:   false,
                    },
                    "allow_multifactor_auth": &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   true,
                        ForceNew:   false,
                    },
                    "allow_public_key_auth": &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   true,
                        ForceNew:   false,
                    },
                    "tags": &schema.Schema{
                        Type:       schema.TypeList,
                        Elem:       &schema.Schema{Type: schema.TypeString},
                        Required:   false,
                        Optional:   true,
                    },
                },
                SchemaVersion:  1,
                Create:     CreateSystem,
                Read:       ReadSystem,
                Update:     UpdateSystem,
                Delete:     DeleteSystem,
                // Importer:   &schema.ResourceImporter{
                //     State: ImportSystem,
                // },
            },
        },
        Schema:         map[string]*schema.Schema{
            "api_key": &schema.Schema{
                Type:           schema.TypeString,
                Required:       true,
                Description:    "JumpCloud API key",
            },
        },
        ConfigureFunc:  providerInit,
    }
}

func providerInit(d *schema.ResourceData) (interface{}, error) {
    jcClient := jcapi.NewJCAPI(d.Get("api_key").(string), apiUrl)

    return &jcClient, nil
}

func CreateSystemUser(d *schema.ResourceData, meta interface{}) error {
    jcUser := jcapi.JCUser{
        UserName:           d.Get("user_name").(string),
        FirstName:          d.Get("first_name").(string),
        LastName:           d.Get("last_name").(string),
        Email:              d.Get("email").(string),
        Password:           d.Get("password").(string),
        Sudo:               d.Get("sudo").(bool),
        PasswordlessSudo:   d.Get("passwordless_sudo").(bool),
        AllowPublicKey:     d.Get("allow_public_key").(bool),
        PublicKey:          d.Get("public_key").(string),
        Activated:          true,
        ExternallyManaged:  false,
    }

    userId, err := meta.(*jcapi.JCAPI).AddUpdateUser(2, jcUser)

    if err != nil {
        return err
    }

    d.SetId(userId)
    return nil
}

// Adding systems in JumpCloud only allowed by Kickstart script.
// Once a system has been created in that way, it can be imported
// using Terraform's "import" command.
func CreateSystem(d *schema.ResourceData, meta interface{}) error {

    return nil
}


func ReadSystemUser(d *schema.ResourceData, meta interface{}) error {
    jcUser, err := meta.(*jcapi.JCAPI).GetSystemUserById(d.Id(), true)

    if err != nil {
        return err
    }

    d.Set("user_name", jcUser.UserName)
    d.Set("first_name", jcUser.FirstName)
    d.Set("last_name", jcUser.LastName)
    d.Set("email", jcUser.Email)
    // Not implemented in getJCUserFieldsFromInterface
    // d.Set("password", jcUser.Password)
    d.Set("sudo", jcUser.Sudo)
    d.Set("passwordless_sudo", jcUser.PasswordlessSudo)
    // Not implemented in getJCUserFieldsFromInterface
    // d.Set("allow_public_key", jcUser.AllowPublicKey)
    d.Set("public_key", jcUser.PublicKey)
    d.Set("uid", jcUser.Uid)
    d.Set("gid", jcUser.Gid)
    d.Set("enable_managed_uid", jcUser.EnableManagedUid)
    d.Set("activated", jcUser.Activated)
    d.Set("externally_managed", jcUser.ExternallyManaged)
    return nil
}

func ImportSystemUser(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
    if err := ReadSystemUser(d, meta); err != nil {
        return nil, err
    }

    return []*schema.ResourceData{d}, nil
}

func ReadSystem(d *schema.ResourceData, meta interface{}) error {

    return nil
}

func UpdateSystemUser(d *schema.ResourceData, meta interface{}) error {
    jcUser, err := meta.(*jcapi.JCAPI).GetSystemUserById(d.Id(), true)

    if err != nil {
        return err
    }

    jcUser.UserName =           d.Get("user_name").(string)
    jcUser.FirstName =          d.Get("first_name").(string)
    jcUser.LastName =           d.Get("last_name").(string)
    jcUser.Email =              d.Get("email").(string)
    jcUser.Password =           d.Get("password").(string)
    jcUser.Sudo  =              d.Get("sudo").(bool)
    jcUser.PasswordlessSudo =   d.Get("passwordless_sudo").(bool)
    jcUser.AllowPublicKey =     d.Get("allow_public_key").(bool)
    jcUser.PublicKey =          d.Get("public_key").(string)
    jcUser.Activated =          true
    jcUser.ExternallyManaged =  false

    userId, err := meta.(*jcapi.JCAPI).AddUpdateUser(3, jcUser)

    if err != nil {
        return err
    }

    d.SetId(userId)
    return nil
}

func UpdateSystem(d *schema.ResourceData, meta interface{}) error {

    return nil
}

func DeleteSystemUser(d *schema.ResourceData, meta interface{}) error {
    jcUser, err := meta.(*jcapi.JCAPI).GetSystemUserById(d.Id(), true)

    if err != nil {
        return err
    }

    err = meta.(*jcapi.JCAPI).DeleteUser(jcUser)

    if err != nil {
        return err
    }

    d.SetId("")
    return nil
}

func DeleteSystem(d *schema.ResourceData, meta interface{}) error {

    return nil
}

