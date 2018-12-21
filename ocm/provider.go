package ocm

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_user", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_pass", nil),
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_VAR_domain", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ocm_storage": resourceStorage(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	config := Config{
		username: d.Get("username").(string),
		password: d.Get("password").(string),
		domain:   d.Get("domain").(string),
	}

	client, ok := config.Client()

	if ok == false {
		return nil, errors.New(`Client configuration failed. Please see https://youtu.be/b_ILDFp5DGA?t=1m22s for more help. `)
	}

	log.Println("[INFO] Initializing ocm client")

	return client, nil
}
