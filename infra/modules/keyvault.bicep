targetScope = 'resourceGroup'

@description('Base name for Key Vault resource')
param name string
@description('Location for Key Vault')
param location string = resourceGroup().location
@description('Tags to apply to Key Vault')
param tags object = {}

var resourceSuffix = take(uniqueString(subscription().id, resourceGroup().name, name), 6)
var keyVaultName = '${name}-${resourceSuffix}'

resource keyVault 'Microsoft.KeyVault/vaults@2023-02-01' = {
  name: keyVaultName
  location: location
  tags: tags
  properties: {
    tenantId: subscription().tenantId
    sku: {
      family: 'A'
      name: 'standard'
    }
    accessPolicies: []
    enabledForDeployment: true
    enabledForTemplateDeployment: true
    enableRbacAuthorization: true
    enableSoftDelete: true
    enablePurgeProtection: true
    publicNetworkAccess: 'Enabled'
  }
}

output keyVaultId string = keyVault.id
