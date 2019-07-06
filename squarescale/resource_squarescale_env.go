package squarescale

import (
	"fmt"
	"log"
	// "strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/squarescale/squarescale-cli/squarescale"
)

func resourceSquarescaleEnv() *schema.Resource {
	return &schema.Resource{
		Create: resourceSquarescaleEnvCreate,
		Read:   resourceSquarescaleEnvRead,
		Delete: resourceSquarescaleEnvDelete,

		Schema: map[string]*schema.Schema{
			"key": {
				Type:        schema.TypeString,
				Description: "key of the env",
				Required:    true,
				ForceNew:    true,
			},
			"project": {
				Type:        schema.TypeString,
				Description: "Project of the env",
				Required:    true,
				ForceNew:    true,
			},
			"value": {
				Type:        schema.TypeString,
				Description: "Value of the env",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func resourceSquarescaleEnvCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Create config")
	config := meta.(*Config)

	project := d.Get("project").(string)
	key := d.Get("key").(string)
	value := d.Get("value").(string)
	d.Partial(true)

	env, err := squarescale.NewEnvironment(config.Client, project)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Add env variable '%s': '%s'", key, value)
	env.Project.SetVariable(key, value)
	env.CommitEnvironment(config.Client, project)
	d.Set("key", key)
	d.Set("value", value)
	d.SetId(fmt.Sprintf("%s_%s", project, key))

	d.Partial(false)

	return resourceSquarescaleEnvRead(d, meta)
}

func resourceSquarescaleEnvRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvRead] Create config")
	config := meta.(*Config)
	project := d.Get("project").(string)
	key := d.Get("key").(string)

	d.Partial(true)

	env, err := squarescale.NewEnvironment(config.Client, project)
	if err != nil {
		return err
	}
	value, err := env.Project.GetVariable(key)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvRead] Value for key '%s': '%q'", key, value.Value)

	d.Set("value", value.Value)

	d.Partial(false)

	return nil
}

func resourceSquarescaleEnvDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvDelete] Create config")
	// config := meta.(*Config)
	// project := d.Get("project").(string)
	// key := d.Get("key").(string)

	d.Partial(true)
	// log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvDelete] Need to delete: %s", key)
	// env, err := squarescale.NewEnvironment(config.Client, project)
	// if err != nil {
	// 	return err
	// }
	// log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvDelete] Env: '%q' error: '%q'", env, err)

	// if err = env.Project.RemoveVariable(key); err != nil {
	// 	return err
	// }
	// env.CommitEnvironment(config.Client, project)

	d.SetId("")
	d.Partial(false)

	return nil
}
