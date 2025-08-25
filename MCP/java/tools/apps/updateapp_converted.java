/**
 * MCP Server function for Updates an existing Amplify app.
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_Apps_App_IdMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Apps_App_IdHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("appId")) {
            queryParams.add("appId=" + args.get("appId"));
        }
        if (args.containsKey("customHeaders")) {
            queryParams.add("customHeaders=" + args.get("customHeaders"));
        }
        if (args.containsKey("iamServiceRoleArn")) {
            queryParams.add("iamServiceRoleArn=" + args.get("iamServiceRoleArn"));
        }
        if (args.containsKey("basicAuthCredentials")) {
            queryParams.add("basicAuthCredentials=" + args.get("basicAuthCredentials"));
        }
        if (args.containsKey("repository")) {
            queryParams.add("repository=" + args.get("repository"));
        }
        if (args.containsKey("name")) {
            queryParams.add("name=" + args.get("name"));
        }
        if (args.containsKey("platform")) {
            queryParams.add("platform=" + args.get("platform"));
        }
        if (args.containsKey("accessToken")) {
            queryParams.add("accessToken=" + args.get("accessToken"));
        }
        if (args.containsKey("buildSpec")) {
            queryParams.add("buildSpec=" + args.get("buildSpec"));
        }
        if (args.containsKey("oauthToken")) {
            queryParams.add("oauthToken=" + args.get("oauthToken"));
        }
        if (args.containsKey("description")) {
            queryParams.add("description=" + args.get("description"));
        }
        if (args.containsKey("enableAutoBranchCreation")) {
            queryParams.add("enableAutoBranchCreation=" + args.get("enableAutoBranchCreation"));
        }
        if (args.containsKey("enableBranchAutoBuild")) {
            queryParams.add("enableBranchAutoBuild=" + args.get("enableBranchAutoBuild"));
        }
        if (args.containsKey("enableBasicAuth")) {
            queryParams.add("enableBasicAuth=" + args.get("enableBasicAuth"));
        }
        if (args.containsKey("enableBranchAutoDeletion")) {
            queryParams.add("enableBranchAutoDeletion=" + args.get("enableBranchAutoDeletion"));
        }
        if (args.containsKey("autoBranchCreationConfig")) {
            queryParams.add("autoBranchCreationConfig=" + args.get("autoBranchCreationConfig"));
        }
        if (args.containsKey("environmentVariables")) {
            queryParams.add("environmentVariables=" + args.get("environmentVariables"));
        }
        if (args.containsKey("autoBranchCreationPatterns")) {
            queryParams.add("autoBranchCreationPatterns=" + args.get("autoBranchCreationPatterns"));
        }
        if (args.containsKey("customRules")) {
            queryParams.add("customRules=" + args.get("customRules"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_apps_app_id" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_Apps_App_IdTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> appIdProperty = new HashMap<>();
        appIdProperty.put("type", "string");
        appIdProperty.put("required", true);
        appIdProperty.put("description", "The unique ID for an Amplify app.");
        properties.put("appId", appIdProperty);
        Map<String, Object> customHeadersProperty = new HashMap<>();
        customHeadersProperty.put("type", "string");
        customHeadersProperty.put("required", false);
        customHeadersProperty.put("description", "Input parameter: The custom HTTP headers for an Amplify app.");
        properties.put("customHeaders", customHeadersProperty);
        Map<String, Object> iamServiceRoleArnProperty = new HashMap<>();
        iamServiceRoleArnProperty.put("type", "string");
        iamServiceRoleArnProperty.put("required", false);
        iamServiceRoleArnProperty.put("description", "Input parameter: The AWS Identity and Access Management (IAM) service role for an Amplify app.");
        properties.put("iamServiceRoleArn", iamServiceRoleArnProperty);
        Map<String, Object> basicAuthCredentialsProperty = new HashMap<>();
        basicAuthCredentialsProperty.put("type", "string");
        basicAuthCredentialsProperty.put("required", false);
        basicAuthCredentialsProperty.put("description", "Input parameter: The basic authorization credentials for an Amplify app. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.");
        properties.put("basicAuthCredentials", basicAuthCredentialsProperty);
        Map<String, Object> repositoryProperty = new HashMap<>();
        repositoryProperty.put("type", "string");
        repositoryProperty.put("required", false);
        repositoryProperty.put("description", "Input parameter: The name of the repository for an Amplify app");
        properties.put("repository", repositoryProperty);
        Map<String, Object> nameProperty = new HashMap<>();
        nameProperty.put("type", "string");
        nameProperty.put("required", false);
        nameProperty.put("description", "Input parameter: The name for an Amplify app.");
        properties.put("name", nameProperty);
        Map<String, Object> platformProperty = new HashMap<>();
        platformProperty.put("type", "string");
        platformProperty.put("required", false);
        platformProperty.put("description", "Input parameter: The platform for the Amplify app. For a static app, set the platform type to <code>WEB</code>. For a dynamic server-side rendered (SSR) app, set the platform type to <code>WEB_COMPUTE</code>. For an app requiring Amplify Hosting's original SSR support only, set the platform type to <code>WEB_DYNAMIC</code>.");
        properties.put("platform", platformProperty);
        Map<String, Object> accessTokenProperty = new HashMap<>();
        accessTokenProperty.put("type", "string");
        accessTokenProperty.put("required", false);
        accessTokenProperty.put("description", "Input parameter: <p>The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.</p> <p>Use <code>accessToken</code> for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use <code>oauthToken</code>.</p> <p>You must specify either <code>accessToken</code> or <code>oauthToken</code> when you update an app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href=\"https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth\">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>");
        properties.put("accessToken", accessTokenProperty);
        Map<String, Object> buildSpecProperty = new HashMap<>();
        buildSpecProperty.put("type", "string");
        buildSpecProperty.put("required", false);
        buildSpecProperty.put("description", "Input parameter: The build specification (build spec) file for an Amplify app build.");
        properties.put("buildSpec", buildSpecProperty);
        Map<String, Object> oauthTokenProperty = new HashMap<>();
        oauthTokenProperty.put("type", "string");
        oauthTokenProperty.put("required", false);
        oauthTokenProperty.put("description", "Input parameter: <p>The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.</p> <p>Use <code>oauthToken</code> for repository providers other than GitHub, such as Bitbucket or CodeCommit.</p> <p>To authorize access to GitHub as your repository provider, use <code>accessToken</code>.</p> <p>You must specify either <code>oauthToken</code> or <code>accessToken</code> when you update an app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href=\"https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth\">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>");
        properties.put("oauthToken", oauthTokenProperty);
        Map<String, Object> descriptionProperty = new HashMap<>();
        descriptionProperty.put("type", "string");
        descriptionProperty.put("required", false);
        descriptionProperty.put("description", "Input parameter: The description for an Amplify app.");
        properties.put("description", descriptionProperty);
        Map<String, Object> enableAutoBranchCreationProperty = new HashMap<>();
        enableAutoBranchCreationProperty.put("type", "string");
        enableAutoBranchCreationProperty.put("required", false);
        enableAutoBranchCreationProperty.put("description", "Input parameter: Enables automated branch creation for an Amplify app.");
        properties.put("enableAutoBranchCreation", enableAutoBranchCreationProperty);
        Map<String, Object> enableBranchAutoBuildProperty = new HashMap<>();
        enableBranchAutoBuildProperty.put("type", "string");
        enableBranchAutoBuildProperty.put("required", false);
        enableBranchAutoBuildProperty.put("description", "Input parameter: Enables branch auto-building for an Amplify app.");
        properties.put("enableBranchAutoBuild", enableBranchAutoBuildProperty);
        Map<String, Object> enableBasicAuthProperty = new HashMap<>();
        enableBasicAuthProperty.put("type", "string");
        enableBasicAuthProperty.put("required", false);
        enableBasicAuthProperty.put("description", "Input parameter: Enables basic authorization for an Amplify app.");
        properties.put("enableBasicAuth", enableBasicAuthProperty);
        Map<String, Object> enableBranchAutoDeletionProperty = new HashMap<>();
        enableBranchAutoDeletionProperty.put("type", "string");
        enableBranchAutoDeletionProperty.put("required", false);
        enableBranchAutoDeletionProperty.put("description", "Input parameter: Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository.");
        properties.put("enableBranchAutoDeletion", enableBranchAutoDeletionProperty);
        Map<String, Object> autoBranchCreationConfigProperty = new HashMap<>();
        autoBranchCreationConfigProperty.put("type", "string");
        autoBranchCreationConfigProperty.put("required", false);
        autoBranchCreationConfigProperty.put("description", "Input parameter: Describes the automated branch creation configuration.");
        properties.put("autoBranchCreationConfig", autoBranchCreationConfigProperty);
        Map<String, Object> environmentVariablesProperty = new HashMap<>();
        environmentVariablesProperty.put("type", "string");
        environmentVariablesProperty.put("required", false);
        environmentVariablesProperty.put("description", "Input parameter: The environment variables for an Amplify app.");
        properties.put("environmentVariables", environmentVariablesProperty);
        Map<String, Object> autoBranchCreationPatternsProperty = new HashMap<>();
        autoBranchCreationPatternsProperty.put("type", "string");
        autoBranchCreationPatternsProperty.put("required", false);
        autoBranchCreationPatternsProperty.put("description", "Input parameter: Describes the automated branch creation glob patterns for an Amplify app.");
        properties.put("autoBranchCreationPatterns", autoBranchCreationPatternsProperty);
        Map<String, Object> customRulesProperty = new HashMap<>();
        customRulesProperty.put("type", "string");
        customRulesProperty.put("required", false);
        customRulesProperty.put("description", "Input parameter: The custom redirect and rewrite rules for an Amplify app.");
        properties.put("customRules", customRulesProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_apps_app_id",
            "Updates an existing Amplify app.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Apps_App_IdHandler(config));
    }
    
}