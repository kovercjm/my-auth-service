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

type AuthAPI interface {
	// AuthLoginPost - Authenticate
	AuthLoginPost(ctx *gin.Context)
	// AuthLogoutPost - Invalidate
	AuthLogoutPost(ctx *gin.Context)
	// UsersMeRolesGet - List roles
	UsersMeRolesGet(ctx *gin.Context)
	// UsersMeRolesIdGet - Check Role
	UsersMeRolesIdGet(ctx *gin.Context)
	// UsersUserIDRolesRoleIDPost - Add role to user
	UsersUserIDRolesRoleIDPost(ctx *gin.Context)
}