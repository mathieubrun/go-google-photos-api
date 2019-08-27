package codegenerator

import (
	"strings"
)

func (r *Resource) GetResourceName(resourceName string) string {
	return strings.Title(resourceName)
}

func (r *Resource) GetServiceName(resourceName string) string {
	return r.GetResourceName(resourceName) + "Service"
}
