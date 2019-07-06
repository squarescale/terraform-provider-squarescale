package squarescale

import (
	"log"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSquarescaleDb() *schema.Resource {
	return &schema.Resource{
		Create: resourceSquarescaleDbCreate,
		Read:   resourceSquarescaleDbRead,
		Delete: resourceSquarescaleDbDelete,

		Schema: map[string]*schema.Schema{
			"engine": {
				Type:        schema.TypeString,
				Description: "Engine for DB",
				Required:    true,
				ForceNew:    true,
			},
			"project": {
				Type:        schema.TypeString,
				Description: "Project name",
				Required:    true,
				ForceNew:    true,
			},
			"size": {
				Type:        schema.TypeString,
				Description: "Size of DB",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceSquarescaleDbCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleDbCreate] Create config")
	config := meta.(*Config)

	project := d.Get("project").(string)
	engine := d.Get("engine").(string)
	size := d.Get("size").(string)
	d.Partial(true)

	dbConfig, err := config.Client.GetDBConfig(project)
	if err != nil {
		return err
	}

	dbConfig.Enabled = true
	dbConfig.Engine = engine
	dbConfig.Size = size

	config.Client.ConfigDB(project, dbConfig)
	d.Set("engine", engine)
	d.Set("size", size)
	d.SetId(fmt.Sprintf("%s_%s", project, engine))

	d.Partial(false)

	return resourceSquarescaleDbRead(d, meta)
}

func resourceSquarescaleDbRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleDbRead] Create config")

	return nil
}

func resourceSquarescaleDbDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleDbDelete] Create config")

	return nil
}
