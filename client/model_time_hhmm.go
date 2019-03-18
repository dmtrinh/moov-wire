/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type TimeHhmm struct {
	// ReceiptTime  HHMM, based on a 24-hour clock, Eastern Time 
	ReceiptTime string `json:"receiptTime,omitempty"`
}