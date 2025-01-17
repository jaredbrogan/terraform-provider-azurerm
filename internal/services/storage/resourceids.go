// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storage

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=BlobInventoryPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/inventoryPolicies/inventoryPolicy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=EncryptionScope -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/encryptionScopes/encryptionScope1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageAccount -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageAccountDefaultBlob -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/blobServices/default
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageContainerResourceManager -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/blobServices/default/containers/container1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageQueueResourceManager -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/queueServices/default/queues/queue1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageShareResourceManager -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/fileServices/fileService1/fileshares/share1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageSyncGroup -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/storageSyncService1/syncGroups/syncGroup1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageSyncService -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/storageSyncService1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageSyncCloudEndpoint -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.StorageSync/storageSyncServices/storageSyncService1/syncGroups/syncGroup1/cloudEndpoints/cloudEndpoint1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageAccountManagementPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resGroup1/providers/Microsoft.Storage/storageAccounts/storageAccount1/managementPolicies/policy1
