// Code generated by openapi-generator. DO NOT EDIT.
/*
    * My Auth Service
    *
    * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
    *
    * API version: 1.0.0
* Generated by: OpenAPI Generator (https://openapi-generator.tech)
*/

package gen

import (
	"github.com/gin-gonic/gin"
)

type RoleAPI interface {
	// RolesIdDelete - Delete Role
	RolesIdDelete(ctx *gin.Context)
	// RolesPost - Create Role
	RolesPost(ctx *gin.Context)
}