package main

import (
	"github.com/aws-amplify/mcp-server/config"
	"github.com/aws-amplify/mcp-server/models"
	tools_apps "github.com/aws-amplify/mcp-server/tools/apps"
	tools_tags "github.com/aws-amplify/mcp-server/tools/tags"
	tools_artifacts "github.com/aws-amplify/mcp-server/tools/artifacts"
	tools_webhooks "github.com/aws-amplify/mcp-server/tools/webhooks"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_apps.CreateDeletebackendenvironmentTool(cfg),
		tools_apps.CreateGetbackendenvironmentTool(cfg),
		tools_apps.CreateStartdeploymentTool(cfg),
		tools_apps.CreateDeletejobTool(cfg),
		tools_apps.CreateGetjobTool(cfg),
		tools_apps.CreateCreatedeploymentTool(cfg),
		tools_apps.CreateListbackendenvironmentsTool(cfg),
		tools_apps.CreateCreatebackendenvironmentTool(cfg),
		tools_apps.CreateListartifactsTool(cfg),
		tools_apps.CreateListappsTool(cfg),
		tools_apps.CreateCreateappTool(cfg),
		tools_apps.CreateDeletebranchTool(cfg),
		tools_apps.CreateGetbranchTool(cfg),
		tools_apps.CreateUpdatebranchTool(cfg),
		tools_tags.CreateUntagresourceTool(cfg),
		tools_apps.CreateStopjobTool(cfg),
		tools_apps.CreateListdomainassociationsTool(cfg),
		tools_apps.CreateCreatedomainassociationTool(cfg),
		tools_apps.CreateListjobsTool(cfg),
		tools_apps.CreateStartjobTool(cfg),
		tools_apps.CreateUpdatedomainassociationTool(cfg),
		tools_apps.CreateDeletedomainassociationTool(cfg),
		tools_apps.CreateGetdomainassociationTool(cfg),
		tools_apps.CreateDeleteappTool(cfg),
		tools_apps.CreateGetappTool(cfg),
		tools_apps.CreateUpdateappTool(cfg),
		tools_apps.CreateGenerateaccesslogsTool(cfg),
		tools_apps.CreateListbranchesTool(cfg),
		tools_apps.CreateCreatebranchTool(cfg),
		tools_apps.CreateListwebhooksTool(cfg),
		tools_apps.CreateCreatewebhookTool(cfg),
		tools_artifacts.CreateGetartifacturlTool(cfg),
		tools_tags.CreateListtagsforresourceTool(cfg),
		tools_tags.CreateTagresourceTool(cfg),
		tools_webhooks.CreateDeletewebhookTool(cfg),
		tools_webhooks.CreateGetwebhookTool(cfg),
		tools_webhooks.CreateUpdatewebhookTool(cfg),
	}
}
