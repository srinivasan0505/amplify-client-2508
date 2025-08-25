/**
 * Main MCP Server - Handles MCP JSON-RPC protocol
 */

import java.io.*;
import java.util.*;
import java.util.function.Function;
import java.util.concurrent.CompletableFuture;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.fasterxml.jackson.databind.node.ArrayNode;

// Import Post_Tags_Resource_ArnMCPTool - included in same package
// Import Get_Tags_Resource_ArnMCPTool - included in same package
// Import Delete_Tags_Resource_Arn_Tag_KeysMCPTool - included in same package
// Import Post_AppsMCPTool - included in same package
// Import Post_Apps_App_Id_Branches_Branch_Name_JobsMCPTool - included in same package
// Import Post_Apps_App_Id_AccesslogsMCPTool - included in same package
// Import Delete_Apps_App_Id_Backendenvironments_Environment_NameMCPTool - included in same package
// Import Get_Apps_App_Id_BackendenvironmentsMCPTool - included in same package
// Import Get_Apps_App_Id_Branches_Branch_Name_Jobs_Job_Id_ArtifactsMCPTool - included in same package
// Import Post_Apps_App_Id_WebhooksMCPTool - included in same package
// Import Delete_Apps_App_Id_Branches_Branch_Name_Jobs_Job_Id_StopMCPTool - included in same package
// Import Get_Apps_App_Id_WebhooksMCPTool - included in same package
// Import Post_Apps_App_Id_Branches_Branch_Name_DeploymentsMCPTool - included in same package
// Import Post_Apps_App_Id_Domains_Domain_NameMCPTool - included in same package
// Import Get_AppsMCPTool - included in same package
// Import Post_Apps_App_Id_Branches_Branch_NameMCPTool - included in same package
// Import Post_Apps_App_IdMCPTool - included in same package
// Import Post_Apps_App_Id_Branches_Branch_Name_Deployments_StartMCPTool - included in same package
// Import Get_Apps_App_Id_Branches_Branch_Name_JobsMCPTool - included in same package
// Import Get_Apps_App_Id_Branches_Branch_NameMCPTool - included in same package
// Import Get_Apps_App_Id_Domains_Domain_NameMCPTool - included in same package
// Import Get_Apps_App_Id_BranchesMCPTool - included in same package
// Import Get_Apps_App_IdMCPTool - included in same package
// Import Delete_Apps_App_Id_Branches_Branch_NameMCPTool - included in same package
// Import Delete_Apps_App_Id_Branches_Branch_Name_Jobs_Job_IdMCPTool - included in same package
// Import Get_Apps_App_Id_Backendenvironments_Environment_NameMCPTool - included in same package
// Import Get_Apps_App_Id_DomainsMCPTool - included in same package
// Import Post_Apps_App_Id_DomainsMCPTool - included in same package
// Import Get_Apps_App_Id_Branches_Branch_Name_Jobs_Job_IdMCPTool - included in same package
// Import Post_Apps_App_Id_BackendenvironmentsMCPTool - included in same package
// Import Delete_Apps_App_IdMCPTool - included in same package
// Import Post_Apps_App_Id_BranchesMCPTool - included in same package
// Import Delete_Apps_App_Id_Domains_Domain_NameMCPTool - included in same package
// Import Get_Artifacts_Artifact_IdMCPTool - included in same package
// Import Get_Webhooks_Webhook_IdMCPTool - included in same package
// Import Post_Webhooks_Webhook_IdMCPTool - included in same package
// Import Delete_Webhooks_Webhook_IdMCPTool - included in same package

public class MCPServer {
    
    // Common classes that all tool classes use
    public static class APIConfig {
        private final String baseUrl;
        private final String apiKey;
        
        public APIConfig(String baseUrl, String apiKey) {
            this.baseUrl = baseUrl;
            this.apiKey = apiKey;
        }
        
        public String getBaseUrl() { return baseUrl; }
        public String getApiKey() { return apiKey; }
    }
    
    public static class MCPRequest {
        private Map<String, Object> params;
        
        public Map<String, Object> getParams() { return params; }
        public void setParams(Map<String, Object> params) { this.params = params; }
        
