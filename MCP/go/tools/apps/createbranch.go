package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/aws-amplify/mcp-server/config"
	"github.com/aws-amplify/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatebranchHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		appIdVal, ok := args["appId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: appId"), nil
		}
		appId, ok := appIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: appId"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateBranchrequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/apps/%s/branches", cfg.BaseURL, appId)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Content-Sha256"]; ok {
			req.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Date"]; ok {
			req.Header.Set("X-Amz-Date", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Algorithm"]; ok {
			req.Header.Set("X-Amz-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Credential"]; ok {
			req.Header.Set("X-Amz-Credential", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Security-Token"]; ok {
			req.Header.Set("X-Amz-Security-Token", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Signature"]; ok {
			req.Header.Set("X-Amz-Signature", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-SignedHeaders"]; ok {
			req.Header.Set("X-Amz-SignedHeaders", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateBranchResult
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatebranchTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps_appId_branches",
		mcp.WithDescription(" Creates a new branch for an Amplify app. "),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("appId", mcp.Required(), mcp.Description(" The unique ID for an Amplify app. ")),
		mcp.WithBoolean("enableAutoBuild", mcp.Description("Input parameter:  Enables auto building for the branch. ")),
		mcp.WithBoolean("enableNotification", mcp.Description("Input parameter:  Enables notifications for the branch. ")),
		mcp.WithString("buildSpec", mcp.Description("Input parameter:  The build specification (build spec) file for an Amplify app build. ")),
		mcp.WithString("basicAuthCredentials", mcp.Description("Input parameter:  The basic authorization credentials for the branch. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.")),
		mcp.WithBoolean("enableBasicAuth", mcp.Description("Input parameter:  Enables basic authorization for the branch. ")),
		mcp.WithString("ttl", mcp.Description("Input parameter:  The content Time to Live (TTL) for the website in seconds. ")),
		mcp.WithObject("environmentVariables", mcp.Description("Input parameter:  The environment variables for the branch. ")),
		mcp.WithString("displayName", mcp.Description("Input parameter:  The display name for a branch. This is used as the default domain prefix. ")),
		mcp.WithString("backendEnvironmentArn", mcp.Description("Input parameter:  The Amazon Resource Name (ARN) for a backend environment that is part of an Amplify app. ")),
		mcp.WithString("framework", mcp.Description("Input parameter:  The framework for the branch. ")),
		mcp.WithBoolean("enablePullRequestPreview", mcp.Description("Input parameter:  Enables pull request previews for this branch. ")),
		mcp.WithString("description", mcp.Description("Input parameter:  The description for the branch. ")),
		mcp.WithBoolean("enablePerformanceMode", mcp.Description("Input parameter: <p>Enables performance mode for the branch.</p> <p>Performance mode optimizes for faster hosting performance by keeping content cached at the edge for a longer interval. When performance mode is enabled, hosting configuration or code changes can take up to 10 minutes to roll out. </p>")),
		mcp.WithString("stage", mcp.Description("Input parameter:  Describes the current stage for the branch. ")),
		mcp.WithString("branchName", mcp.Required(), mcp.Description("Input parameter:  The name for the branch. ")),
		mcp.WithString("pullRequestEnvironmentName", mcp.Description("Input parameter:  The Amplify environment name for the pull request. ")),
		mcp.WithObject("tags", mcp.Description("Input parameter:  The tag for the branch. ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatebranchHandler(cfg),
	}
}
