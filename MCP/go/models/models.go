package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListDomainAssociationsResult represents the ListDomainAssociationsResult schema from the OpenAPI specification
type ListDomainAssociationsResult struct {
	Domainassociations interface{} `json:"domainAssociations"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateAppRequestautoBranchCreationConfig represents the CreateAppRequestautoBranchCreationConfig schema from the OpenAPI specification
type CreateAppRequestautoBranchCreationConfig struct {
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview,omitempty"`
	Framework interface{} `json:"framework,omitempty"`
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
}

// DeleteDomainAssociationRequest represents the DeleteDomainAssociationRequest schema from the OpenAPI specification
type DeleteDomainAssociationRequest struct {
}

// UpdateWebhookRequest represents the UpdateWebhookRequest schema from the OpenAPI specification
type UpdateWebhookRequest struct {
	Branchname interface{} `json:"branchName,omitempty"`
	Description interface{} `json:"description,omitempty"`
}

// App represents the App schema from the OpenAPI specification
type App struct {
	Customrules interface{} `json:"customRules,omitempty"`
	Productionbranch AppproductionBranch `json:"productionBranch,omitempty"`
	Appid interface{} `json:"appId"`
	Tags interface{} `json:"tags,omitempty"`
	Customheaders interface{} `json:"customHeaders,omitempty"`
	Enablebranchautobuild interface{} `json:"enableBranchAutoBuild"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Defaultdomain interface{} `json:"defaultDomain"`
	Autobranchcreationconfig AppautoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty"`
	Autobranchcreationpatterns interface{} `json:"autoBranchCreationPatterns,omitempty"`
	Enableautobranchcreation interface{} `json:"enableAutoBranchCreation,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth"`
	Name interface{} `json:"name"`
	Repositoryclonemethod interface{} `json:"repositoryCloneMethod,omitempty"`
	Apparn interface{} `json:"appArn"`
	Createtime interface{} `json:"createTime"`
	Enablebranchautodeletion interface{} `json:"enableBranchAutoDeletion,omitempty"`
	Iamservicerolearn interface{} `json:"iamServiceRoleArn,omitempty"`
	Repository interface{} `json:"repository"`
	Updatetime interface{} `json:"updateTime"`
	Environmentvariables interface{} `json:"environmentVariables"`
	Platform interface{} `json:"platform"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Description interface{} `json:"description"`
}

// GetBackendEnvironmentResult represents the GetBackendEnvironmentResult schema from the OpenAPI specification
type GetBackendEnvironmentResult struct {
	Backendenvironment CreateBackendEnvironmentResultbackendEnvironment `json:"backendEnvironment"`
}

// DeleteBranchResultbranch represents the DeleteBranchResultbranch schema from the OpenAPI specification
type DeleteBranchResultbranch struct {
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Stage interface{} `json:"stage"`
	Framework interface{} `json:"framework"`
	Thumbnailurl interface{} `json:"thumbnailUrl,omitempty"`
	Brancharn interface{} `json:"branchArn"`
	Customdomains interface{} `json:"customDomains"`
	Ttl interface{} `json:"ttl"`
	Enablebasicauth interface{} `json:"enableBasicAuth"`
	Environmentvariables interface{} `json:"environmentVariables"`
	Displayname interface{} `json:"displayName"`
	Totalnumberofjobs interface{} `json:"totalNumberOfJobs"`
	Branchname interface{} `json:"branchName"`
	Description interface{} `json:"description"`
	Activejobid interface{} `json:"activeJobId"`
	Enablenotification interface{} `json:"enableNotification"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview"`
	Sourcebranch interface{} `json:"sourceBranch,omitempty"`
	Associatedresources interface{} `json:"associatedResources,omitempty"`
	Updatetime interface{} `json:"updateTime"`
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild"`
	Createtime interface{} `json:"createTime"`
	Destinationbranch interface{} `json:"destinationBranch,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
}

// CreateDomainAssociationrequest represents the CreateDomainAssociationrequest schema from the OpenAPI specification
type CreateDomainAssociationrequest struct {
	Enableautosubdomain bool `json:"enableAutoSubDomain,omitempty"` // Enables the automated creation of subdomains for branches.
	Subdomainsettings []SubDomainSetting `json:"subDomainSettings"` // The setting for the subdomain.
	Autosubdomaincreationpatterns []string `json:"autoSubDomainCreationPatterns,omitempty"` // Sets the branch patterns for automatic subdomain creation.
	Autosubdomainiamrole string `json:"autoSubDomainIAMRole,omitempty"` // The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.
	Domainname string `json:"domainName"` // The domain name for the domain association.
}

// StartJobResult represents the StartJobResult schema from the OpenAPI specification
type StartJobResult struct {
	Jobsummary StartDeploymentResultjobSummary `json:"jobSummary"`
}

// StopJobRequest represents the StopJobRequest schema from the OpenAPI specification
type StopJobRequest struct {
}

// ListJobsRequest represents the ListJobsRequest schema from the OpenAPI specification
type ListJobsRequest struct {
}

// DomainAssociation represents the DomainAssociation schema from the OpenAPI specification
type DomainAssociation struct {
	Enableautosubdomain interface{} `json:"enableAutoSubDomain"`
	Autosubdomainiamrole interface{} `json:"autoSubDomainIAMRole,omitempty"`
	Domainstatus interface{} `json:"domainStatus"`
	Subdomains interface{} `json:"subDomains"`
	Domainassociationarn interface{} `json:"domainAssociationArn"`
	Domainname interface{} `json:"domainName"`
	Statusreason interface{} `json:"statusReason"`
	Autosubdomaincreationpatterns interface{} `json:"autoSubDomainCreationPatterns,omitempty"`
	Certificateverificationdnsrecord interface{} `json:"certificateVerificationDNSRecord,omitempty"`
}

// CreateDomainAssociationResult represents the CreateDomainAssociationResult schema from the OpenAPI specification
type CreateDomainAssociationResult struct {
	Domainassociation CreateDomainAssociationResultdomainAssociation `json:"domainAssociation"`
}

// UpdateBranchResult represents the UpdateBranchResult schema from the OpenAPI specification
type UpdateBranchResult struct {
	Branch DeleteBranchResultbranch `json:"branch"`
}

// CreateBackendEnvironmentResult represents the CreateBackendEnvironmentResult schema from the OpenAPI specification
type CreateBackendEnvironmentResult struct {
	Backendenvironment CreateBackendEnvironmentResultbackendEnvironment `json:"backendEnvironment"`
}

// CreateBackendEnvironmentRequest represents the CreateBackendEnvironmentRequest schema from the OpenAPI specification
type CreateBackendEnvironmentRequest struct {
	Deploymentartifacts interface{} `json:"deploymentArtifacts,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Stackname interface{} `json:"stackName,omitempty"`
}

