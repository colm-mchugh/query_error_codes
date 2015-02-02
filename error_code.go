//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

/*

Package errors provides user-visible errors and warnings. These errors
include error codes and will eventually provide multi-language
messages.

*/
package errors

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

// service level errors - errors that are created in the service package
const (
	// A readonly error is returned when
	SERVICE_READONLY_CODE    = 1000
	SERVICE_READONLY_MESSAGE = "The server or request is read-only and cannot accept this write statement."

	// This error is returned if the REST API receives a non-GET or POST request
	SERVICE_UNSUPPORTED_HTTP_CODE    = 1010
	SERVICE_UNSUPPORTED_HTTP_MESSAGE = "Unsupported http method: %s"

	// Returned for not yet implemented features
	// Example: ZIP compression not yet implemented
	SERVICE_NOT_YET_IMPLEMENTED_CODE    = 1020
	SERVICE_NOT_YET_IMPLEMENTED_MESSAGE = "%s %s not yet implemented"

	// Returned if the value the value for a parameter is incorrect
	// Example: Unknown scan_consistency value: foo
	SERVICE_UNRECOGNIZED_VALUE_CODE    = 1030
	SERVICE_UNRECOGNIZED_VALUE_MESSAGE = "Unknown %s value: %s"

	// Returned if there is an error
	SERVICE_INTERNAL_ERROR_CODE    = 1040
	SERVICE_INTERNAL_ERROR_MESSAGE = "Error processing %s"

	// Returned if the request is missing a requried parameter
	SERVICE_MISSING_REQUIRED_CODE    = 1050
	SERVICE_MISSING_REQUIRED_MESSAGE = "No %s value"

	// Returned if a request includes multiple values of a parameter
	SERVICE_MULTIPLE_VALUES_CODE    = 1060
	SERVICE_MULTIPLE_VALUES_MESSAGE = "Multiple values for %s."

	// Returned if the value for a parameter is of incorrect type
	SERVICE_TYPE_MISMATCH_CODE    = 1070
	SERVICE_TYPE_MISMATCH_MESSAGE = "%s has to be of type %s"

	// Returned if results contain an invalid json object
	SERVICE_INVALID_JSON_CODE    = 1100
	SERVICE_INVALID_JSON_MESSAGE = "Invalid JSON in results"
)

// Parse errors - errors that are created in the parse package
func NewParseSyntaxError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 3000, IKey: "parse.syntax_error", ICause: e,
		InternalMsg: msg, InternalCaller: CallerN(1)}
}

// Plan errors - errors that are created in the plan package
func NewPlanError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 4000, IKey: "plan_error", ICause: e, InternalMsg: msg, InternalCaller: CallerN(1)}
}

// admin level errors - errors that are created in the clustering and accounting packages

func NewAdminConnectionError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2000, IKey: "admin.clustering.connection_error", ICause: e,
		InternalMsg: "Error connecting to " + msg, InternalCaller: CallerN(1)}
}

func NewAdminInvalidURL(component string, url string) Error {
	return &err{level: EXCEPTION, ICode: 2010, IKey: "admin.invalid_url",
		InternalMsg: fmt.Sprintf("Invalid % url: %s", component, url), InternalCaller: CallerN(1)}
}

func NewAdminDecodingError(e error) Error {
	return &err{level: EXCEPTION, ICode: 2020, IKey: "admin.json_decoding_error", ICause: e,
		InternalMsg: "Error in JSON decoding", InternalCaller: CallerN(1)}
}

func NewAdminEncodingError(e error) Error {
	return &err{level: EXCEPTION, ICode: 2030, IKey: "admin.json_encoding_error", ICause: e,
		InternalMsg: "Error in JSON encoding", InternalCaller: CallerN(1)}
}

func NewAdminGetClusterError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2040, IKey: "admin.clustering.get_cluster_error", ICause: e,
		InternalMsg: "Error retrieving cluster " + msg, InternalCaller: CallerN(1)}
}

