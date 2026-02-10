targetScope = 'resourceGroup'

@description('Database name (Postgres)')
param name string
@description('Location for database')
param location string = resourceGroup().location
@description('Tags for database')
param tags object = {}
@description('Key Vault resource ID for secret integration')
param keyVaultId string

var resourceSuffix = take(uniqueString(subscription().id, resourceGroup().name, name), 6)
var postgresServerName = '${name}-${resourceSuffix}'

resource postgres 'Microsoft.DBforPostgreSQL/flexibleServers@2023-12-01' = {
  name: postgresServerName
  location: location
  tags: tags
  properties: {
    administratorLogin: 'azdadmin'
    administratorLoginPassword: '' // Store password in Key Vault, reference here for completeness
    version: '15'
    storage: {
      storageSizeGB: 32
    }
    backup: {
      backupRetentionDays: 7
      geoRedundantBackup: 'Disabled'
    }
    createMode: 'Default'
    highAvailability: {
      mode: 'ZoneRedundant'
    }
    network: {
      delegatedSubnetResourceId: '' // Use VNet/subnet if required
      privateEndpointConnections: []
    }
  }
  sku: {
    name: 'Standard_D2s_v3'
    tier: 'GeneralPurpose'
    family: 'Gen5'
    capacity: 2
  }
}

output postgresServerId string = postgres.id
output postgresServerName string = postgres.name
