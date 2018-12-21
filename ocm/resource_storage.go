package ocm

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/parnurzeal/gorequest"
	"log"
	"os"
)

func resourceStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceStorageCreate,
		Read:   resourceStorageRead,
		Update: resourceStorageUpdate,
		Delete: resourceStorageDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"size": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"properties": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"bootable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceStorageCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	size := d.Get("size").(string)
	path := d.Get("path").(string)
	prop := d.Get("properties").(string)
	boot := d.Get("bootable").(bool)
	image := d.Get("image").(string)

	//domainOCM := "10.8.120.38"
	domainOCM := m.(string)

	type storageJSON struct {
		Size       string   `json:"size"`
		Name       string   `json:"name"`
		Properties []string `json:"properties"`
		Bootable   bool     `json:"bootable"`
		Image      string   `json:"imagelist,omitempty"`
	}

	storageName := fmt.Sprintf("%s/%s", path, name)
	storeJSON := storageJSON{Size: size, Name: storageName, Properties: []string{prop}, Bootable: boot, Image: image}
	urlStorageVolume := fmt.Sprintf("http://%s/storage/volume%s/", domainOCM, path)

	resp, body, errs := gorequest.New().
		Post(urlStorageVolume).
		Set("Content-Type", "application/oracle-compute-v3+json").
		Set("Cookie", os.ExpandEnv("$COMPUTE_COOKIE")).
		Send(storeJSON).
		End()

	if resp.StatusCode == 409 {
		log.Println("[ERROR] Object already exists:", body, errs)
		return errors.New(`Storage already exists. Please see https://youtu.be/dQw4w9WgXcQ for more information on providing credentials for the OCM `)
	}

	d.SetId(name)
	return nil
}

func resourceStorageRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStorageUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceStorageDelete(d *schema.ResourceData, m interface{}) error {
	//domainOCM := "10.8.120.38"
	domainOCM := m.(string)
	name := d.Get("name").(string)
	path := d.Get("path").(string)

	b := destroyStorage(domainOCM, name, path)
	if b == false {
		//return errors.New(`Destruction failed. Please see https://youtu.be/SYnVYJDxu2Q for more information on destroy things in the OCM `)
	}
	//d.SetId("1")
	return nil
}

func destroyStorage(domain, name, path string) bool {
	urlStorageVolume := fmt.Sprintf("http://%s/storage/volume%s/%s", domain, path, name)
	resp, body, errs := gorequest.New().
		Delete(urlStorageVolume).
		Set("Content-Type", "application/oracle-compute-v3+json").
		Set("Cookie", os.ExpandEnv("$COMPUTE_COOKIE")).
		End()
	if resp.StatusCode != 204 {
		log.Println(os.ExpandEnv("$COMPUTE_COOKIE"))
		log.Println("[ERROR] Object not found:", body, errs)
		log.Println("[HD-ERROR]", resp.StatusCode, resp)
		return false
	} else {
		return true
	}
}