func NewAdminAddClusterError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2050, IKey: "admin.clustering.add_cluster_error", ICause: e,
		InternalMsg: "Error adding cluster " + msg, InternalCaller: CallerN(1)}
}

func NewAdminRemoveClusterError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2060, IKey: "admin.clustering.remove_cluster_error", ICause: e,
		InternalMsg: "Error removing cluster " + msg, InternalCaller: CallerN(1)}
}

func NewAdminGetNodeError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2070, IKey: "admin.clustering.get_node_error", ICause: e,
		InternalMsg: "Error retrieving node " + msg, InternalCaller: CallerN(1)}
}

func NewAdminNoNodeError(msg string) Error {
	return &err{level: EXCEPTION, ICode: 2080, IKey: "admin.clustering.no_such_node",
		InternalMsg: "No such  node " + msg, InternalCaller: CallerN(1)}
}

func NewAdminAddNodeError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2090, IKey: "admin.clustering.add_node_error", ICause: e,
		InternalMsg: "Error adding node " + msg, InternalCaller: CallerN(1)}
}

func NewAdminRemoveNodeError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2100, IKey: "admin.clustering.remove_node_error", ICause: e,
		InternalMsg: "Error removing node " + msg, InternalCaller: CallerN(1)}
}

func NewAdminMakeMetricError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 2110, IKey: "admin.accounting.metric.create", ICause: e,
		InternalMsg: "Error creating metric " + msg, InternalCaller: CallerN(1)}
}

// Authorization Errors
func NewDatastoreAuthorizationError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 10000, IKey: "datastore.couchbase.authorization_error", ICause: e,
		InternalMsg: "Authorization Failed " + msg, InternalCaller: CallerN(1)}
}

// System datastore error codes
func NewSystemDatastoreError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11000, IKey: "datastore.system.generic_error", ICause: e,
		InternalMsg: "System datastore error " + msg, InternalCaller: CallerN(1)}

}

func NewSystemNamespaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11001, IKey: "datastore.system.namespace_not_found", ICause: e,
		InternalMsg: "Datastore : namespace not found " + msg, InternalCaller: CallerN(1)}

}

func NewSystemKeyspaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11002, IKey: "datastore.system.keyspace_not_found", ICause: e,
		InternalMsg: "Keyspace not found " + msg, InternalCaller: CallerN(1)}

}

func NewSystemNotImplementedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11003, IKey: "datastore.system.not_implemented", ICause: e,
		InternalMsg: "System datastore :  Not implemented " + msg, InternalCaller: CallerN(1)}

}

func NewSystemNotSupportedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11004, IKey: "datastore.system.not_supported", ICause: e,
		InternalMsg: "System datastore : Not supported " + msg, InternalCaller: CallerN(1)}

}

func NewSystemIdxNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11005, IKey: "datastore.system.idx_not_found", ICause: e,
		InternalMsg: "System datastore : Index not found " + msg, InternalCaller: CallerN(1)}

}

func NewSystemIdxNoDropError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 11006, IKey: "datastore.system.idx_no_drop", ICause: e,
		InternalMsg: "System datastore : This  index cannot be dropped " + msg, InternalCaller: CallerN(1)}

}

// Datastore/couchbase error codes
func NewCbConnectionError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12000, IKey: "datastore.couchbase.connection_error", ICause: e,
		InternalMsg: "Cannot connect " + msg, InternalCaller: CallerN(1)}

}

func NewCbUrlParseError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12001, IKey: "datastore.couchbase.url_parse", ICause: e,
		InternalMsg: "Cannot parse url " + msg, InternalCaller: CallerN(1)}
}

func NewCbNamespaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12002, IKey: "datastore.couchbase.namespace_not_found", ICause: e,
		InternalMsg: "Namespace not found " + msg, InternalCaller: CallerN(1)}
}

func NewCbKeyspaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12003, IKey: "datastore.couchbase.keyspace_not_found", ICause: e,
		InternalMsg: "Keyspace not found " + msg, InternalCaller: CallerN(1)}
}

func NewCbPrimaryIndexNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12004, IKey: "datastore.couchbase.primary_idx_not_found", ICause: e,
		InternalMsg: "Primary Index not found " + msg, InternalCaller: CallerN(1)}
}

func NewCbIndexerNotImplementedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12005, IKey: "datastore.couchbase.indexer_not_implemented", ICause: e,
		InternalMsg: "Indexer not implemented " + msg, InternalCaller: CallerN(1)}
}

func NewCbKeyspaceCountError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12006, IKey: "datastore.couchbase.keyspace_count_error", ICause: e,
		InternalMsg: "Failed to get keyspace count " + msg, InternalCaller: CallerN(1)}
}

func NewCbNoKeysFetchError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12007, IKey: "datastore.couchbase.no_keys_fetch", ICause: e,
		InternalMsg: "No keys to fetch " + msg, InternalCaller: CallerN(1)}
}

func NewCbBulkGetError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12008, IKey: "datastore.couchbase.bulk_get_error", ICause: e,
		InternalMsg: "Error performing buck get " + msg, InternalCaller: CallerN(1)}
}

func NewCbDMLError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12009, IKey: "datastore.couchbase.DML_error", ICause: e,
		InternalMsg: "DML Error" + msg, InternalCaller: CallerN(1)}
}

func NewCbNoKeysInsertError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12010, IKey: "datastore.couchbase.no_keys_insert", ICause: e,
		InternalMsg: "No keys to insert " + msg, InternalCaller: CallerN(1)}
}

func NewCbDeleteFailedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12011, IKey: "datastore.couchbase.delete_failed", ICause: e,
		InternalMsg: msg, InternalCaller: CallerN(1)}
}

func NewCbLoadIndexesError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 12012, IKey: "datastore.couchbase.load_index_failed", ICause: e,
		InternalMsg: "Failed to load indexes" + msg, InternalCaller: CallerN(1)}
}

// Datastore/couchbase/view index error codes
func NewCbViewCreateError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13000, IKey: "datastore.couchbase.view.create_failed", ICause: e,
		InternalMsg: "Failed to create view" + msg, InternalCaller: CallerN(1)}
}

func NewCbViewNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13001, IKey: "datastore.couchbase.view.not_found", ICause: e,
		InternalMsg: "View Index not found " + msg, InternalCaller: CallerN(1)}
}

func NewCbViewExistsError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13003, IKey: "datastore.couchbase.view.exists", ICause: e,
		InternalMsg: "View index exists" + msg, InternalCaller: CallerN(1)}
}

func NewCbViewsWithNotAllowedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13004, IKey: "datastore.couchbase.view.with_not_allowed", ICause: e,
		InternalMsg: "Views not allowed for WITH keyword" + msg, InternalCaller: CallerN(1)}
}

func NewCbViewsNotSupportedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13005, IKey: "datastore.couchbase.view.not_supported", ICause: e,
		InternalMsg: "View indexes not supported" + msg, InternalCaller: CallerN(1)}
}

func NewCbViewsDropIndexError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13006, IKey: "datastore.couchbase.view.drop_index_error", ICause: e,
		InternalMsg: "Failed to drop index " + msg, InternalCaller: CallerN(1)}
}

func NewCbViewsAccessError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 13007, IKey: "datastore.couchbase.view.access_error", ICause: e,
		InternalMsg: "Failed to access view " + msg, InternalCaller: CallerN(1)}
}

// Datastore File based error codes

func NewFileDatastoreError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15000, IKey: "datastore.file.generic_file_error", ICause: e,
		InternalMsg: "Error in file datastore " + msg, InternalCaller: CallerN(1)}
}

func NewFileNamespaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15001, IKey: "datastore.file.namespace_not_found", ICause: e,
		InternalMsg: "Namespace not found " + msg, InternalCaller: CallerN(1)}
}

func NewFileKeyspaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15002, IKey: "datastore.file.keyspace_not_found", ICause: e,
		InternalMsg: "Keyspace not found " + msg, InternalCaller: CallerN(1)}
}

