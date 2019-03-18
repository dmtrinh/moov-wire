/*
 * WIRE API
 *
 * Moov WIRE () implements an HTTP API for creating, parsing and validating WIRE files.
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PaymentMethodToBeneficiary struct {
	// PaymentMethod
	PaymentMethod string `json:"paymentMethod,omitempty"`
	Additional string `json:"Additional,omitempty"`
}