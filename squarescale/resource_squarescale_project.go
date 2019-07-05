package squarescale

import (
	"log"
	"time"
	"fmt"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/squarescale/squarescale-cli/squarescale"
)

func resourceSquarescaleProject() *schema.Resource {
	return &schema.Resource{
		Create: resourceSquarescaleProjectCreate,
		Read:   resourceSquarescaleProjectRead,
		Delete: resourceSquarescaleProjectDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "Name of the project",
				Required:    true,
				ForceNew:    true,
			},
		},
	}
}

func waitForSquarescaleProjectActive(c *squarescale.Client, definitiveName string, taskId int) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Printf("[DEBUG][SQSC][waitForSquarescaleProjectActive] Get project status: %s", definitiveName)
		log.Printf("[DEBUG][SQSC][waitForSquarescaleProjectActive] Get task id: %s", taskId)
		task, err := c.WaitTask(taskId)
		log.Printf("[DEBUG][SQSC][waitForSquarescaleProjectActive] TaskId & err status: %s // %s", task, err)
		projectStatus, err := c.ProjectStatus(definitiveName)
		log.Printf("[DEBUG][SQSC][waitForSquarescaleProjectActive] Project & err status: %s // %s", projectStatus, err)
		if err != nil {
			return nil, "", err
		}

		log.Printf("[DEBUG][SQSC][waitForSquarescaleProjectActive] Project done: %s", definitiveName)
		return projectStatus, "ok", nil
	}
}

func getProject(definitiveName string,c *squarescale.Client, project map[string]string) error {
	log.Printf("[DEBUG][SQSC][getProject] Create project map for %s", definitiveName)
	log.Printf("[DEBUG][SQSC][getProject] Current project map for %s", project)

	project["name"] = definitiveName
	log.Printf("[DEBUG][SQSC][getProject] Current project map for %s", project)

	return nil
}

func resourceSquarescaleProjectCreate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Create config")
	config := meta.(*Config)

	d.Partial(true)

	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Create cluster config")
	cluster := squarescale.ClusterConfig{
		InfraType: "single-node",
	}

	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Create Db config")
	db := squarescale.DbConfig{
		Size:   "dev",
		Engine: "postgres",
	}

	definitiveName := d.Get("name").(string)
	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Will create project: %s", definitiveName)
	taskId, err := config.Client.CreateProject(definitiveName, cluster, db)
	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Waiting for project create: %s", definitiveName)
	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Waiting for taskId: %s", taskId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"creating"},
		Target:     []string{"ok"},
		Refresh:    waitForSquarescaleProjectActive(config.Client, definitiveName, taskId),
		Timeout:    10 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 3 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("waiting for project: %s (%s)", definitiveName, err)
	}
	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectCreate] Created project %s", definitiveName)

	d.SetId(definitiveName)
	d.Set("name", definitiveName)

	d.Partial(false)

	return resourceSquarescaleProjectRead(d, meta)
}

func resourceSquarescaleProjectRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	d.Partial(true)

	definitiveName := d.Get("name").(string)
	projectStatus, err := config.Client.ProjectStatus(definitiveName)

	log.Printf("[DEBUG][SQSC][resourceSquarescaleProjectRead] Project & err status: %s // %s", projectStatus, err)
	if err != nil {
		return err
	}

	d.Set("name", definitiveName)

	d.Partial(false)

	return nil
}

func resourceSquarescaleProjectDelete(d *schema.ResourceData, _ interface{}) error {
	return nil
}