        @SuppressWarnings("unchecked")
        public Map<String, Object> getArguments() {
            if (params != null && params.containsKey("arguments")) {
                return (Map<String, Object>) params.get("arguments");
            }
            return new HashMap<>();
        }
    }
    
    public static class MCPToolResult {
        private final String content;
        private final boolean isError;
        
        public MCPToolResult(String content, boolean isError) {
            this.content = content;
            this.isError = isError;
        }
        
        public MCPToolResult(String content) {
            this(content, false);
        }
        
        public String getContent() { return content; }
        public boolean isError() { return isError; }
    }
    
    public static class ToolDefinition {
        private final String name;
        private final String description;
        private final Map<String, Object> parameters;
        
        public ToolDefinition(String name, String description, Map<String, Object> parameters) {
            this.name = name;
            this.description = description;
            this.parameters = parameters;
        }
        
        public String getName() { return name; }
        public String getDescription() { return description; }
        public Map<String, Object> getParameters() { return parameters; }
    }
    
    public static class Tool {
        private final ToolDefinition definition;
        private final Function<MCPRequest, MCPToolResult> handler;
        
        public Tool(ToolDefinition definition, Function<MCPRequest, MCPToolResult> handler) {
            this.definition = definition;
            this.handler = handler;
        }
        
        public ToolDefinition getDefinition() { return definition; }
        public Function<MCPRequest, MCPToolResult> getHandler() { return handler; }
    }
    
    private static final ObjectMapper mapper = new ObjectMapper();
    private static List<Tool> tools = new ArrayList<>();
    private static APIConfig config;
    
