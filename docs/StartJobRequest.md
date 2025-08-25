

# StartJobRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**jobId** | **String** |  The unique ID for an existing job. This is required if the value of &lt;code&gt;jobType&lt;/code&gt; is &lt;code&gt;RETRY&lt;/code&gt;.  |  [optional] |
|**jobType** | [**JobTypeEnum**](#JobTypeEnum) |  Describes the type for the job. The job type &lt;code&gt;RELEASE&lt;/code&gt; starts a new job with the latest change from the specified branch. This value is available only for apps that are connected to a repository. The job type &lt;code&gt;RETRY&lt;/code&gt; retries an existing job. If the job type value is &lt;code&gt;RETRY&lt;/code&gt;, the &lt;code&gt;jobId&lt;/code&gt; is also required.  |  |
|**jobReason** | **String** |  A descriptive reason for starting this job.  |  [optional] |
|**commitId** | **String** |  The commit ID from a third-party repository provider for the job.  |  [optional] |
|**commitMessage** | **String** |  The commit message from a third-party repository provider for the job.  |  [optional] |
|**commitTime** | **OffsetDateTime** |  The commit date and time for the job.  |  [optional] |



## Enum: JobTypeEnum

| Name | Value |
|---- | -----|
| RELEASE | &quot;RELEASE&quot; |
| RETRY | &quot;RETRY&quot; |
| MANUAL | &quot;MANUAL&quot; |
| WEB_HOOK | &quot;WEB_HOOK&quot; |



