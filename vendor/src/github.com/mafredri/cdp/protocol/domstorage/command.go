// Code generated by cdpgen. DO NOT EDIT.

package domstorage

// ClearArgs represents the arguments for Clear in the DOMStorage domain.
type ClearArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
}

// NewClearArgs initializes ClearArgs with the required arguments.
func NewClearArgs(storageID StorageID) *ClearArgs {
	args := new(ClearArgs)
	args.StorageID = storageID
	return args
}

// GetDOMStorageItemsArgs represents the arguments for GetDOMStorageItems in the DOMStorage domain.
type GetDOMStorageItemsArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
}

// NewGetDOMStorageItemsArgs initializes GetDOMStorageItemsArgs with the required arguments.
func NewGetDOMStorageItemsArgs(storageID StorageID) *GetDOMStorageItemsArgs {
	args := new(GetDOMStorageItemsArgs)
	args.StorageID = storageID
	return args
}

// GetDOMStorageItemsReply represents the return values for GetDOMStorageItems in the DOMStorage domain.
type GetDOMStorageItemsReply struct {
	Entries []Item `json:"entries"` // No description.
}

// RemoveDOMStorageItemArgs represents the arguments for RemoveDOMStorageItem in the DOMStorage domain.
type RemoveDOMStorageItemArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
}

// NewRemoveDOMStorageItemArgs initializes RemoveDOMStorageItemArgs with the required arguments.
func NewRemoveDOMStorageItemArgs(storageID StorageID, key string) *RemoveDOMStorageItemArgs {
	args := new(RemoveDOMStorageItemArgs)
	args.StorageID = storageID
	args.Key = key
	return args
}

// SetDOMStorageItemArgs represents the arguments for SetDOMStorageItem in the DOMStorage domain.
type SetDOMStorageItemArgs struct {
	StorageID StorageID `json:"storageId"` // No description.
	Key       string    `json:"key"`       // No description.
	Value     string    `json:"value"`     // No description.
}

// NewSetDOMStorageItemArgs initializes SetDOMStorageItemArgs with the required arguments.
func NewSetDOMStorageItemArgs(storageID StorageID, key string, value string) *SetDOMStorageItemArgs {
	args := new(SetDOMStorageItemArgs)
	args.StorageID = storageID
	args.Key = key
	args.Value = value
	return args
}
