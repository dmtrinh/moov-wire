/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// IdentificationTypeEnum : IdentificationType  * `OI` - Organization ID * `PI` - Private ID 
type IdentificationTypeEnum string

// List of IdentificationTypeEnum
const (
	OI IdentificationTypeEnum = "OI"
	PI IdentificationTypeEnum = "PI"
)