// UpdateDomainAssociationResultdomainAssociation represents the UpdateDomainAssociationResultdomainAssociation schema from the OpenAPI specification
type UpdateDomainAssociationResultdomainAssociation struct {
	Domainassociationarn interface{} `json:"domainAssociationArn"`
	Domainname interface{} `json:"domainName"`
	Statusreason interface{} `json:"statusReason"`
	Autosubdomaincreationpatterns interface{} `json:"autoSubDomainCreationPatterns,omitempty"`
	Certificateverificationdnsrecord interface{} `json:"certificateVerificationDNSRecord,omitempty"`
	Enableautosubdomain interface{} `json:"enableAutoSubDomain"`
	Autosubdomainiamrole interface{} `json:"autoSubDomainIAMRole,omitempty"`
	Domainstatus interface{} `json:"domainStatus"`
	Subdomains interface{} `json:"subDomains"`
}

// GetBranchResult represents the GetBranchResult schema from the OpenAPI specification
type GetBranchResult struct {
	Branch Branch `json:"branch"` // The branch for an Amplify app, which maps to a third-party repository branch.
}

// BackendEnvironment represents the BackendEnvironment schema from the OpenAPI specification
type BackendEnvironment struct {
	Stackname interface{} `json:"stackName,omitempty"`
	Updatetime interface{} `json:"updateTime"`
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn"`
	Createtime interface{} `json:"createTime"`
	Deploymentartifacts interface{} `json:"deploymentArtifacts,omitempty"`
	Environmentname interface{} `json:"environmentName"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// CreateApprequestautoBranchCreationConfig represents the CreateApprequestautoBranchCreationConfig schema from the OpenAPI specification
type CreateApprequestautoBranchCreationConfig struct {
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Framework interface{} `json:"framework,omitempty"`
}

// GenerateAccessLogsResult represents the GenerateAccessLogsResult schema from the OpenAPI specification
type GenerateAccessLogsResult struct {
	Logurl interface{} `json:"logUrl,omitempty"`
}

// ListBackendEnvironmentsResult represents the ListBackendEnvironmentsResult schema from the OpenAPI specification
type ListBackendEnvironmentsResult struct {
	Backendenvironments interface{} `json:"backendEnvironments"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// StartDeploymentRequest represents the StartDeploymentRequest schema from the OpenAPI specification
type StartDeploymentRequest struct {
	Jobid interface{} `json:"jobId,omitempty"`
	Sourceurl interface{} `json:"sourceUrl,omitempty"`
}

// ListAppsResult represents the ListAppsResult schema from the OpenAPI specification
type ListAppsResult struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Apps interface{} `json:"apps"`
}

// ListArtifactsResult represents the ListArtifactsResult schema from the OpenAPI specification
type ListArtifactsResult struct {
	Artifacts interface{} `json:"artifacts"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// GetWebhookResultwebhook represents the GetWebhookResultwebhook schema from the OpenAPI specification
type GetWebhookResultwebhook struct {
	Description interface{} `json:"description"`
	Updatetime interface{} `json:"updateTime"`
	Webhookarn interface{} `json:"webhookArn"`
	Webhookid interface{} `json:"webhookId"`
	Webhookurl interface{} `json:"webhookUrl"`
	Branchname interface{} `json:"branchName"`
	Createtime interface{} `json:"createTime"`
}

// UpdateWebhookResult represents the UpdateWebhookResult schema from the OpenAPI specification
type UpdateWebhookResult struct {
	Webhook CreateWebhookResultwebhook `json:"webhook"`
}

// TagMap represents the TagMap schema from the OpenAPI specification
type TagMap struct {
}

// CreateDomainAssociationResultdomainAssociation represents the CreateDomainAssociationResultdomainAssociation schema from the OpenAPI specification
type CreateDomainAssociationResultdomainAssociation struct {
	Statusreason interface{} `json:"statusReason"`
	Autosubdomaincreationpatterns interface{} `json:"autoSubDomainCreationPatterns,omitempty"`
	Certificateverificationdnsrecord interface{} `json:"certificateVerificationDNSRecord,omitempty"`
	Enableautosubdomain interface{} `json:"enableAutoSubDomain"`
	Autosubdomainiamrole interface{} `json:"autoSubDomainIAMRole,omitempty"`
	Domainstatus interface{} `json:"domainStatus"`
	Subdomains interface{} `json:"subDomains"`
	Domainassociationarn interface{} `json:"domainAssociationArn"`
	Domainname interface{} `json:"domainName"`
}

// StopJobResult represents the StopJobResult schema from the OpenAPI specification
type StopJobResult struct {
	Jobsummary StartDeploymentResultjobSummary `json:"jobSummary"`
}

// CreateWebhookResult represents the CreateWebhookResult schema from the OpenAPI specification
type CreateWebhookResult struct {
	Webhook CreateWebhookResultwebhook `json:"webhook"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Tags interface{} `json:"tags"`
}

