// This file won't be included in the client build
// +build modhack

package kafkamgmt

// Necessary for safely adding multi-module repo. See https://github.com/golang/go/wiki/Modules#is-it-possible-to-add-a-module-to-a-multi-module-repository
import _ "github.com/redhat-developer/app-services-sdk-go"