    public static void main(String[] args) {
        try {
            // Initialize configuration
            String baseUrl = System.getenv("API_BASE_URL");
            String apiKey = System.getenv("API_KEY");
            
            if (baseUrl == null || apiKey == null) {
                System.err.println("Error: Please set API_BASE_URL and API_KEY environment variables");
                System.exit(1);
            }
            
            config = new APIConfig(baseUrl, apiKey);
            
            // Register all tools
            tools = Arrays.asList(
            Post_Tags_Resource_ArnMCPTool.createPost_Tags_Resource_ArnTool(config),
            Get_Tags_Resource_ArnMCPTool.createGet_Tags_Resource_ArnTool(config),
            Delete_Tags_Resource_Arn_Tag_KeysMCPTool.createDelete_Tags_Resource_Arn_Tag_KeysTool(config),
            Post_AppsMCPTool.createPost_AppsTool(config),
            Post_Apps_App_Id_Branches_Branch_Name_JobsMCPTool.createPost_Apps_App_Id_Branches_Branch_Name_JobsTool(config),
            Post_Apps_App_Id_AccesslogsMCPTool.createPost_Apps_App_Id_AccesslogsTool(config),
            Delete_Apps_App_Id_Backendenvironments_Environment_NameMCPTool.createDelete_Apps_App_Id_Backendenvironments_Environment_NameTool(config),
            Get_Apps_App_Id_BackendenvironmentsMCPTool.createGet_Apps_App_Id_BackendenvironmentsTool(config),
            Get_Apps_App_Id_Branches_Branch_Name_Jobs_Job_Id_ArtifactsMCPTool.createGet_Apps_App_Id_Branches_Branch_Name_Jobs_Job_Id_ArtifactsTool(config),
            Post_Apps_App_Id_WebhooksMCPTool.createPost_Apps_App_Id_WebhooksTool(config),
            Delete_Apps_App_Id_Branches_Branch_Name_Jobs_Job_Id_StopMCPTool.createDelete_Apps_App_Id_Branches_Branch_Name_Jobs_Job_Id_StopTool(config),
            Get_Apps_App_Id_WebhooksMCPTool.createGet_Apps_App_Id_WebhooksTool(config),
            Post_Apps_App_Id_Branches_Branch_Name_DeploymentsMCPTool.createPost_Apps_App_Id_Branches_Branch_Name_DeploymentsTool(config),
            Post_Apps_App_Id_Domains_Domain_NameMCPTool.createPost_Apps_App_Id_Domains_Domain_NameTool(config),
            Get_AppsMCPTool.createGet_AppsTool(config),
            Post_Apps_App_Id_Branches_Branch_NameMCPTool.createPost_Apps_App_Id_Branches_Branch_NameTool(config),
            Post_Apps_App_IdMCPTool.createPost_Apps_App_IdTool(config),
            Post_Apps_App_Id_Branches_Branch_Name_Deployments_StartMCPTool.createPost_Apps_App_Id_Branches_Branch_Name_Deployments_StartTool(config),
            Get_Apps_App_Id_Branches_Branch_Name_JobsMCPTool.createGet_Apps_App_Id_Branches_Branch_Name_JobsTool(config),
            Get_Apps_App_Id_Branches_Branch_NameMCPTool.createGet_Apps_App_Id_Branches_Branch_NameTool(config),
            Get_Apps_App_Id_Domains_Domain_NameMCPTool.createGet_Apps_App_Id_Domains_Domain_NameTool(config),
            Get_Apps_App_Id_BranchesMCPTool.createGet_Apps_App_Id_BranchesTool(config),
            Get_Apps_App_IdMCPTool.createGet_Apps_App_IdTool(config),
            Delete_Apps_App_Id_Branches_Branch_NameMCPTool.createDelete_Apps_App_Id_Branches_Branch_NameTool(config),
            Delete_Apps_App_Id_Branches_Branch_Name_Jobs_Job_IdMCPTool.createDelete_Apps_App_Id_Branches_Branch_Name_Jobs_Job_IdTool(config),
            Get_Apps_App_Id_Backendenvironments_Environment_NameMCPTool.createGet_Apps_App_Id_Backendenvironments_Environment_NameTool(config),
            Get_Apps_App_Id_DomainsMCPTool.createGet_Apps_App_Id_DomainsTool(config),
            Post_Apps_App_Id_DomainsMCPTool.createPost_Apps_App_Id_DomainsTool(config),
            Get_Apps_App_Id_Branches_Branch_Name_Jobs_Job_IdMCPTool.createGet_Apps_App_Id_Branches_Branch_Name_Jobs_Job_IdTool(config),
            Post_Apps_App_Id_BackendenvironmentsMCPTool.createPost_Apps_App_Id_BackendenvironmentsTool(config),
            Delete_Apps_App_IdMCPTool.createDelete_Apps_App_IdTool(config),
            Post_Apps_App_Id_BranchesMCPTool.createPost_Apps_App_Id_BranchesTool(config),
            Delete_Apps_App_Id_Domains_Domain_NameMCPTool.createDelete_Apps_App_Id_Domains_Domain_NameTool(config),
            Get_Artifacts_Artifact_IdMCPTool.createGet_Artifacts_Artifact_IdTool(config),
            Get_Webhooks_Webhook_IdMCPTool.createGet_Webhooks_Webhook_IdTool(config),
            Post_Webhooks_Webhook_IdMCPTool.createPost_Webhooks_Webhook_IdTool(config),
            Delete_Webhooks_Webhook_IdMCPTool.createDelete_Webhooks_Webhook_IdTool(config)
            );
            
            System.err.println("MCP Server started with " + tools.size() + " tools");
            
            // Start JSON-RPC server
            runServer();
            
        } catch (Exception e) {
            System.err.println("Failed to start MCP server: " + e.getMessage());
            e.printStackTrace();
            System.exit(1);
        }
    }
    
    private static void runServer() throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String line;
        
