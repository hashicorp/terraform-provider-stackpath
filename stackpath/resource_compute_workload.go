package stackpath

import (
	"context"
	"net/http"

	workload "github.com/stackpath/terraform-provider-stackpath/stackpath/internal/client"
	"github.com/stackpath/terraform-provider-stackpath/stackpath/internal/models"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceComputeWorkload() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeWorkloadCreate,
		Read:   resourceComputeWorkloadRead,
		Update: resourceComputeWorkloadUpdate,
		Delete: resourceComputeWorkloadDelete,
		Importer: &schema.ResourceImporter{
			State: resourceComputeWorkloadImportState,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"slug": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"network_interface": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"image_pull_credentials": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"docker_registry": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"username": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"password": &schema.Schema{
										Type:      schema.TypeString,
										Required:  true,
										Sensitive: true,
									},
									"email": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
								},
							},
						},
					},
				},
			},
			"virtual_machine": &schema.Schema{
				Type:          schema.TypeList,
				ConflictsWith: []string{"container"},
				MaxItems:      1,
				Optional:      true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"port":            resourceComputeWorkloadPortSchema(),
						"liveness_probe":  resourceComputeWorkloadProbeSchema(),
						"readiness_probe": resourceComputeWorkloadProbeSchema(),
						"resources":       resourceComputeWorkloadResourcesSchema(),
						"volume_mount":    resourceComputeWorkloadVolumeMountSchema(),
					},
				},
			},
			"container": &schema.Schema{
				Type:          schema.TypeList,
				Optional:      true,
				ConflictsWith: []string{"virtual_machine"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"image": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"command": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"env": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"value": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"secret_value": &schema.Schema{
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
								},
							},
						},
						"port":            resourceComputeWorkloadPortSchema(),
						"readiness_probe": resourceComputeWorkloadProbeSchema(),
						"liveness_probe":  resourceComputeWorkloadProbeSchema(),
						"resources":       resourceComputeWorkloadResourcesSchema(),
						"volume_mount":    resourceComputeWorkloadVolumeMountSchema(),
					},
				},
			},
			"volume_claim": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"slug": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"resources": resourceComputeWorkloadResourcesSchema(),
					},
				},
			},
			"target": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MinItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"min_replicas": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
						"deployment_scope": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "cityCode",
						},
						"selector": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MinItems: 1,
							Elem:     resourceComputeMatchExpressionSchema(),
						},
					},
				},
			},
		},
	}
}

func resourceComputeWorkloadVolumeMountSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"slug": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"mount_path": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	}
}

func resourceComputeWorkloadProbeSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"http_get": &schema.Schema{
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"path": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
								Default:  "/",
							},
							"port": &schema.Schema{
								Type:     schema.TypeInt,
								Required: true,
							},
							"scheme": &schema.Schema{
								Type:     schema.TypeString,
								Optional: true,
								Default:  "http",
							},
							"http_headers": &schema.Schema{
								Type:     schema.TypeMap,
								Optional: true,
								Elem: &schema.Schema{
									Type: schema.TypeString,
								},
							},
						},
					},
				},
				"tcp_socket": &schema.Schema{
					Type:     schema.TypeList,
					MaxItems: 1,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"port": &schema.Schema{
								Type:     schema.TypeInt,
								Required: true,
							},
						},
					},
				},
				"initial_delay_seconds": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  0,
				},
				"timeout_seconds": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  10,
				},
				"period_seconds": &schema.Schema{
					Type:     schema.TypeInt,
					Optional: true,
					Default:  60,
				},
				"success_threshold": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"failure_threshold": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
			},
		},
	}
}

func resourceComputeWorkloadResourcesSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		MaxItems: 1,
		Required: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"requests": &schema.Schema{
					Type:     schema.TypeMap,
					Required: true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func resourceComputeWorkloadPortSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"port": &schema.Schema{
					Type:     schema.TypeInt,
					Required: true,
				},
				"protocol": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "tcp",
				},
			},
		},
	}
}

func resourceComputeWorkloadCreate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	// Create the workload
	resp, err := config.compute.CreateWorkload(&workload.CreateWorkloadParams{
		Context: context.Background(),
		StackID: config.Stack,
		Body: &models.V1CreateWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if err != nil {
		return err
	}

	// Set the ID based on the workload created in the API
	data.SetId(resp.Payload.Workload.ID)

	return nil
}

func resourceComputeWorkloadUpdate(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	_, err := config.compute.UpdateWorkload(&workload.UpdateWorkloadParams{
		Context:    context.Background(),
		StackID:    config.Stack,
		WorkloadID: data.Id(),
		Body: &models.V1UpdateWorkloadRequest{
			Workload: convertComputeWorkload(data),
		},
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return err
	}
	return err
}

func resourceComputeWorkloadRead(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	resp, err := config.compute.GetWorkload(&workload.GetWorkloadParams{
		Context:    context.Background(),
		StackID:    config.Stack,
		WorkloadID: data.Id(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return err
	}

	flattenComputeWorkload(data, resp.Payload.Workload)
	return nil
}

func resourceComputeWorkloadDelete(data *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := config.compute.DeleteWorkload(&workload.DeleteWorkloadParams{
		Context:    context.Background(),
		StackID:    config.Stack,
		WorkloadID: data.Id(),
	}, nil)
	if c, ok := err.(interface{ Code() int }); ok && c.Code() == http.StatusNotFound {
		// Clear out the ID in terraform if the
		// resource no longer exists in the API
		data.SetId("")
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func resourceComputeWorkloadImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	// We expect that to import a resource, the user will pass in
	// the full UUID of the workload they're attempting to import.
	return []*schema.ResourceData{d}, nil
}
