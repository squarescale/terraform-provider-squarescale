package squarescale

import (
	"log"

	"github.com/squarescale/squarescale-cli/squarescale"
	"github.com/squarescale/squarescale-cli/tokenstore"
)

type Config struct {
	Endpoint       string
	GithubLogin    string
	GithubPassword string
	Client         *squarescale.Client
}

func (c *Config) connectClient() error {
	log.Printf("[DEBUG][SQSC] Use endpoint: %s", c.Endpoint)
	token, err := tokenstore.GetToken(c.Endpoint)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG][SQSC] Use token: %s", token)
	log.Printf("[DEBUG][SQSC] Token error: %s", err)

	c.Client = squarescale.NewClient(
		c.Endpoint,
		token,
	)
	//c.ApplicationKey,
	//c.ApplicationSecret,
	//c.ConsumerKey,
	if err != nil {
		return err
	}
	return nil
}
