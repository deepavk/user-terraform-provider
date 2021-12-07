package httpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const baseURL = "http://localhost:8000/"

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func getURL(opType string) string {
	urlMap := map[string]string{
		"create": "createUser",
		"get":    "getUser",
		"update": "updateUser",
		"delete": "deleteUser",
	}

	return baseURL + urlMap[opType]
}

func createUserResource(ctx context.Context, resourceData *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics
	// Get fields from terraform resource data.

	name := resourceData.Get("name").(string)
	phone := resourceData.Get("phone").(string)

	// build user resource
	user := &User{
		Name:  name,
		Phone: phone,
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		return diag.FromErr(err)
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", getURL("create"), body)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get response from API
	var response User
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	// todo: add validation & error handling
	resourceData.Set("name", response.Name)
	resourceData.Set("phone", response.Phone)
	resourceData.SetId(response.Id)

	return diags
}

func updateUserResource(ctx context.Context, resourceData *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics
	// Get fields from terraform resource data.

	name := resourceData.Get("name").(string)
	phone := resourceData.Get("phone").(string)
	userid := resourceData.Id()

	// Build User object
	user := &User{
		Id:    userid,
		Name:  name,
		Phone: phone,
	}

	requestBody, err := json.Marshal(user)
	if err != nil {
		return diag.FromErr(err)
	}

	body := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", getURL("update"), body)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	// Get response from API
	var response User
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func deleteUserResource(ctx context.Context, resourceData *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	var diags diag.Diagnostics

	userid := resourceData.Id()
	url := getURL("delete") + "?id=" + userid
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	defer r.Body.Close()
	return diags
}

func readUserResource(ctx context.Context, resourceData *schema.ResourceData, m interface{}) diag.Diagnostics {
	userid := resourceData.Id()
	var diags diag.Diagnostics
	client := &http.Client{Timeout: 10 * time.Second}
	url := getURL("get") + "?id=" + userid
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}

	defer r.Body.Close()

	var user *User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func getUserResourceSchema() *schema.Resource {
	return &schema.Resource{
		CreateContext: createUserResource,
		UpdateContext: updateUserResource,
		ReadContext:   readUserResource,
		DeleteContext: deleteUserResource,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true, // Field is required
			},
			"phone": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