// UpdateApprequest represents the UpdateApprequest schema from the OpenAPI specification
type UpdateApprequest struct {
	Repository string `json:"repository,omitempty"` // The name of the repository for an Amplify app
	Autobranchcreationconfig CreateApprequestautoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty"` // Describes the automated branch creation configuration.
	Enablebranchautodeletion bool `json:"enableBranchAutoDeletion,omitempty"` // Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	Name string `json:"name,omitempty"` // The name for an Amplify app.
	Platform string `json:"platform,omitempty"` // The platform for the Amplify app. For a static app, set the platform type to <code>WEB</code>. For a dynamic server-side rendered (SSR) app, set the platform type to <code>WEB_COMPUTE</code>. For an app requiring Amplify Hosting's original SSR support only, set the platform type to <code>WEB_DYNAMIC</code>.
	Environmentvariables map[string]interface{} `json:"environmentVariables,omitempty"` // The environment variables for an Amplify app.
	Accesstoken string `json:"accessToken,omitempty"` // <p>The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.</p> <p>Use <code>accessToken</code> for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use <code>oauthToken</code>.</p> <p>You must specify either <code>accessToken</code> or <code>oauthToken</code> when you update an app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href="https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>
	Buildspec string `json:"buildSpec,omitempty"` // The build specification (build spec) file for an Amplify app build.
	Oauthtoken string `json:"oauthToken,omitempty"` // <p>The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.</p> <p>Use <code>oauthToken</code> for repository providers other than GitHub, such as Bitbucket or CodeCommit.</p> <p>To authorize access to GitHub as your repository provider, use <code>accessToken</code>.</p> <p>You must specify either <code>oauthToken</code> or <code>accessToken</code> when you update an app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href="https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>
	Customrules []CustomRule `json:"customRules,omitempty"` // The custom redirect and rewrite rules for an Amplify app.
	Description string `json:"description,omitempty"` // The description for an Amplify app.
	Autobranchcreationpatterns []string `json:"autoBranchCreationPatterns,omitempty"` // Describes the automated branch creation glob patterns for an Amplify app.
	Enableautobranchcreation bool `json:"enableAutoBranchCreation,omitempty"` // Enables automated branch creation for an Amplify app.
	Customheaders string `json:"customHeaders,omitempty"` // The custom HTTP headers for an Amplify app.
	Iamservicerolearn string `json:"iamServiceRoleArn,omitempty"` // The AWS Identity and Access Management (IAM) service role for an Amplify app.
	Enablebranchautobuild bool `json:"enableBranchAutoBuild,omitempty"` // Enables branch auto-building for an Amplify app.
	Basicauthcredentials string `json:"basicAuthCredentials,omitempty"` // The basic authorization credentials for an Amplify app. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.
	Enablebasicauth bool `json:"enableBasicAuth,omitempty"` // Enables basic authorization for an Amplify app.
}

// CreateBackendEnvironmentResultbackendEnvironment represents the CreateBackendEnvironmentResultbackendEnvironment schema from the OpenAPI specification
type CreateBackendEnvironmentResultbackendEnvironment struct {
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn"`
	Createtime interface{} `json:"createTime"`
	Deploymentartifacts interface{} `json:"deploymentArtifacts,omitempty"`
	Environmentname interface{} `json:"environmentName"`
	Stackname interface{} `json:"stackName,omitempty"`
	Updatetime interface{} `json:"updateTime"`
}

// ProductionBranch represents the ProductionBranch schema from the OpenAPI specification
type ProductionBranch struct {
	Thumbnailurl interface{} `json:"thumbnailUrl,omitempty"`
	Branchname interface{} `json:"branchName,omitempty"`
	Lastdeploytime interface{} `json:"lastDeployTime,omitempty"`
	Status interface{} `json:"status,omitempty"`
}

// StartDeploymentrequest represents the StartDeploymentrequest schema from the OpenAPI specification
type StartDeploymentrequest struct {
	Jobid string `json:"jobId,omitempty"` // The job ID for this deployment, generated by the create deployment request.
	Sourceurl string `json:"sourceUrl,omitempty"` // The source URL for this deployment, used when calling start deployment without create deployment. The source URL can be any HTTP GET URL that is publicly accessible and downloads a single .zip file.
}

// TagResourcerequest represents the TagResourcerequest schema from the OpenAPI specification
type TagResourcerequest struct {
	Tags map[string]interface{} `json:"tags"` // The tags used to tag the resource.
}

// GetAppRequest represents the GetAppRequest schema from the OpenAPI specification
type GetAppRequest struct {
}

// UpdateAppResultapp represents the UpdateAppResultapp schema from the OpenAPI specification
type UpdateAppResultapp struct {
	Environmentvariables interface{} `json:"environmentVariables"`
	Platform interface{} `json:"platform"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Description interface{} `json:"description"`
	Customrules interface{} `json:"customRules,omitempty"`
	Productionbranch AppproductionBranch `json:"productionBranch,omitempty"`
	Appid interface{} `json:"appId"`
	Tags interface{} `json:"tags,omitempty"`
	Customheaders interface{} `json:"customHeaders,omitempty"`
	Enablebranchautobuild interface{} `json:"enableBranchAutoBuild"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Defaultdomain interface{} `json:"defaultDomain"`
	Autobranchcreationconfig AppautoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty"`
	Autobranchcreationpatterns interface{} `json:"autoBranchCreationPatterns,omitempty"`
	Enableautobranchcreation interface{} `json:"enableAutoBranchCreation,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth"`
	Name interface{} `json:"name"`
	Repositoryclonemethod interface{} `json:"repositoryCloneMethod,omitempty"`
	Apparn interface{} `json:"appArn"`
	Createtime interface{} `json:"createTime"`
	Enablebranchautodeletion interface{} `json:"enableBranchAutoDeletion,omitempty"`
	Iamservicerolearn interface{} `json:"iamServiceRoleArn,omitempty"`
	Repository interface{} `json:"repository"`
	Updatetime interface{} `json:"updateTime"`
}

// DeleteJobRequest represents the DeleteJobRequest schema from the OpenAPI specification
type DeleteJobRequest struct {
}

// DeleteBackendEnvironmentResult represents the DeleteBackendEnvironmentResult schema from the OpenAPI specification
type DeleteBackendEnvironmentResult struct {
	Backendenvironment CreateBackendEnvironmentResultbackendEnvironment `json:"backendEnvironment"`
}

// CreateDeploymentRequest represents the CreateDeploymentRequest schema from the OpenAPI specification
type CreateDeploymentRequest struct {
	Filemap interface{} `json:"fileMap,omitempty"`
}

