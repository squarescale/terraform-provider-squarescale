package squarescale

import (
	"log"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://www.squarescale.io",
			},
			"github_login": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"github_password": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"squarescale_project": resourceSquarescaleProject(),
		},

		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"endpoint": "The Squarescale API endpoint (ex: \"https://www.squarescale.io\").",
		"github_login": "The github login",
		"github_password": "The github password",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	newconfig := Config{
		Endpoint:       d.Get("endpoint").(string),
		GithubLogin:    d.Get("github_login").(string),
		GithubPassword: d.Get("github_password").(string),
	}

	log.Printf("[DEBUG][SQSC] Provider, create new client")
	if err := newconfig.connectClient(); err != nil {
		return nil, err
	}

	return &newconfig, nil
}
