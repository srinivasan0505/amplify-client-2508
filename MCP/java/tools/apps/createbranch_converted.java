/**
 * MCP Server function for Creates a new branch for an Amplify app.
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

class Post_Apps_App_Id_BranchesMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Apps_App_Id_BranchesHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("buildSpec")) {
            queryParams.add("buildSpec=" + args.get("buildSpec"));
        }
        if (args.containsKey("basicAuthCredentials")) {
            queryParams.add("basicAuthCredentials=" + args.get("basicAuthCredentials"));
        }
        if (args.containsKey("ttl")) {
            queryParams.add("ttl=" + args.get("ttl"));
        }
        if (args.containsKey("displayName")) {
            queryParams.add("displayName=" + args.get("displayName"));
        }
        if (args.containsKey("backendEnvironmentArn")) {
            queryParams.add("backendEnvironmentArn=" + args.get("backendEnvironmentArn"));
        }
        if (args.containsKey("framework")) {
            queryParams.add("framework=" + args.get("framework"));
        }
        if (args.containsKey("description")) {
            queryParams.add("description=" + args.get("description"));
        }
        if (args.containsKey("stage")) {
            queryParams.add("stage=" + args.get("stage"));
        }
        if (args.containsKey("branchName")) {
            queryParams.add("branchName=" + args.get("branchName"));
        }
        if (args.containsKey("pullRequestEnvironmentName")) {
            queryParams.add("pullRequestEnvironmentName=" + args.get("pullRequestEnvironmentName"));
        }
        if (args.containsKey("enableAutoBuild")) {
            queryParams.add("enableAutoBuild=" + args.get("enableAutoBuild"));
        }
        if (args.containsKey("enableNotification")) {
            queryParams.add("enableNotification=" + args.get("enableNotification"));
        }
        if (args.containsKey("enableBasicAuth")) {
            queryParams.add("enableBasicAuth=" + args.get("enableBasicAuth"));
        }
        if (args.containsKey("enablePullRequestPreview")) {
            queryParams.add("enablePullRequestPreview=" + args.get("enablePullRequestPreview"));
        }
        if (args.containsKey("enablePerformanceMode")) {
            queryParams.add("enablePerformanceMode=" + args.get("enablePerformanceMode"));
        }
        if (args.containsKey("environmentVariables")) {
            queryParams.add("environmentVariables=" + args.get("environmentVariables"));
        }
        if (args.containsKey("tags")) {
            queryParams.add("tags=" + args.get("tags"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_apps_app_id_branches" + queryString;
                
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
    
    public static MCPServer.Tool createPost_Apps_App_Id_BranchesTool(MCPServer.APIConfig config) {
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
        Map<String, Object> buildSpecProperty = new HashMap<>();
        buildSpecProperty.put("type", "string");
        buildSpecProperty.put("required", false);
        buildSpecProperty.put("description", "Input parameter: The build specification (build spec) file for an Amplify app build.");
        properties.put("buildSpec", buildSpecProperty);
        Map<String, Object> basicAuthCredentialsProperty = new HashMap<>();
        basicAuthCredentialsProperty.put("type", "string");
        basicAuthCredentialsProperty.put("required", false);
        basicAuthCredentialsProperty.put("description", "Input parameter: The basic authorization credentials for the branch. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.");
        properties.put("basicAuthCredentials", basicAuthCredentialsProperty);
        Map<String, Object> ttlProperty = new HashMap<>();
        ttlProperty.put("type", "string");
        ttlProperty.put("required", false);
        ttlProperty.put("description", "Input parameter: The content Time to Live (TTL) for the website in seconds.");
        properties.put("ttl", ttlProperty);
        Map<String, Object> displayNameProperty = new HashMap<>();
        displayNameProperty.put("type", "string");
        displayNameProperty.put("required", false);
        displayNameProperty.put("description", "Input parameter: The display name for a branch. This is used as the default domain prefix.");
        properties.put("displayName", displayNameProperty);
        Map<String, Object> backendEnvironmentArnProperty = new HashMap<>();
        backendEnvironmentArnProperty.put("type", "string");
        backendEnvironmentArnProperty.put("required", false);
        backendEnvironmentArnProperty.put("description", "Input parameter: The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app.");
        properties.put("backendEnvironmentArn", backendEnvironmentArnProperty);
        Map<String, Object> frameworkProperty = new HashMap<>();
        frameworkProperty.put("type", "string");
        frameworkProperty.put("required", false);
        frameworkProperty.put("description", "Input parameter: The framework for the branch.");
        properties.put("framework", frameworkProperty);
        Map<String, Object> descriptionProperty = new HashMap<>();
        descriptionProperty.put("type", "string");
        descriptionProperty.put("required", false);
        descriptionProperty.put("description", "Input parameter: The description for the branch.");
        properties.put("description", descriptionProperty);
        Map<String, Object> stageProperty = new HashMap<>();
        stageProperty.put("type", "string");
        stageProperty.put("required", false);
        stageProperty.put("description", "Input parameter: Describes the current stage for the branch.");
        properties.put("stage", stageProperty);
        Map<String, Object> branchNameProperty = new HashMap<>();
        branchNameProperty.put("type", "string");
        branchNameProperty.put("required", true);
        branchNameProperty.put("description", "Input parameter: The name for the branch.");
        properties.put("branchName", branchNameProperty);
        Map<String, Object> pullRequestEnvironmentNameProperty = new HashMap<>();
        pullRequestEnvironmentNameProperty.put("type", "string");
        pullRequestEnvironmentNameProperty.put("required", false);
        pullRequestEnvironmentNameProperty.put("description", "Input parameter: The Amplify environment name for the pull request.");
        properties.put("pullRequestEnvironmentName", pullRequestEnvironmentNameProperty);
        Map<String, Object> enableAutoBuildProperty = new HashMap<>();
        enableAutoBuildProperty.put("type", "string");
        enableAutoBuildProperty.put("required", false);
        enableAutoBuildProperty.put("description", "Input parameter: Enables auto building for the branch.");
        properties.put("enableAutoBuild", enableAutoBuildProperty);
        Map<String, Object> enableNotificationProperty = new HashMap<>();
        enableNotificationProperty.put("type", "string");
        enableNotificationProperty.put("required", false);
        enableNotificationProperty.put("description", "Input parameter: Enables notifications for the branch.");
        properties.put("enableNotification", enableNotificationProperty);
        Map<String, Object> enableBasicAuthProperty = new HashMap<>();
        enableBasicAuthProperty.put("type", "string");
        enableBasicAuthProperty.put("required", false);
        enableBasicAuthProperty.put("description", "Input parameter: Enables basic authorization for the branch.");
        properties.put("enableBasicAuth", enableBasicAuthProperty);
        Map<String, Object> enablePullRequestPreviewProperty = new HashMap<>();
        enablePullRequestPreviewProperty.put("type", "string");
        enablePullRequestPreviewProperty.put("required", false);
        enablePullRequestPreviewProperty.put("description", "Input parameter: Enables pull request previews for this branch.");
        properties.put("enablePullRequestPreview", enablePullRequestPreviewProperty);
        Map<String, Object> enablePerformanceModeProperty = new HashMap<>();
        enablePerformanceModeProperty.put("type", "string");
        enablePerformanceModeProperty.put("required", false);
        enablePerformanceModeProperty.put("description", "Input parameter: <p>Enables performance mode for the branch.</p> <p>Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out. </p>");
        properties.put("enablePerformanceMode", enablePerformanceModeProperty);
        Map<String, Object> environmentVariablesProperty = new HashMap<>();
        environmentVariablesProperty.put("type", "string");
        environmentVariablesProperty.put("required", false);
        environmentVariablesProperty.put("description", "Input parameter: The environment variables for the branch.");
        properties.put("environmentVariables", environmentVariablesProperty);
        Map<String, Object> tagsProperty = new HashMap<>();
        tagsProperty.put("type", "string");
        tagsProperty.put("required", false);
        tagsProperty.put("description", "Input parameter: The tag for the branch.");
        properties.put("tags", tagsProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_apps_app_id_branches",
            "Creates a new branch for an Amplify app.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Apps_App_Id_BranchesHandler(config));
    }
    
}