// AppautoBranchCreationConfig represents the AppautoBranchCreationConfig schema from the OpenAPI specification
type AppautoBranchCreationConfig struct {
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview,omitempty"`
	Framework interface{} `json:"framework,omitempty"`
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
}

// UpdateBranchRequest represents the UpdateBranchRequest schema from the OpenAPI specification
type UpdateBranchRequest struct {
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Framework interface{} `json:"framework,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Ttl interface{} `json:"ttl,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild,omitempty"`
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Enablenotification interface{} `json:"enableNotification,omitempty"`
}

// DeleteBranchResult represents the DeleteBranchResult schema from the OpenAPI specification
type DeleteBranchResult struct {
	Branch DeleteBranchResultbranch `json:"branch"`
}

// UpdateWebhookrequest represents the UpdateWebhookrequest schema from the OpenAPI specification
type UpdateWebhookrequest struct {
	Branchname string `json:"branchName,omitempty"` // The name for a branch that is part of an Amplify app.
	Description string `json:"description,omitempty"` // The description for a webhook.
}

// CreateBranchrequest represents the CreateBranchrequest schema from the OpenAPI specification
type CreateBranchrequest struct {
	Branchname string `json:"branchName"` // The name for the branch.
	Pullrequestenvironmentname string `json:"pullRequestEnvironmentName,omitempty"` // The Amplify environment name for the pull request.
	Tags map[string]interface{} `json:"tags,omitempty"` // The tag for the branch.
	Enableautobuild bool `json:"enableAutoBuild,omitempty"` // Enables auto building for the branch.
	Enablenotification bool `json:"enableNotification,omitempty"` // Enables notifications for the branch.
	Buildspec string `json:"buildSpec,omitempty"` // The build specification (build spec) file for an Amplify app build.
	Basicauthcredentials string `json:"basicAuthCredentials,omitempty"` // The basic authorization credentials for the branch. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.
	Enablebasicauth bool `json:"enableBasicAuth,omitempty"` // Enables basic authorization for the branch.
	Ttl string `json:"ttl,omitempty"` // The content Time to Live (TTL) for the website in seconds.
	Environmentvariables map[string]interface{} `json:"environmentVariables,omitempty"` // The environment variables for the branch.
	Displayname string `json:"displayName,omitempty"` // The display name for a branch. This is used as the default domain prefix.
	Backendenvironmentarn string `json:"backendEnvironmentArn,omitempty"` // The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	Framework string `json:"framework,omitempty"` // The framework for the branch.
	Enablepullrequestpreview bool `json:"enablePullRequestPreview,omitempty"` // Enables pull request previews for this branch.
	Description string `json:"description,omitempty"` // The description for the branch.
	Enableperformancemode bool `json:"enablePerformanceMode,omitempty"` // <p>Enables performance mode for the branch.</p> <p>Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out. </p>
	Stage string `json:"stage,omitempty"` // Describes the current stage for the branch.
}

// CreateWebhookRequest represents the CreateWebhookRequest schema from the OpenAPI specification
type CreateWebhookRequest struct {
	Description interface{} `json:"description,omitempty"`
	Branchname interface{} `json:"branchName"`
}

// ListAppsRequest represents the ListAppsRequest schema from the OpenAPI specification
type ListAppsRequest struct {
}

// UpdateDomainAssociationrequest represents the UpdateDomainAssociationrequest schema from the OpenAPI specification
type UpdateDomainAssociationrequest struct {
	Autosubdomaincreationpatterns []string `json:"autoSubDomainCreationPatterns,omitempty"` // Sets the branch patterns for automatic subdomain creation.
	Autosubdomainiamrole string `json:"autoSubDomainIAMRole,omitempty"` // The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.
	Enableautosubdomain bool `json:"enableAutoSubDomain,omitempty"` // Enables the automated creation of subdomains for branches.
	Subdomainsettings []SubDomainSetting `json:"subDomainSettings,omitempty"` // Describes the settings for the subdomain.
}

// CreateDeploymentrequest represents the CreateDeploymentrequest schema from the OpenAPI specification
type CreateDeploymentrequest struct {
	Filemap map[string]interface{} `json:"fileMap,omitempty"` // An optional file map that contains the file name as the key and the file content md5 hash as the value. If this argument is provided, the service will generate a unique upload URL per file. Otherwise, the service will only generate a single upload URL for the zipped files.
}

// DeleteBackendEnvironmentRequest represents the DeleteBackendEnvironmentRequest schema from the OpenAPI specification
type DeleteBackendEnvironmentRequest struct {
}

// Screenshots represents the Screenshots schema from the OpenAPI specification
type Screenshots struct {
}

// GetArtifactUrlRequest represents the GetArtifactUrlRequest schema from the OpenAPI specification
type GetArtifactUrlRequest struct {
}

// Artifact represents the Artifact schema from the OpenAPI specification
type Artifact struct {
	Artifactfilename interface{} `json:"artifactFileName"`
	Artifactid interface{} `json:"artifactId"`
}

// SubDomainsubDomainSetting represents the SubDomainsubDomainSetting schema from the OpenAPI specification
type SubDomainsubDomainSetting struct {
	Branchname interface{} `json:"branchName"`
	Prefix interface{} `json:"prefix"`
}

// AutoBranchCreationConfig represents the AutoBranchCreationConfig schema from the OpenAPI specification
type AutoBranchCreationConfig struct {
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview,omitempty"`
	Framework interface{} `json:"framework,omitempty"`
}

// GetBackendEnvironmentRequest represents the GetBackendEnvironmentRequest schema from the OpenAPI specification
type GetBackendEnvironmentRequest struct {
}

// CreateAppResult represents the CreateAppResult schema from the OpenAPI specification
type CreateAppResult struct {
	App App `json:"app"` // Represents the different branches of a repository for building, deploying, and hosting an Amplify app.
}

// DeleteAppRequest represents the DeleteAppRequest schema from the OpenAPI specification
type DeleteAppRequest struct {
}

