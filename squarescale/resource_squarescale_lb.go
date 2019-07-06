package squarescale

import (
	"log"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSquarescaleLb() *schema.Resource {
	return &schema.Resource{
		Create: resourceSquarescaleLbCreate,
		Read:   resourceSquarescaleLbRead,
		Delete: resourceSquarescaleLbDelete,

		Schema: map[string]*schema.Schema{
			"container": {
				Type:        schema.TypeString,
				Description: "name of the container",
				Required:    true,
				ForceNew:    true,
			},
			"project": {
				Type:        schema.TypeString,
				Description: "Project of the lb",
				Required:    true,
				ForceNew:    true,
			},
			"port": {
				Type:        schema.TypeInt,
				Description: "Instances port",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceSquarescaleLbCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleLbCreate] Create config")
	config := meta.(*Config)

	project := d.Get("project").(string)
	containerName := d.Get("container").(string)
	port := d.Get("port").(int)

	d.Partial(true)

	container, err := config.Client.GetContainerInfo(project, containerName)
	if err != nil {
		return err
	}


	config.Client.ConfigLB(project, container.ID, port, false, "", nil, "")
	d.Set("container", container)
	d.Set("port", port)
	d.SetId(fmt.Sprintf("%s_%s", project, container))

	d.Partial(false)

	return resourceSquarescaleLbRead(d, meta)
}

func resourceSquarescaleLbRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleLbRead] Create config")

	return nil
}

func resourceSquarescaleLbDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleLbDelete] Create config")

	return nil
}
