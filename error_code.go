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

// If a query is submitted using an HTTP method not supported by the query engine
// then this error is returned. 
SERVICE_UNSUPPORTED_HTTP_CODE    = 1010
SERVICE_UNSUPPORTED_HTTP_MESSAGE = "Unsupported http method: %s"

// For a request that includes a parameter value that is not yet implemented. 
// Example: ZIP compression not yet implemented
SERVICE_NOT_YET_IMPLEMENTED_CODE    = 1020
SERVICE_NOT_YET_IMPLEMENTED_MESSAGE = "%s %s not yet implemented"

// For a request that includes a parameter value that is not recognized by the query engine.
// Example: Unknown scan_consistency value: foo
SERVICE_UNRECOGNIZED_VALUE_CODE    = 1030
SERVICE_UNRECOGNIZED_VALUE_MESSAGE = "Unknown %s value: %s"

// Returned if there is an internal error while processing a request.
SERVICE_INTERNAL_ERROR_CODE    = 1040
SERVICE_INTERNAL_ERROR_MESSAGE = "Error processing %s"

// Returned if a request is missing a required parameter.
// Example: No statement or prepared value
SERVICE_MISSING_REQUIRED_CODE    = 1050
SERVICE_MISSING_REQUIRED_MESSAGE = "No %s value"

// Returned if a request includes multiple values of a parameter
SERVICE_MULTIPLE_VALUES_CODE    = 1060
SERVICE_MULTIPLE_VALUES_MESSAGE = "Multiple values for %s."

// Returned if the value for a parameter is of incorrect type
// Example: pretty has to be of type boolean
SERVICE_TYPE_MISMATCH_CODE    = 1070
SERVICE_TYPE_MISMATCH_MESSAGE = "%s has to be of type %s"

// Returned if results contain an invalid json object
SERVICE_INVALID_JSON_CODE    = 1100
SERVICE_INVALID_JSON_MESSAGE = "Invalid JSON in results"
)

const (
// ADMIN errors are returned by the clustering API and accounting API

// Returned if there is an errror connecting to a third party component
ADMIN_CONNECTION_ERROR_CODE    = 2000
ADMIN_CONNECTION_ERROR_MESSAGE = "Error connecting to %S"

// Returned if the URL for a component cannot be recognized
// Example: Invalid datastore url: htttp://hostname:8091
ADMIN_INVALID_URL_CODE    = 2010
ADMIN_INVALID_URL_MESSAGE = "Invalid % url: %s"

// Returned if an error occurs reading a JSON payload
ADMIN_JSON_DECODING_ERROR_CODE    = 2020
ADMIN_JSON_DECODING_ERROR_MESSAGE = "Error in JSON decoding: %s"

// Returned if an error occurs when writing a JSON payload
ADMIN_JSON_ENCODING_ERROR_CODE    = 2030
ADMIN_JSON_ENCODING_ERROR_MESSAGE = "Error in JSON encoding: %s"

// If an error occurs retrieving a cluster configuration
ADMIN_RETRIEVE_CLUSTER_ERROR_CODE    = 2040
ADMIN_RETRIEVE_CLUSTER_ERROR_MESSAGE = "Error retrieving cluster: %s"

// If an error occurs adding a cluster configuration
ADMIN_ADD_CLUSTER_ERROR_CODE    = 2050
ADMIN_ADD_CLUSTER_ERROR_MESSAGE = "Error adding cluster: %s"

// If an error occurs removing a cluster configuration
ADMIN_REMOVE_CLUSTER_ERROR_CODE    = 2060
ADMIN_REMOVE_CLUSTER_ERROR_MESSAGE = "Error removing cluster: %s"

// If an error occurs retrieving a query node configuration
ADMIN_GET_NODE_ERROR_CODE    = 2070
ADMIN_GET_NODE_ERROR_MESSAGE = "Error retrieving node: %s"

// If there is a request for the configuration of a non-existent query node
ADMIN_NO_SUCH_NODE_CODE    = 2080
ADMIN_NO_SUCH_NODE_MESSAGE = "No such node %s"

// If an error occurs adding a query node configuration
ADMIN_ADD_NODE_ERROR_CODE    = 2090
ADMIN_ADD_NODE_ERROR_MESSAGE = "Error adding node: %s"

// If an error occurs removing a query node configuration
ADMIN_REMOVE_NODE_ERROR_CODE    = 2100
ADMIN_REMOVE_NODE_ERROR_MESSAGE = "Error removing node: %s"

// If an error occurs adding a metric 
ADMIN_MAKE_METRIC_ERROR_CODE    = 2110
ADMIN_MAKE_METRIC_ERROR_MESSAGE = "Error creating metric: %s "
)
