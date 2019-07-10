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
			"project": {
				Type:        schema.TypeString,
				Description: "Project of the env",
				Required:    true,
				ForceNew:    true,
			},
			"environnement": {
				Type:        schema.TypeMap,
				Description: "environnement dict",
				Required:    true,
				ForceNew:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceSquarescaleEnvCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Create config")
	config := meta.(*Config)

	project := d.Get("project").(string)
	d.Partial(true)
	env, err := squarescale.NewEnvironment(config.Client, project)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Current env: '%q'", env.Project)

	if params, ok := d.GetOk("environnement"); ok {
		log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] params: '%q'", params)
		for key, _ := range params.(map[string]interface{}) {
			currentValue, err := env.Project.GetVariable(key)
			if err == nil && currentValue.Predefined {
				log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Can't redefine predefined keys (%s)", key)
				return fmt.Errorf("Error Can't redefine predefined key: %s", key)
			}
		}
		for key, value := range params.(map[string]interface{}) {
			currentValue, err := env.Project.GetVariable(key)
			log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Current key (%s) has value: %s", key, currentValue.Value)
			if err != nil {
				log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Add env variable '%s': '%s'", key, value)
				env.Project.SetVariable(key, value.(string))
			} else if value != currentValue.Value {
				log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] '%s' != '%s'", currentValue.Value, value)
				log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvCreate] Set env variable '%s' to '%s'", key, value)
				env.Project.SetVariable(key, value.(string))
			}
		}
		env.CommitEnvironment(config.Client, project)
		d.Set("environnement", params)
	}

	d.SetId(fmt.Sprintf("%s", project))

	d.Partial(false)

	return resourceSquarescaleEnvRead(d, meta)
}

func resourceSquarescaleEnvRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvRead] Create config")

	d.Partial(true)
	d.Partial(false)

	return nil
}

func resourceSquarescaleEnvDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvDelete] Create config")
	// config := meta.(*Config)

	// project := d.Get("project").(string)
	d.Partial(true)
	// env, err := squarescale.NewEnvironment(config.Client, project)
	// if err != nil {
	// 	return err
	// }
	// log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvDelete] Current env: '%q'", env.Project)

	// if params, ok := d.GetOk("environnement"); ok {
	// 	log.Printf("[DEBUG][SQSC][resourceSquarescaleEnvDelete] params: '%q'", params)
	// 	for key, _ := range params.(map[string]interface{}) {
	// 		env.Project.RemoveVariable(key)
	// 	}
	// 	env.CommitEnvironment(config.Client, project)
	// }

	d.SetId("")

	d.Partial(false)

	return nil
}