// ListDomainAssociationsRequest represents the ListDomainAssociationsRequest schema from the OpenAPI specification
type ListDomainAssociationsRequest struct {
}

// GetWebhookRequest represents the GetWebhookRequest schema from the OpenAPI specification
type GetWebhookRequest struct {
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
}

// Job represents the Job schema from the OpenAPI specification
type Job struct {
	Steps interface{} `json:"steps"`
	Summary Jobsummary `json:"summary"`
}

// CreateDomainAssociationRequest represents the CreateDomainAssociationRequest schema from the OpenAPI specification
type CreateDomainAssociationRequest struct {
	Autosubdomaincreationpatterns interface{} `json:"autoSubDomainCreationPatterns,omitempty"`
	Autosubdomainiamrole interface{} `json:"autoSubDomainIAMRole,omitempty"`
	Domainname interface{} `json:"domainName"`
	Enableautosubdomain interface{} `json:"enableAutoSubDomain,omitempty"`
	Subdomainsettings interface{} `json:"subDomainSettings"`
}

// UpdateBranchrequest represents the UpdateBranchrequest schema from the OpenAPI specification
type UpdateBranchrequest struct {
	Basicauthcredentials string `json:"basicAuthCredentials,omitempty"` // The basic authorization credentials for the branch. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.
	Description string `json:"description,omitempty"` // The description for the branch.
	Framework string `json:"framework,omitempty"` // The framework for the branch.
	Pullrequestenvironmentname string `json:"pullRequestEnvironmentName,omitempty"` // The Amplify environment name for the pull request.
	Enablepullrequestpreview bool `json:"enablePullRequestPreview,omitempty"` // Enables pull request previews for this branch.
	Backendenvironmentarn string `json:"backendEnvironmentArn,omitempty"` // The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.
	Environmentvariables map[string]interface{} `json:"environmentVariables,omitempty"` // The environment variables for the branch.
	Enableautobuild bool `json:"enableAutoBuild,omitempty"` // Enables auto building for the branch.
	Stage string `json:"stage,omitempty"` // Describes the current stage for the branch.
	Displayname string `json:"displayName,omitempty"` // The display name for a branch. This is used as the default domain prefix.
	Buildspec string `json:"buildSpec,omitempty"` // The build specification (build spec) file for an Amplify app build.
	Ttl string `json:"ttl,omitempty"` // The content Time to Live (TTL) for the website in seconds.
	Enablenotification bool `json:"enableNotification,omitempty"` // Enables notifications for the branch.
	Enableperformancemode bool `json:"enablePerformanceMode,omitempty"` // <p>Enables performance mode for the branch.</p> <p>Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out. </p>
	Enablebasicauth bool `json:"enableBasicAuth,omitempty"` // Enables basic authorization for the branch.
}

// GetDomainAssociationRequest represents the GetDomainAssociationRequest schema from the OpenAPI specification
type GetDomainAssociationRequest struct {
}

// DeleteWebhookRequest represents the DeleteWebhookRequest schema from the OpenAPI specification
type DeleteWebhookRequest struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// ListArtifactsRequest represents the ListArtifactsRequest schema from the OpenAPI specification
type ListArtifactsRequest struct {
}

// Webhook represents the Webhook schema from the OpenAPI specification
type Webhook struct {
	Description interface{} `json:"description"`
	Updatetime interface{} `json:"updateTime"`
	Webhookarn interface{} `json:"webhookArn"`
	Webhookid interface{} `json:"webhookId"`
	Webhookurl interface{} `json:"webhookUrl"`
	Branchname interface{} `json:"branchName"`
	Createtime interface{} `json:"createTime"`
}

// CreateBranchRequest represents the CreateBranchRequest schema from the OpenAPI specification
type CreateBranchRequest struct {
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Branchname interface{} `json:"branchName"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview,omitempty"`
	Enablenotification interface{} `json:"enableNotification,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Ttl interface{} `json:"ttl,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild,omitempty"`
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Framework interface{} `json:"framework,omitempty"`
	Displayname interface{} `json:"displayName,omitempty"`
	Stage interface{} `json:"stage,omitempty"`
}

// CreateWebhookrequest represents the CreateWebhookrequest schema from the OpenAPI specification
type CreateWebhookrequest struct {
	Description string `json:"description,omitempty"` // The description for a webhook.
	Branchname string `json:"branchName"` // The name for a branch that is part of an Amplify app.
}

// DeleteAppResult represents the DeleteAppResult schema from the OpenAPI specification
type DeleteAppResult struct {
	App App `json:"app"` // Represents the different branches of a repository for building, deploying, and hosting an Amplify app.
}

// ListBranchesRequest represents the ListBranchesRequest schema from the OpenAPI specification
type ListBranchesRequest struct {
}

// FileUploadUrls represents the FileUploadUrls schema from the OpenAPI specification
type FileUploadUrls struct {
}

// CreateBackendEnvironmentrequest represents the CreateBackendEnvironmentrequest schema from the OpenAPI specification
type CreateBackendEnvironmentrequest struct {
	Deploymentartifacts string `json:"deploymentArtifacts,omitempty"` // The name of deployment artifacts.
	Environmentname string `json:"environmentName"` // The name for the backend environment.
	Stackname string `json:"stackName,omitempty"` // The AWS CloudFormation stack name of a backend environment.
}

// UpdateAppRequest represents the UpdateAppRequest schema from the OpenAPI specification
type UpdateAppRequest struct {
	Iamservicerolearn interface{} `json:"iamServiceRoleArn,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Autobranchcreationconfig CreateAppRequestautoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty"`
	Enablebranchautodeletion interface{} `json:"enableBranchAutoDeletion,omitempty"`
	Oauthtoken interface{} `json:"oauthToken,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Customheaders interface{} `json:"customHeaders,omitempty"`
	Enablebranchautobuild interface{} `json:"enableBranchAutoBuild,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Repository interface{} `json:"repository,omitempty"`
	Enableautobranchcreation interface{} `json:"enableAutoBranchCreation,omitempty"`
	Platform interface{} `json:"platform,omitempty"`
	Autobranchcreationpatterns interface{} `json:"autoBranchCreationPatterns,omitempty"`
	Customrules interface{} `json:"customRules,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
}

