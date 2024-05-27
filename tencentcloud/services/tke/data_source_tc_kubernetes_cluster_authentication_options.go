// Code generated by iacg; DO NOT EDIT.
package tke

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	tke "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tke/v20180525"

	tccommon "github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/common"
	"github.com/tencentcloudstack/terraform-provider-tencentcloud/tencentcloud/internal/helper"
)

func DataSourceTencentCloudKubernetesClusterAuthenticationOptions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTencentCloudKubernetesClusterAuthenticationOptionsRead,
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Cluster ID.",
			},

			"service_accounts": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "ServiceAccount authentication configuration. Note: this field may return `null`, indicating that no valid values can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_tke_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use TKE default issuer and jwksuri. Note: This field may return `null`, indicating that no valid values can be obtained.",
						},
						"issuer": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "service-account-issuer. Note: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"jwks_uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "service-account-jwks-uri. Note: this field may return `null`, indicating that no valid values can be obtained.",
						},
						"auto_create_discovery_anonymous_auth": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If it is set to `true`, a RABC rule is automatically created to allow anonymous users to access `/.well-known/openid-configuration` and `/openid/v1/jwks`. Note: this field may return `null`, indicating that no valid values can be obtained.",
						},
					},
				},
			},

			"latest_operation_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Result of the last modification. Values: `Updating`, `Success`, `Failed` or `TimeOut`. Note: this field may return `null`, indicating that no valid values can be obtained.",
			},

			"oidc_config": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "OIDC authentication configurations. Note: This field may return `null`, indicating that no valid value can be obtained.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_create_oidc_config": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Creating an identity provider. Note: This field may return `null`, indicating that no valid value can be obtained.",
						},
						"auto_create_client_id": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Creating ClientId of the identity provider. Note: This field may return `null`, indicating that no valid value can be obtained.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"auto_install_pod_identity_webhook_addon": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Creating the PodIdentityWebhook component. Note: This field may return `null`, indicating that no valid value can be obtained.",
						},
					},
				},
			},

			"result_output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Used to save results.",
			},
		},
	}
}

func dataSourceTencentCloudKubernetesClusterAuthenticationOptionsRead(d *schema.ResourceData, meta interface{}) error {
	defer tccommon.LogElapsed("data_source.tencentcloud_kubernetes_cluster_authentication_options.read")()
	defer tccommon.InconsistentCheck(d, meta)()

	logId := tccommon.GetLogId(nil)
	ctx := tccommon.NewResourceLifeCycleHandleFuncContext(context.Background(), logId, d, meta)

	service := TkeService{client: meta.(tccommon.ProviderMeta).GetAPIV3Conn()}

	var (
		clusterId string
	)
	if v, ok := d.GetOk("cluster_id"); ok {
		clusterId = v.(string)
	}
	paramMap := make(map[string]interface{})
	if v, ok := d.GetOk("cluster_id"); ok {
		paramMap["ClusterId"] = helper.String(v.(string))
	}

	var respData *tke.DescribeClusterAuthenticationOptionsResponseParams
	err := resource.Retry(tccommon.ReadRetryTimeout, func() *resource.RetryError {
		result, e := service.DescribeKubernetesClusterAuthenticationOptionsByFilter(ctx, paramMap)
		if e != nil {
			return tccommon.RetryError(e)
		}
		respData = result
		return nil
	})
	if err != nil {
		return err
	}

	serviceAccountsMap := map[string]interface{}{}

	if respData.ServiceAccounts != nil {
		if respData.ServiceAccounts.UseTKEDefault != nil {
			serviceAccountsMap["use_tke_default"] = respData.ServiceAccounts.UseTKEDefault
		}

		if respData.ServiceAccounts.Issuer != nil {
			serviceAccountsMap["issuer"] = respData.ServiceAccounts.Issuer
		}

		if respData.ServiceAccounts.JWKSURI != nil {
			serviceAccountsMap["jwks_uri"] = respData.ServiceAccounts.JWKSURI
		}

		if respData.ServiceAccounts.AutoCreateDiscoveryAnonymousAuth != nil {
			serviceAccountsMap["auto_create_discovery_anonymous_auth"] = respData.ServiceAccounts.AutoCreateDiscoveryAnonymousAuth
		}

		_ = d.Set("service_accounts", []interface{}{serviceAccountsMap})
	}

	if respData.LatestOperationState != nil {
		_ = d.Set("latest_operation_state", respData.LatestOperationState)
	}

	oIDCConfigMap := map[string]interface{}{}

	if respData.OIDCConfig != nil {
		if respData.OIDCConfig.AutoCreateOIDCConfig != nil {
			oIDCConfigMap["auto_create_oidc_config"] = respData.OIDCConfig.AutoCreateOIDCConfig
		}

		if respData.OIDCConfig.AutoCreateClientId != nil {
			oIDCConfigMap["auto_create_client_id"] = respData.OIDCConfig.AutoCreateClientId
		}

		if respData.OIDCConfig.AutoInstallPodIdentityWebhookAddon != nil {
			oIDCConfigMap["auto_install_pod_identity_webhook_addon"] = respData.OIDCConfig.AutoInstallPodIdentityWebhookAddon
		}

		_ = d.Set("oidc_config", []interface{}{oIDCConfigMap})
	}

	d.SetId(clusterId)

	output, ok := d.GetOk("result_output_file")
	if ok && output.(string) != "" {
		if e := tccommon.WriteToFile(output.(string), dataSourceTencentCloudKubernetesClusterAuthenticationOptionsReadOutputContent(ctx)); e != nil {
			return e
		}
	}

	return nil
}
