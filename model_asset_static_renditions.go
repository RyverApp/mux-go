// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

type AssetStaticRenditions struct {
	// Indicates the status of downloadable MP4 versions of this asset.
	Status string                       `json:"status,omitempty"`
	Files  []AssetStaticRenditionsFiles `json:"files,omitempty"`
}