func NewFileDuplicateNamespaceError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15003, IKey: "datastore.file.duplicate_namespace", ICause: e,
		InternalMsg: "Duplicate Namespace " + msg, InternalCaller: CallerN(1)}
}

func NewFileDuplicateKeyspaceError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15004, IKey: "datastore.file.duplicate_keyspace", ICause: e,
		InternalMsg: "Duplicate Keyspace " + msg, InternalCaller: CallerN(1)}
}

func NewFileNoKeysInsertError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15005, IKey: "datastore.file.no_keys_insert", ICause: e,
		InternalMsg: "No keys to insert " + msg, InternalCaller: CallerN(1)}
}

func NewFileKeyExists(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15006, IKey: "datastore.file.key_exists", ICause: e,
		InternalMsg: "Key Exists " + msg, InternalCaller: CallerN(1)}
}

func NewFileDMLError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15007, IKey: "datastore.file.DML_error", ICause: e,
		InternalMsg: "DML Error " + msg, InternalCaller: CallerN(1)}
}

func NewFileKeyspaceNotDirError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15008, IKey: "datastore.file.keyspacenot_dir", ICause: e,
		InternalMsg: "Keyspace path must be a directory " + msg, InternalCaller: CallerN(1)}
}

func NewFileIdxNotFound(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15009, IKey: "datastore.file.idx_not_found", ICause: e,
		InternalMsg: "Index not found " + msg, InternalCaller: CallerN(1)}
}

func NewFileNotSupported(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15010, IKey: "datastore.file.not_supported", ICause: e,
		InternalMsg: "Operation not supported " + msg, InternalCaller: CallerN(1)}
}

func NewFilePrimaryIdxNoDropError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 15011, IKey: "datastore.file.primary_idx_no_drop", ICause: e,
		InternalMsg: "Primary Index cannot be dropped " + msg, InternalCaller: CallerN(1)}
}

// Error codes for all other datastores, e.g Mock
func NewOtherDatastoreError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16000, IKey: "datastore.other.datastore_generic_error", ICause: e,
		InternalMsg: "Error in datastore " + msg, InternalCaller: CallerN(1)}
}

func NewOtherNamespaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16001, IKey: "datastore.other.namespace_not_found", ICause: e,
		InternalMsg: "Namespace Not Found " + msg, InternalCaller: CallerN(1)}
}

func NewOtherKeyspaceNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16002, IKey: "datastore.other.keyspace_not_found", ICause: e,
		InternalMsg: "Keyspace Not Found " + msg, InternalCaller: CallerN(1)}
}

func NewOtherNotImplementedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16003, IKey: "datastore.other.not_implemented", ICause: e,
		InternalMsg: "Not Implemented " + msg, InternalCaller: CallerN(1)}
}

func NewOtherIdxNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16004, IKey: "datastore.other.idx_not_found", ICause: e,
		InternalMsg: "Index not found  " + msg, InternalCaller: CallerN(1)}
}

func NewOtherIdxNoDrop(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16005, IKey: "datastore.other.idx_no_drop", ICause: e,
		InternalMsg: "Index Cannot be dropped " + msg, InternalCaller: CallerN(1)}
}

func NewOtherNotSupportedError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16006, IKey: "datastore.other.not_supported", ICause: e,
		InternalMsg: "Not supported for this datastore " + msg, InternalCaller: CallerN(1)}
}

func NewOtherKeyNotFoundError(e error, msg string) Error {
	return &err{level: EXCEPTION, ICode: 16007, IKey: "datastore.other.key_not_found", ICause: e,
		InternalMsg: "Key not found " + msg, InternalCaller: CallerN(1)}
}

// Returns "FileName:LineNum" of caller.
func Caller() string {
	return CallerN(1)
}

// Returns "FileName:LineNum" of the Nth caller on the call stack,
// where level of 0 is the caller of CallerN.
func CallerN(level int) string {
	_, fname, lineno, ok := runtime.Caller(1 + level)
	if !ok {
		return "unknown:0"
	}
	return fmt.Sprintf("%s:%d",
		strings.Split(path.Base(fname), ".")[0], lineno)
}