// CustomRule represents the CustomRule schema from the OpenAPI specification
type CustomRule struct {
	Status interface{} `json:"status,omitempty"`
	Target interface{} `json:"target"`
	Condition interface{} `json:"condition,omitempty"`
	Source interface{} `json:"source"`
}

// StartJobRequest represents the StartJobRequest schema from the OpenAPI specification
type StartJobRequest struct {
	Jobid interface{} `json:"jobId,omitempty"`
	Jobreason interface{} `json:"jobReason,omitempty"`
	Jobtype interface{} `json:"jobType"`
	Commitid interface{} `json:"commitId,omitempty"`
	Commitmessage interface{} `json:"commitMessage,omitempty"`
	Committime interface{} `json:"commitTime,omitempty"`
}

// GenerateAccessLogsRequest represents the GenerateAccessLogsRequest schema from the OpenAPI specification
type GenerateAccessLogsRequest struct {
	Domainname interface{} `json:"domainName"`
	Endtime interface{} `json:"endTime,omitempty"`
	Starttime interface{} `json:"startTime,omitempty"`
}

// GetWebhookResult represents the GetWebhookResult schema from the OpenAPI specification
type GetWebhookResult struct {
	Webhook GetWebhookResultwebhook `json:"webhook"`
}

// ListBranchesResult represents the ListBranchesResult schema from the OpenAPI specification
type ListBranchesResult struct {
	Branches interface{} `json:"branches"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// ListJobsResult represents the ListJobsResult schema from the OpenAPI specification
type ListJobsResult struct {
	Jobsummaries interface{} `json:"jobSummaries"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// CreateDeploymentResult represents the CreateDeploymentResult schema from the OpenAPI specification
type CreateDeploymentResult struct {
	Jobid interface{} `json:"jobId,omitempty"`
	Zipuploadurl interface{} `json:"zipUploadUrl"`
	Fileuploadurls interface{} `json:"fileUploadUrls"`
}

// GetBranchRequest represents the GetBranchRequest schema from the OpenAPI specification
type GetBranchRequest struct {
}

// DeleteDomainAssociationResult represents the DeleteDomainAssociationResult schema from the OpenAPI specification
type DeleteDomainAssociationResult struct {
	Domainassociation DomainAssociation `json:"domainAssociation"` // Describes a domain association that associates a custom domain with an Amplify app.
}

// GetJobRequest represents the GetJobRequest schema from the OpenAPI specification
type GetJobRequest struct {
}

// SubDomain represents the SubDomain schema from the OpenAPI specification
type SubDomain struct {
	Dnsrecord interface{} `json:"dnsRecord"`
	Subdomainsetting SubDomainsubDomainSetting `json:"subDomainSetting"`
	Verified interface{} `json:"verified"`
}

// EnvironmentVariables represents the EnvironmentVariables schema from the OpenAPI specification
type EnvironmentVariables struct {
}

// UpdateDomainAssociationResult represents the UpdateDomainAssociationResult schema from the OpenAPI specification
type UpdateDomainAssociationResult struct {
	Domainassociation UpdateDomainAssociationResultdomainAssociation `json:"domainAssociation"`
}

// DeleteWebhookResult represents the DeleteWebhookResult schema from the OpenAPI specification
type DeleteWebhookResult struct {
	Webhook CreateWebhookResultwebhook `json:"webhook"`
}

// StartDeploymentResultjobSummary represents the StartDeploymentResultjobSummary schema from the OpenAPI specification
type StartDeploymentResultjobSummary struct {
	Jobtype interface{} `json:"jobType"`
	Commitid interface{} `json:"commitId"`
	Status interface{} `json:"status"`
	Committime interface{} `json:"commitTime"`
	Endtime interface{} `json:"endTime,omitempty"`
	Jobid interface{} `json:"jobId"`
	Starttime interface{} `json:"startTime"`
	Commitmessage interface{} `json:"commitMessage"`
	Jobarn interface{} `json:"jobArn"`
}

// DeleteJobResult represents the DeleteJobResult schema from the OpenAPI specification
type DeleteJobResult struct {
	Jobsummary JobSummary `json:"jobSummary"` // Describes the summary for an execution job for an Amplify app.
}

// UpdateDomainAssociationRequest represents the UpdateDomainAssociationRequest schema from the OpenAPI specification
type UpdateDomainAssociationRequest struct {
	Subdomainsettings interface{} `json:"subDomainSettings,omitempty"`
	Autosubdomaincreationpatterns interface{} `json:"autoSubDomainCreationPatterns,omitempty"`
	Autosubdomainiamrole interface{} `json:"autoSubDomainIAMRole,omitempty"`
	Enableautosubdomain interface{} `json:"enableAutoSubDomain,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// SubDomainSetting represents the SubDomainSetting schema from the OpenAPI specification
type SubDomainSetting struct {
	Branchname interface{} `json:"branchName"`
	Prefix interface{} `json:"prefix"`
}

// ListWebhooksRequest represents the ListWebhooksRequest schema from the OpenAPI specification
type ListWebhooksRequest struct {
}

// CreateApprequest represents the CreateApprequest schema from the OpenAPI specification
type CreateApprequest struct {
	Basicauthcredentials string `json:"basicAuthCredentials,omitempty"` // The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.
	Customheaders string `json:"customHeaders,omitempty"` // The custom HTTP headers for an Amplify app.
	Buildspec string `json:"buildSpec,omitempty"` // The build specification (build spec) file for an Amplify app build.
	Enablebranchautodeletion bool `json:"enableBranchAutoDeletion,omitempty"` // Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.
	Iamservicerolearn string `json:"iamServiceRoleArn,omitempty"` // The AWS Identity and Access Management (IAM) service role for an Amplify app.
	Accesstoken string `json:"accessToken,omitempty"` // <p>The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.</p> <p>Use <code>accessToken</code> for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use <code>oauthToken</code>.</p> <p>You must specify either <code>accessToken</code> or <code>oauthToken</code> when you create a new app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href="https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>
	Autobranchcreationpatterns []string `json:"autoBranchCreationPatterns,omitempty"` // The automated branch creation glob patterns for an Amplify app.
	Description string `json:"description,omitempty"` // The description for an Amplify app.
	Enablebasicauth bool `json:"enableBasicAuth,omitempty"` // Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.
	Name string `json:"name"` // The name for an Amplify app.
	Platform string `json:"platform,omitempty"` // The platform for the Amplify app. For a static app, set the platform type to <code>WEB</code>. For a dynamic server-side rendered (SSR) app, set the platform type to <code>WEB_COMPUTE</code>. For an app requiring Amplify Hosting's original SSR support only, set the platform type to <code>WEB_DYNAMIC</code>.
	Tags map[string]interface{} `json:"tags,omitempty"` // The tag for an Amplify app.
	Autobranchcreationconfig CreateApprequestautoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty"` // Describes the automated branch creation configuration.
	Customrules []CustomRule `json:"customRules,omitempty"` // The custom rewrite and redirect rules for an Amplify app.
	Oauthtoken string `json:"oauthToken,omitempty"` // <p>The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.</p> <p>Use <code>oauthToken</code> for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use <code>accessToken</code>.</p> <p>You must specify either <code>oauthToken</code> or <code>accessToken</code> when you create a new app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href="https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>
	Environmentvariables map[string]interface{} `json:"environmentVariables,omitempty"` // The environment variables map for an Amplify app.
	Repository string `json:"repository,omitempty"` // The repository for an Amplify app.
	Enableautobranchcreation bool `json:"enableAutoBranchCreation,omitempty"` // Enables automated branch creation for an Amplify app.
	Enablebranchautobuild bool `json:"enableBranchAutoBuild,omitempty"` // Enables the auto building of branches for an Amplify app.
}

// AppproductionBranch represents the AppproductionBranch schema from the OpenAPI specification
type AppproductionBranch struct {
	Lastdeploytime interface{} `json:"lastDeployTime,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Thumbnailurl interface{} `json:"thumbnailUrl,omitempty"`
	Branchname interface{} `json:"branchName,omitempty"`
}

// StartJobrequest represents the StartJobrequest schema from the OpenAPI specification
type StartJobrequest struct {
	Jobid string `json:"jobId,omitempty"` // The unique ID for an existing job. This is required if the value of <code>jobType</code> is <code>RETRY</code>.
	Jobreason string `json:"jobReason,omitempty"` // A descriptive reason for starting this job.
	Jobtype string `json:"jobType"` // Describes the type for the job. The job type <code>RELEASE</code> starts a new job with the latest change from the specified branch. This value is available only for apps that are connected to a repository. The job type <code>RETRY</code> retries an existing job. If the job type value is <code>RETRY</code>, the <code>jobId</code> is also required.
	Commitid string `json:"commitId,omitempty"` // The commit ID from a third-party repository provider for the job.
	Commitmessage string `json:"commitMessage,omitempty"` // The commit message from a third-party repository provider for the job.
	Committime string `json:"commitTime,omitempty"` // The commit date and time for the job.
}

// GetArtifactUrlResult represents the GetArtifactUrlResult schema from the OpenAPI specification
type GetArtifactUrlResult struct {
	Artifacturl interface{} `json:"artifactUrl"`
	Artifactid interface{} `json:"artifactId"`
}

// StartDeploymentResult represents the StartDeploymentResult schema from the OpenAPI specification
type StartDeploymentResult struct {
	Jobsummary StartDeploymentResultjobSummary `json:"jobSummary"`
}

// JobSummary represents the JobSummary schema from the OpenAPI specification
type JobSummary struct {
	Jobarn interface{} `json:"jobArn"`
	Jobtype interface{} `json:"jobType"`
	Commitid interface{} `json:"commitId"`
	Status interface{} `json:"status"`
	Committime interface{} `json:"commitTime"`
	Endtime interface{} `json:"endTime,omitempty"`
	Jobid interface{} `json:"jobId"`
	Starttime interface{} `json:"startTime"`
	Commitmessage interface{} `json:"commitMessage"`
}

// GetJobResult represents the GetJobResult schema from the OpenAPI specification
type GetJobResult struct {
	Job Job `json:"job"` // Describes an execution job for an Amplify app.
}

// CreateBranchResultbranch represents the CreateBranchResultbranch schema from the OpenAPI specification
type CreateBranchResultbranch struct {
	Environmentvariables interface{} `json:"environmentVariables"`
	Displayname interface{} `json:"displayName"`
	Totalnumberofjobs interface{} `json:"totalNumberOfJobs"`
	Branchname interface{} `json:"branchName"`
	Description interface{} `json:"description"`
	Activejobid interface{} `json:"activeJobId"`
	Enablenotification interface{} `json:"enableNotification"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview"`
	Sourcebranch interface{} `json:"sourceBranch,omitempty"`
	Associatedresources interface{} `json:"associatedResources,omitempty"`
	Updatetime interface{} `json:"updateTime"`
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild"`
	Createtime interface{} `json:"createTime"`
	Destinationbranch interface{} `json:"destinationBranch,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Stage interface{} `json:"stage"`
	Framework interface{} `json:"framework"`
	Thumbnailurl interface{} `json:"thumbnailUrl,omitempty"`
	Brancharn interface{} `json:"branchArn"`
	Customdomains interface{} `json:"customDomains"`
	Ttl interface{} `json:"ttl"`
	Enablebasicauth interface{} `json:"enableBasicAuth"`
}

// Jobsummary represents the Jobsummary schema from the OpenAPI specification
type Jobsummary struct {
	Jobid interface{} `json:"jobId"`
	Starttime interface{} `json:"startTime"`
	Commitmessage interface{} `json:"commitMessage"`
	Jobarn interface{} `json:"jobArn"`
	Jobtype interface{} `json:"jobType"`
	Commitid interface{} `json:"commitId"`
	Status interface{} `json:"status"`
	Committime interface{} `json:"commitTime"`
	Endtime interface{} `json:"endTime,omitempty"`
}

// UpdateAppResult represents the UpdateAppResult schema from the OpenAPI specification
type UpdateAppResult struct {
	App UpdateAppResultapp `json:"app"`
}

// ListBackendEnvironmentsRequest represents the ListBackendEnvironmentsRequest schema from the OpenAPI specification
type ListBackendEnvironmentsRequest struct {
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// Step represents the Step schema from the OpenAPI specification
type Step struct {
	Starttime interface{} `json:"startTime"`
	Statusreason interface{} `json:"statusReason,omitempty"`
	Testconfigurl interface{} `json:"testConfigUrl,omitempty"`
	Artifactsurl interface{} `json:"artifactsUrl,omitempty"`
	Endtime interface{} `json:"endTime"`
	Context interface{} `json:"context,omitempty"`
	Logurl interface{} `json:"logUrl,omitempty"`
	Stepname interface{} `json:"stepName"`
	Testartifactsurl interface{} `json:"testArtifactsUrl,omitempty"`
	Screenshots interface{} `json:"screenshots,omitempty"`
	Status interface{} `json:"status"`
}

// CreateAppRequest represents the CreateAppRequest schema from the OpenAPI specification
type CreateAppRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Customheaders interface{} `json:"customHeaders,omitempty"`
	Iamservicerolearn interface{} `json:"iamServiceRoleArn,omitempty"`
	Autobranchcreationpatterns interface{} `json:"autoBranchCreationPatterns,omitempty"`
	Enablebranchautobuild interface{} `json:"enableBranchAutoBuild,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Enableautobranchcreation interface{} `json:"enableAutoBranchCreation,omitempty"`
	Name interface{} `json:"name"`
	Autobranchcreationconfig CreateAppRequestautoBranchCreationConfig `json:"autoBranchCreationConfig,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Description interface{} `json:"description,omitempty"`
	Enablebranchautodeletion interface{} `json:"enableBranchAutoDeletion,omitempty"`
	Accesstoken interface{} `json:"accessToken,omitempty"`
	Platform interface{} `json:"platform,omitempty"`
	Repository interface{} `json:"repository,omitempty"`
	Environmentvariables interface{} `json:"environmentVariables,omitempty"`
	Enablebasicauth interface{} `json:"enableBasicAuth,omitempty"`
	Oauthtoken interface{} `json:"oauthToken,omitempty"`
	Customrules interface{} `json:"customRules,omitempty"`
}

// CreateBranchResult represents the CreateBranchResult schema from the OpenAPI specification
type CreateBranchResult struct {
	Branch CreateBranchResultbranch `json:"branch"`
}

// CreateWebhookResultwebhook represents the CreateWebhookResultwebhook schema from the OpenAPI specification
type CreateWebhookResultwebhook struct {
	Branchname interface{} `json:"branchName"`
	Createtime interface{} `json:"createTime"`
	Description interface{} `json:"description"`
	Updatetime interface{} `json:"updateTime"`
	Webhookarn interface{} `json:"webhookArn"`
	Webhookid interface{} `json:"webhookId"`
	Webhookurl interface{} `json:"webhookUrl"`
}

// Branch represents the Branch schema from the OpenAPI specification
type Branch struct {
	Brancharn interface{} `json:"branchArn"`
	Customdomains interface{} `json:"customDomains"`
	Ttl interface{} `json:"ttl"`
	Enablebasicauth interface{} `json:"enableBasicAuth"`
	Environmentvariables interface{} `json:"environmentVariables"`
	Displayname interface{} `json:"displayName"`
	Totalnumberofjobs interface{} `json:"totalNumberOfJobs"`
	Branchname interface{} `json:"branchName"`
	Description interface{} `json:"description"`
	Activejobid interface{} `json:"activeJobId"`
	Enablenotification interface{} `json:"enableNotification"`
	Enableperformancemode interface{} `json:"enablePerformanceMode,omitempty"`
	Enablepullrequestpreview interface{} `json:"enablePullRequestPreview"`
	Sourcebranch interface{} `json:"sourceBranch,omitempty"`
	Associatedresources interface{} `json:"associatedResources,omitempty"`
	Updatetime interface{} `json:"updateTime"`
	Backendenvironmentarn interface{} `json:"backendEnvironmentArn,omitempty"`
	Buildspec interface{} `json:"buildSpec,omitempty"`
	Basicauthcredentials interface{} `json:"basicAuthCredentials,omitempty"`
	Enableautobuild interface{} `json:"enableAutoBuild"`
	Createtime interface{} `json:"createTime"`
	Destinationbranch interface{} `json:"destinationBranch,omitempty"`
	Tags interface{} `json:"tags,omitempty"`
	Pullrequestenvironmentname interface{} `json:"pullRequestEnvironmentName,omitempty"`
	Stage interface{} `json:"stage"`
	Framework interface{} `json:"framework"`
	Thumbnailurl interface{} `json:"thumbnailUrl,omitempty"`
}

// GenerateAccessLogsrequest represents the GenerateAccessLogsrequest schema from the OpenAPI specification
type GenerateAccessLogsrequest struct {
	Domainname string `json:"domainName"` // The name of the domain.
	Endtime string `json:"endTime,omitempty"` // The time at which the logs should end. The time range specified is inclusive of the end time.
	Starttime string `json:"startTime,omitempty"` // The time at which the logs should start. The time range specified is inclusive of the start time.
}

// ListWebhooksResult represents the ListWebhooksResult schema from the OpenAPI specification
type ListWebhooksResult struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Webhooks interface{} `json:"webhooks"`
}

// DeleteBranchRequest represents the DeleteBranchRequest schema from the OpenAPI specification
type DeleteBranchRequest struct {
}

// GetDomainAssociationResult represents the GetDomainAssociationResult schema from the OpenAPI specification
type GetDomainAssociationResult struct {
	Domainassociation CreateDomainAssociationResultdomainAssociation `json:"domainAssociation"`
}

// GetAppResult represents the GetAppResult schema from the OpenAPI specification
type GetAppResult struct {
	App App `json:"app"` // Represents the different branches of a repository for building, deploying, and hosting an Amplify app.
}

// FileMap represents the FileMap schema from the OpenAPI specification
type FileMap struct {
}
