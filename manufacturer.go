/* 
 * Simple Inventory API
 *
 * This is a simple API
 *
 * OpenAPI spec version: 1.0.0
 * Contact: you@your-company.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type Manufacturer struct {

	Name string `json:"name"`

	HomePage string `json:"homePage,omitempty"`

	Phone string `json:"phone,omitempty"`
}