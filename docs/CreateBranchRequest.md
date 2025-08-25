

# CreateBranchRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**branchName** | **String** |  The name for the branch.  |  |
|**description** | **String** |  The description for the branch.  |  [optional] |
|**stage** | [**StageEnum**](#StageEnum) |  Describes the current stage for the branch.  |  [optional] |
|**framework** | **String** |  The framework for the branch.  |  [optional] |
|**enableNotification** | **Boolean** |  Enables notifications for the branch.  |  [optional] |
|**enableAutoBuild** | **Boolean** |  Enables auto building for the branch.  |  [optional] |
|**environmentVariables** | **Map&lt;String, String&gt;** |  The environment variables for the branch.  |  [optional] |
|**basicAuthCredentials** | **String** |  The basic authorization credentials for the branch. You must base64-encode the authorization credentials and provide them in the format &lt;code&gt;user:password&lt;/code&gt;. |  [optional] |
|**enableBasicAuth** | **Boolean** |  Enables basic authorization for the branch.  |  [optional] |
|**enablePerformanceMode** | **Boolean** | &lt;p&gt;Enables performance mode for the branch.&lt;/p&gt; &lt;p&gt;Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out. &lt;/p&gt; |  [optional] |
|**tags** | **Map&lt;String, String&gt;** |  The tag for the branch.  |  [optional] |
|**buildSpec** | **String** |  The build specification (build spec) file for an Amplify app build.  |  [optional] |
|**ttl** | **String** |  The content Time to Live (TTL) for the website in seconds.  |  [optional] |
|**displayName** | **String** |  The display name for a branch. This is used as the default domain prefix.  |  [optional] |
|**enablePullRequestPreview** | **Boolean** |  Enables pull request previews for this branch.  |  [optional] |
|**pullRequestEnvironmentName** | **String** |  The Amplify environment name for the pull request.  |  [optional] |
|**backendEnvironmentArn** | **String** |  The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.  |  [optional] |



## Enum: StageEnum

| Name | Value |
|---- | -----|
| PRODUCTION | &quot;PRODUCTION&quot; |
| BETA | &quot;BETA&quot; |
| DEVELOPMENT | &quot;DEVELOPMENT&quot; |
| EXPERIMENTAL | &quot;EXPERIMENTAL&quot; |
| PULL_REQUEST | &quot;PULL_REQUEST&quot; |



