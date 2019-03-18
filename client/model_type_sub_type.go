/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// TypeSubtype {1510}
type TypeSubType struct {
	TypeCode TypeCodeEnum `json:"typeCode"`
	SubTypeCode SubTypeCodeEnum `json:"subTypeCode"`
}