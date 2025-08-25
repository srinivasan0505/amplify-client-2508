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

func CreateappHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateApprequest
		
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
		url := fmt.Sprintf("%s/apps", cfg.BaseURL)
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
		var result models.CreateAppResult
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

func CreateCreateappTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_apps",
		mcp.WithDescription(" Creates a new Amplify app. "),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithArray("autoBranchCreationPatterns", mcp.Description("Input parameter:  The automated branch creation glob patterns for an Amplify app. ")),
		mcp.WithString("description", mcp.Description("Input parameter:  The description for an Amplify app. ")),
		mcp.WithBoolean("enableBasicAuth", mcp.Description("Input parameter:  Enables basic authorization for an Amplify app. This will apply to all branches that are part of this app. ")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Input parameter:  The name for an Amplify app. ")),
		mcp.WithString("platform", mcp.Description("Input parameter:  The platform for the Amplify app. For a static app, set the platform type to <code>WEB</code>. For a dynamic server-side rendered (SSR) app, set the platform type to <code>WEB_COMPUTE</code>. For an app requiring Amplify Hosting's original SSR support only, set the platform type to <code>WEB_DYNAMIC</code>.")),
		mcp.WithObject("tags", mcp.Description("Input parameter:  The tag for an Amplify app. ")),
		mcp.WithObject("autoBranchCreationConfig", mcp.Description("Input parameter:  Describes the automated branch creation configuration. ")),
		mcp.WithArray("customRules", mcp.Description("Input parameter:  The custom rewrite and redirect rules for an Amplify app. ")),
		mcp.WithString("oauthToken", mcp.Description("Input parameter: <p>The OAuth token for a third-party source control system for an Amplify app. The OAuth token is used to create a webhook and a read-only deploy key using SSH cloning. The OAuth token is not stored.</p> <p>Use <code>oauthToken</code> for repository providers other than GitHub, such as Bitbucket or CodeCommit. To authorize access to GitHub as your repository provider, use <code>accessToken</code>.</p> <p>You must specify either <code>oauthToken</code> or <code>accessToken</code> when you create a new app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href=\"https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth\">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>")),
		mcp.WithObject("environmentVariables", mcp.Description("Input parameter:  The environment variables map for an Amplify app. ")),
		mcp.WithString("repository", mcp.Description("Input parameter:  The repository for an Amplify app. ")),
		mcp.WithBoolean("enableAutoBranchCreation", mcp.Description("Input parameter:  Enables automated branch creation for an Amplify app. ")),
		mcp.WithBoolean("enableBranchAutoBuild", mcp.Description("Input parameter:  Enables the auto building of branches for an Amplify app. ")),
		mcp.WithString("basicAuthCredentials", mcp.Description("Input parameter:  The credentials for basic authorization for an Amplify app. You must base64-encode the authorization credentials and provide them in the format <code>user:password</code>.")),
		mcp.WithString("customHeaders", mcp.Description("Input parameter: The custom HTTP headers for an Amplify app.")),
		mcp.WithString("buildSpec", mcp.Description("Input parameter:  The build specification (build spec) file for an Amplify app build. ")),
		mcp.WithBoolean("enableBranchAutoDeletion", mcp.Description("Input parameter:  Automatically disconnects a branch in the Amplify Console when you delete a branch from your Git repository. ")),
		mcp.WithString("iamServiceRoleArn", mcp.Description("Input parameter:  The AWS Identity and Access Management (IAM) service role for an Amplify app. ")),
		mcp.WithString("accessToken", mcp.Description("Input parameter: <p>The personal access token for a GitHub repository for an Amplify app. The personal access token is used to authorize access to a GitHub repository using the Amplify GitHub App. The token is not stored.</p> <p>Use <code>accessToken</code> for GitHub repositories only. To authorize access to a repository provider such as Bitbucket or CodeCommit, use <code>oauthToken</code>.</p> <p>You must specify either <code>accessToken</code> or <code>oauthToken</code> when you create a new app.</p> <p>Existing Amplify apps deployed from a GitHub repository using OAuth continue to work with CI/CD. However, we strongly recommend that you migrate these apps to use the GitHub App. For more information, see <a href=\"https://docs.aws.amazon.com/amplify/latest/UserGuide/setting-up-GitHub-access.html#migrating-to-github-app-auth\">Migrating an existing OAuth app to the Amplify GitHub App</a> in the <i>Amplify User Guide</i> .</p>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateappHandler(cfg),
	}
}
