

# UpdateDomainAssociationRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**enableAutoSubDomain** | **Boolean** |  Enables the automated creation of subdomains for branches.  |  [optional] |
|**subDomainSettings** | [**List&lt;SubDomainSetting&gt;**](SubDomainSetting.md) |  Describes the settings for the subdomain.  |  [optional] |
|**autoSubDomainCreationPatterns** | **List&lt;String&gt;** |  Sets the branch patterns for automatic subdomain creation.  |  [optional] |
|**autoSubDomainIAMRole** | **String** |  The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.  |  [optional] |



