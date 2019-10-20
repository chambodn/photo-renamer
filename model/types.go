package model

import (
	"encoding/json"
)

// APIError is the base type for endpoint-specific errors.
type APIError struct {
	ErrorSummary string `json:"error_summary"`
}

// Tagged is used for tagged unions.
type Tagged struct {
	Tag string `json:".tag"`
}

// GetMetadataResponse is the struct representation for the json answer to the /get_metadata directive
// https://www.dropbox.com/developers/documentation/http/documentation#files-get_metadata
type GetMetadataResponse struct {
	Tag         string `json:".tag"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	PathLower   string `json:"path_lower"`
	PathDisplay string `json:"path_display"`
	SharingInfo struct {
		ReadOnly             bool   `json:"read_only"`
		ParentSharedFolderID string `json:"parent_shared_folder_id"`
		TraverseOnly         bool   `json:"traverse_only"`
		NoAccess             bool   `json:"no_access"`
	} `json:"sharing_info"`
	PropertyGroups []struct {
		TemplateID string `json:"template_id"`
		Fields     []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"fields"`
	} `json:"property_groups"`
}

// GetMetadataArg is the struct representation for the json answer to the /get_metadata directive
// https://www.dropbox.com/developers/documentation/http/documentation#files-get_metadata
type GetMetadataArg struct {
	Path                            string `json:"path"`
	IncludeMediaInfo                bool   `json:"include_media_info"`
	IncludeDeleted                  bool   `json:"include_deleted"`
	IncludeHasExplicitSharedMembers bool   `json:"include_has_explicit_shared_members"`
}

// GetMetadataError : has no documentation (yet)
type GetMetadataError struct {
	Tagged
	// Path : has no documentation (yet)
	Path *LookupError `json:"path,omitempty"`
}

// Valid tag values for GetMetadataError
const (
	GetMetadataErrorPath = "path"
)

// UnmarshalJSON deserializes into a GetMetadataError instance
func (u *GetMetadataError) UnmarshalJSON(body []byte) error {
	type wrap struct {
		Tagged
		// Path : has no documentation (yet)
		Path json.RawMessage `json:"path,omitempty"`
	}
	var w wrap
	var err error
	if err = json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "path":
		err = json.Unmarshal(w.Path, &u.Path)

		if err != nil {
			return err
		}
	}
	return nil
}

// LookupError : has no documentation (yet)
type LookupError struct {
	Tagged
	// MalformedPath : The given path does not satisfy the required path format.
	// Please refer to the `Path formats documentation`
	// <https://www.dropbox.com/developers/documentation/http/documentation#path-formats>
	// for more information.
	MalformedPath string `json:"malformed_path,omitempty"`
}

// Valid tag values for LookupError
const (
	LookupErrorMalformedPath     = "malformed_path"
	LookupErrorNotFound          = "not_found"
	LookupErrorNotFile           = "not_file"
	LookupErrorNotFolder         = "not_folder"
	LookupErrorRestrictedContent = "restricted_content"
	LookupErrorOther             = "other"
)

// UnmarshalJSON deserializes into a LookupError instance
func (u *LookupError) UnmarshalJSON(body []byte) error {
	type wrap struct {
		Tagged
		// MalformedPath : The given path does not satisfy the required path
		// format. Please refer to the `Path formats documentation`
		// <https://www.dropbox.com/developers/documentation/http/documentation#path-formats>
		// for more information.
		MalformedPath json.RawMessage `json:"malformed_path,omitempty"`
	}
	var w wrap
	var err error
	if err = json.Unmarshal(body, &w); err != nil {
		return err
	}
	u.Tag = w.Tag
	switch u.Tag {
	case "malformed_path":
		err = json.Unmarshal(body, &u.MalformedPath)

		if err != nil {
			return err
		}
	}
	return nil
}
