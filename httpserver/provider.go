package httpserver

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"httpserver_user": getUserResourceSchema(), // Add this line
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}