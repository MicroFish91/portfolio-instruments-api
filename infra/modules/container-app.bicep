targetScope = 'resourceGroup'

@description('Service name for Container App (API)')
param name string
@description('Location for Container App')
param location string = resourceGroup().location
@description('Tags for Container App')
param tags object = {}
@description('Key Vault resource ID for secret integration')
param keyVaultId string
@description('Azure Container Registry login server')
param acrLoginServer string
@description('Docker image name')
param imageName string
@description('Image tag')
param imageTag string = 'latest'

var resourceSuffix = take(uniqueString(subscription().id, resourceGroup().name, name), 6)
var containerAppName = '${name}-${resourceSuffix}'

resource managedIdentity 'Microsoft.ManagedIdentity/userAssignedIdentities@2023-01-31' = {
  name: '${containerAppName}-identity'
  location: location
  tags: tags
}

resource containerAppEnv 'Microsoft.App/managedEnvironments@2023-05-01' = {
  name: '${containerAppName}-env'
  location: location
  tags: tags
  properties: {
    daprEnabled: false
    zoneRedundant: false
    vnetConfiguration: {}
    diagnostics: {
      logAnalyticsWorkspaceId: '' // Provide workspace if Application Insights/logs required
    }
  }
}

resource containerApp 'Microsoft.App/containerApps@2023-05-01' = {
  name: containerAppName
  location: location
  tags: tags
  identity: {
    type: 'UserAssigned'
    userAssignedIdentities: {
      '${managedIdentity.id}': {}
    }
  }
  properties: {
    managedEnvironmentId: containerAppEnv.id
    configuration: {
      secrets: [
        {
          name: 'POSTGRES_CONNECTION_STRING'
          value: '' // Use Key Vault reference in application config
        }
      ]
      ingress: {
        external: true
        targetPort: 3000
        transport: 'Auto'
        allowInsecure: false
      }
    }
    template: {
      containers: [
        {
          name: 'api'
          image: '${acrLoginServer}/${imageName}:${imageTag}' // Dynamic registry/image/tag reference
          resources: {
            cpu: json('0.25')
            memory: '0.5Gi'
          }
        }
      ]
      scale: {
        minReplicas: 1
        maxReplicas: 3
      }
    }
  }
}

output containerAppId string = containerApp.id
output managedIdentityId string = managedIdentity.id