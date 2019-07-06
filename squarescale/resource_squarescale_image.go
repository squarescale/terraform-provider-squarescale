package squarescale

import (
	"log"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSquarescaleImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceSquarescaleImageCreate,
		Read:   resourceSquarescaleImageRead,
		Delete: resourceSquarescaleImageDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "name of the Docker image",
				Required:    true,
				ForceNew:    true,
			},
			"project": {
				Type:        schema.TypeString,
				Description: "Project of the image",
				Required:    true,
				ForceNew:    true,
			},
			"instances": {
				Type:        schema.TypeInt,
				Description: "Instances count",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceSquarescaleImageCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleImageCreate] Create config")
	config := meta.(*Config)

	project := d.Get("project").(string)
	name := d.Get("name").(string)
	instances := d.Get("instances").(int)
	d.Partial(true)

	config.Client.AddImage(project, name, instances)
	d.Set("name", name)
	d.Set("instances", instances)
	d.SetId(fmt.Sprintf("%s_%s", project, name))

	d.Partial(false)

	return resourceSquarescaleImageRead(d, meta)
}

func resourceSquarescaleImageRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleImageRead] Create config")

	return nil
}

func resourceSquarescaleImageDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleImageDelete] Create config")

	return nil
}
