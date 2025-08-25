/**
 * MCP Server function for Starts a new job for a branch of an Amplify app.
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

class Post_Apps_App_Id_Branches_Branch_Name_JobsMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_Apps_App_Id_Branches_Branch_Name_JobsHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("branchName")) {
            queryParams.add("branchName=" + args.get("branchName"));
        }
        if (args.containsKey("commitId")) {
            queryParams.add("commitId=" + args.get("commitId"));
        }
        if (args.containsKey("commitMessage")) {
            queryParams.add("commitMessage=" + args.get("commitMessage"));
        }
        if (args.containsKey("commitTime")) {
            queryParams.add("commitTime=" + args.get("commitTime"));
        }
        if (args.containsKey("jobId")) {
            queryParams.add("jobId=" + args.get("jobId"));
        }
        if (args.containsKey("jobReason")) {
            queryParams.add("jobReason=" + args.get("jobReason"));
        }
        if (args.containsKey("jobType")) {
            queryParams.add("jobType=" + args.get("jobType"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_apps_app_id_branches_branch_name_jobs" + queryString;
                
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
    
    public static MCPServer.Tool createPost_Apps_App_Id_Branches_Branch_Name_JobsTool(MCPServer.APIConfig config) {
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
        Map<String, Object> branchNameProperty = new HashMap<>();
        branchNameProperty.put("type", "string");
        branchNameProperty.put("required", true);
        branchNameProperty.put("description", "The branch name for the job.");
        properties.put("branchName", branchNameProperty);
        Map<String, Object> commitIdProperty = new HashMap<>();
        commitIdProperty.put("type", "string");
        commitIdProperty.put("required", false);
        commitIdProperty.put("description", "Input parameter: The commit ID from a third-party repository provider for the job.");
        properties.put("commitId", commitIdProperty);
        Map<String, Object> commitMessageProperty = new HashMap<>();
        commitMessageProperty.put("type", "string");
        commitMessageProperty.put("required", false);
        commitMessageProperty.put("description", "Input parameter: The commit message from a third-party repository provider for the job.");
        properties.put("commitMessage", commitMessageProperty);
        Map<String, Object> commitTimeProperty = new HashMap<>();
        commitTimeProperty.put("type", "string");
        commitTimeProperty.put("required", false);
        commitTimeProperty.put("description", "Input parameter: The commit date and time for the job.");
        properties.put("commitTime", commitTimeProperty);
        Map<String, Object> jobIdProperty = new HashMap<>();
        jobIdProperty.put("type", "string");
        jobIdProperty.put("required", false);
        jobIdProperty.put("description", "Input parameter: The unique ID for an existing job. This is required if the value of <code>jobType</code> is <code>RETRY</code>.");
        properties.put("jobId", jobIdProperty);
        Map<String, Object> jobReasonProperty = new HashMap<>();
        jobReasonProperty.put("type", "string");
        jobReasonProperty.put("required", false);
        jobReasonProperty.put("description", "Input parameter: A descriptive reason for starting this job.");
        properties.put("jobReason", jobReasonProperty);
        Map<String, Object> jobTypeProperty = new HashMap<>();
        jobTypeProperty.put("type", "string");
        jobTypeProperty.put("required", true);
        jobTypeProperty.put("description", "Input parameter: Describes the type for the job. The job type <code>RELEASE</code> starts a new job with the latest change from the specified branch. This value is available only for apps that are connected to a repository. The job type <code>RETRY</code> retries an existing job. If the job type value is <code>RETRY</code>, the <code>jobId</code> is also required.");
        properties.put("jobType", jobTypeProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_apps_app_id_branches_branch_name_jobs",
            "Starts a new job for a branch of an Amplify app.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_Apps_App_Id_Branches_Branch_Name_JobsHandler(config));
    }
    
}