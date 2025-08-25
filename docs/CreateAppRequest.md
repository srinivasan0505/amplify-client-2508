

# CreateAppRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**name** | **String** |  The name for an Amplify app.  |  |
|**description** | **String** |  The description for an Amplify app.  |  [optional] |
|**repository** | **String** |  The repository for an Amplify app.  |  [optional] |
|**platform** | [**PlatformEnum**](#PlatformEnum) |  The platform for the Amplify app. For a static app, set the platform type to &lt;code&gt;WEB&lt;/code&gt;. For a dynamic server-side rendered (SSR) app, set the platform type to &lt;code&gt;WEB_COMPUTE&lt;/code&gt;. For an app requiring Amplify Hosting&#39;s original SSR support only, set the platform type to &lt;code&gt;WEB_DYNAMIC&lt;/code&gt;. |  [optional] |
|**iamServiceRoleArn** | **String** |  The AWS Identity and Access Management (IAM) service role for an Amplify app.  |  [optional] |
|**oauthToken** | **String** | &lt;p&gt;The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.&lt;/p&gt; &lt;p&gt;Use &lt;code&gt;oauthToken&lt;/code&gt; for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use &lt;code&gt;accessToken&lt;/code&gt;.&lt;/p&gt; &lt;p&gt;You must specify either &lt;code&gt;oauthToken&lt;/code&gt; or &lt;code&gt;accessToken&lt;/code&gt; when you create a new app.&lt;/p&gt; &lt;p&gt;Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth\&quot;&gt;Migrating an existing OAuth app to the Amplify GitHub App&lt;/a&gt; in the &lt;i&gt;Amplify User Guide&lt;/i&gt; .&lt;/p&gt; |  [optional] |
|**accessToken** | **String** | &lt;p&gt;The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.&lt;/p&gt; &lt;p&gt;Use &lt;code&gt;accessToken&lt;/code&gt; for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use &lt;code&gt;oauthToken&lt;/code&gt;.&lt;/p&gt; &lt;p&gt;You must specify either &lt;code&gt;accessToken&lt;/code&gt; or &lt;code&gt;oauthToken&lt;/code&gt; when you create a new app.&lt;/p&gt; &lt;p&gt;Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth\&quot;&gt;Migrating an existing OAuth app to the Amplify GitHub App&lt;/a&gt; in the &lt;i&gt;Amplify User Guide&lt;/i&gt; .&lt;/p&gt; |  [optional] |
|**environmentVariables** | **Map&lt;String, String&gt;** |  The environment variables map for an Amplify app.  |  [optional] |
|**enableBranchAutoBuild** | **Boolean** |  Enables the auto building of branches for an Amplify app.  |  [optional] |
|**enableBranchAutoDeletion** | **Boolean** |  Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.  |  [optional] |
|**enableBasicAuth** | **Boolean** |  Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app.  |  [optional] |
|**basicAuthCredentials** | **String** |  The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format &lt;code&gt;user:password&lt;/code&gt;. |  [optional] |
|**customRules** | [**List&lt;CustomRule&gt;**](CustomRule.md) |  The custom rewrite and redirect rules for an Amplify app.  |  [optional] |
|**tags** | **Map&lt;String, String&gt;** |  The tag for an Amplify app.  |  [optional] |
|**buildSpec** | **String** |  The build specification (build spec) file for an Amplify app build.  |  [optional] |
|**customHeaders** | **String** | The custom HTTP headers for an Amplify app. |  [optional] |
|**enableAutoBranchCreation** | **Boolean** |  Enables automated branch creation for an Amplify app.  |  [optional] |
|**autoBranchCreationPatterns** | **List&lt;String&gt;** |  The automated branch creation glob patterns for an Amplify app.  |  [optional] |
|**autoBranchCreationConfig** | [**CreateAppRequestAutoBranchCreationConfig**](CreateAppRequestAutoBranchCreationConfig.md) |  |  [optional] |



## Enum: PlatformEnum

| Name | Value |
|---- | -----|
| WEB | &quot;WEB&quot; |
| WEB_DYNAMIC | &quot;WEB_DYNAMIC&quot; |
| WEB_COMPUTE | &quot;WEB_COMPUTE&quot; |



