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


const (
// SERVICE errors are returned by the service API if a request cannot be fulfilled
// by the query engine.
//
// If a request includes a query that produces a mutation to a server that 
// is read-only, this error is returned
SERVICE_READONLY_CODE    = 1000
SERVICE_READONLY_MESSAGE = "The server or request is read-only and " +
		"cannot accept this write statement."
)

const (
// If a query is submitted using an HTTP method not supported by the query engine
// then this error is returned. 
SERVICE_UNSUPPORTED_HTTP_CODE    = 1010
SERVICE_UNSUPPORTED_HTTP_MESSAGE = "Unsupported http method: %s"
)

const (
// For a request that includes a parameter value that is not yet implemented. 
// Example: ZIP compression not yet implemented
SERVICE_NOT_YET_IMPLEMENTED_CODE    = 1020
SERVICE_NOT_YET_IMPLEMENTED_MESSAGE = "%s %s not yet implemented"
)

const (
// For a request that includes a parameter value that is not recognized by the query engine.
// Example: Unknown scan_consistency value: foo
SERVICE_UNRECOGNIZED_VALUE_CODE    = 1030
SERVICE_UNRECOGNIZED_VALUE_MESSAGE = "Unknown %s value: %s"
)

const (
// Returned if there is an internal error while processing a request.
SERVICE_INTERNAL_ERROR_CODE    = 1040
SERVICE_INTERNAL_ERROR_MESSAGE = "Error processing %s"
)

const (
// Returned if a request is missing a required parameter.
// Example: No statement or prepared value
SERVICE_MISSING_REQUIRED_CODE    = 1050
SERVICE_MISSING_REQUIRED_MESSAGE = "No %s value"
)

const (
// Returned if a request includes multiple values of a parameter
SERVICE_MULTIPLE_VALUES_CODE    = 1060
SERVICE_MULTIPLE_VALUES_MESSAGE = "Multiple values for %s."
)

const (
// Returned if the value for a parameter is of incorrect type
// Example: pretty has to be of type boolean
SERVICE_TYPE_MISMATCH_CODE    = 1070
SERVICE_TYPE_MISMATCH_MESSAGE = "%s has to be of type %s"
)

const (
// Returned if results contain an invalid json object
SERVICE_INVALID_JSON_CODE    = 1100
SERVICE_INVALID_JSON_MESSAGE = "Invalid JSON in results"
)


const (
// admin level errors - errors that are created in the clustering
// and accounting packages; these errors can be returned by the
// clustering API

// Comments
ADMIN_CONNECTION_ERROR_CODE    = 2000
ADMIN_CONNECTION_ERROR_MESSAGE = "Error connecting to %S"

// Put documentation here
ADMIN_INVALID_URL_CODE    = 2010
ADMIN_INVALID_URL_MESSAGE = "Invalid % url: %s"

// Great documentation
ADMIN_JSON_DECODING_ERROR_CODE    = 2020
ADMIN_JSON_DECODING_ERROR_MESSAGE = "Error in JSON decoding: %s"

ADMIN_JSON_ENCODING_ERROR_CODE    = 2030
ADMIN_JSON_ENCODING_ERROR_MESSAGE = "Error in JSON encoding: %s"

ADMIN_RETRIEVE_CLUSTER_ERROR_CODE    = 2040
ADMIN_RETRIEVE_CLUSTER_ERROR_MESSAGE = "Error retrieving cluster: %s"

ADMIN_ADD_CLUSTER_ERROR_CODE    = 2050
ADMIN_ADD_CLUSTER_ERROR_MESSAGE = "Error adding cluster: %s"

ADMIN_REMOVE_CLUSTER_ERROR_CODE    = 2060
ADMIN_REMOVE_CLUSTER_ERROR_MESSAGE = "Error removing cluster: %s"

ADMIN_GET_NODE_ERROR_CODE    = 2070
ADMIN_GET_NODE_ERROR_MESSAGE = "Error retrieving node: %s"

ADMIN_NO_SUCH_NODE_CODE    = 2080
ADMIN_NO_SUCH_NODE_MESSAGE = "No such node %s"

ADMIN_ADD_NODE_ERROR_CODE    = 2090
ADMIN_ADD_NODE_ERROR_MESSAGE = "Error adding node: %s"

ADMIN_REMOVE_NODE_ERROR_CODE    = 2100
ADMIN_REMOVE_NODE_ERROR_MESSAGE = "Error removing node: %s"

ADMIN_MAKE_METRIC_ERROR_CODE    = 2110
ADMIN_MAKE_METRIC_ERROR_MESSAGE = "Error creating metric: %s "
)
