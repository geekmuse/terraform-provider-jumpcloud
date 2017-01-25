
package main

import (
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/TheJumpCloud/jcapi"
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
                        Required:   true,
                    },
                    "sudo":  &schema.Schema{
                        Type:       schema.TypeBool,
                        Required:   true,
                    },
                },
                SchemaVersion:  1,
                Create:     CreateSystemUser,
                Read:       ReadSystemUser,
                Update:     UpdateSystemUser,
                Delete:     DeleteSystemUser,
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

func ReadSystemUser(d *schema.ResourceData, meta interface{}) error {

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
    jcUser.Activated =          true
    jcUser.ExternallyManaged =  false

    userId, err := meta.(*jcapi.JCAPI).AddUpdateUser(3, jcUser)

    if err != nil {
        return err
    }

    d.SetId(userId)
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

