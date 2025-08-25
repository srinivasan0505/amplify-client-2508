

# CreateDomainAssociationRequest


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**domainName** | **String** |  The domain name for the domain association.  |  |
|**enableAutoSubDomain** | **Boolean** |  Enables the automated creation of subdomains for branches.  |  [optional] |
|**subDomainSettings** | [**List&lt;SubDomainSetting&gt;**](SubDomainSetting.md) |  The setting for the subdomain.  |  |
|**autoSubDomainCreationPatterns** | **List&lt;String&gt;** |  Sets the branch patterns for automatic subdomain creation.  |  [optional] |
|**autoSubDomainIAMRole** | **String** |  The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.  |  [optional] |