        while ((line = reader.readLine()) != null) {
            JsonNode request = null;
            try {
                request = mapper.readTree(line);
                JsonNode response = handleRequest(request);
                
                if (response != null) {
                    System.out.println(mapper.writeValueAsString(response));
                }
                
            } catch (Exception e) {
                // Send error response
                ObjectNode errorResponse = mapper.createObjectNode();
                errorResponse.put("jsonrpc", "2.0");
                if (request != null && request.has("id")) {
                    errorResponse.put("id", request.get("id").asText());
                } else {
                    errorResponse.putNull("id");
                }
                
                ObjectNode error = mapper.createObjectNode();
                error.put("code", -32603);
                error.put("message", "Internal error: " + e.getMessage());
                errorResponse.set("error", error);
                
                System.out.println(mapper.writeValueAsString(errorResponse));
            }
        }
    }
    
    private static JsonNode handleRequest(JsonNode request) {
        if (!request.has("method")) {
            return createErrorResponse(request, -32600, "Invalid Request");
        }
        
        String method = request.get("method").asText();
        JsonNode params = request.get("params");
        String id = request.has("id") ? request.get("id").asText() : null;
        
        switch (method) {
            case "initialize":
                return handleInitialize(id);
            case "tools/list":
                return handleToolsList(id);
            case "tools/call":
                return handleToolsCall(id, params);
            default:
                return createErrorResponse(request, -32601, "Method not found");
        }
    }
    
    private static JsonNode handleInitialize(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        result.put("protocolVersion", "2024-11-05");
        
        ObjectNode capabilities = mapper.createObjectNode();
        ObjectNode toolsCapability = mapper.createObjectNode();
        toolsCapability.put("listChanged", false);
        capabilities.set("tools", toolsCapability);
        result.set("capabilities", capabilities);
        
        ObjectNode serverInfo = mapper.createObjectNode();
        serverInfo.put("name", "Opsera MCP Server");
        serverInfo.put("version", "1.0.0");
        result.set("serverInfo", serverInfo);
        
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsList(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        ArrayNode toolsArray = mapper.createArrayNode();
        
        for (Tool tool : tools) {
            ObjectNode toolDef = mapper.createObjectNode();
            toolDef.put("name", tool.getDefinition().getName());
            toolDef.put("description", tool.getDefinition().getDescription());
            
            // Convert parameters to JSON
            ObjectNode inputSchema = mapper.valueToTree(tool.getDefinition().getParameters());
            toolDef.set("inputSchema", inputSchema);
            
            toolsArray.add(toolDef);
        }
        
        result.set("tools", toolsArray);
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsCall(String id, JsonNode params) {
        if (!params.has("name")) {
            return createErrorResponse(null, -32602, "Invalid params: missing 'name'");
        }
        
        String toolName = params.get("name").asText();
        JsonNode arguments = params.has("arguments") ? params.get("arguments") : mapper.createObjectNode();
        
        // Find the tool
        Tool tool = null;
        for (Tool t : tools) {
            if (t.getDefinition().getName().equals(toolName)) {
                tool = t;
                break;
            }
        }
        
        if (tool == null) {
            return createErrorResponse(null, -32602, "Tool not found: " + toolName);
        }
        
        try {
            // Convert arguments to Map
            Map<String, Object> argsMap = mapper.convertValue(arguments, Map.class);
            
            // Create MCP request
            MCPRequest mcpRequest = new MCPRequest();
            Map<String, Object> requestParams = new HashMap<>();
            requestParams.put("arguments", argsMap);
            mcpRequest.setParams(requestParams);
            
            // Execute tool
            MCPToolResult result = tool.getHandler().apply(mcpRequest);
            
            // Create response
            ObjectNode response = mapper.createObjectNode();
            response.put("jsonrpc", "2.0");
            response.put("id", id);
            
            ObjectNode resultObj = mapper.createObjectNode();
            ArrayNode content = mapper.createArrayNode();
            
            ObjectNode textContent = mapper.createObjectNode();
            textContent.put("type", "text");
            textContent.put("text", result.getContent());
            content.add(textContent);
            
            resultObj.set("content", content);
            resultObj.put("isError", result.isError());
            
            response.set("result", resultObj);
            return response;
            
        } catch (Exception e) {
            return createErrorResponse(null, -32603, "Tool execution failed: " + e.getMessage());
        }
    }
    
    private static JsonNode createErrorResponse(JsonNode request, int code, String message) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", request != null && request.has("id") ? request.get("id").asText() : null);
        
        ObjectNode error = mapper.createObjectNode();
        error.put("code", code);
        error.put("message", message);
        response.set("error", error);
        
        return response;
    }
}