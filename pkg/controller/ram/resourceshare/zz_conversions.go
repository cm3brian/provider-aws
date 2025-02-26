/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package resourceshare

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/ram"

	svcapitypes "github.com/crossplane/provider-aws/apis/ram/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateGetResourceSharesInput returns input for read
// operation.
func GenerateGetResourceSharesInput(cr *svcapitypes.ResourceShare) *svcsdk.GetResourceSharesInput {
	res := &svcsdk.GetResourceSharesInput{}

	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// GenerateResourceShare returns the current state in the form of *svcapitypes.ResourceShare.
func GenerateResourceShare(resp *svcsdk.GetResourceSharesOutput) *svcapitypes.ResourceShare {
	cr := &svcapitypes.ResourceShare{}

	found := false
	for _, elem := range resp.ResourceShares {
		if elem.AllowExternalPrincipals != nil {
			cr.Spec.ForProvider.AllowExternalPrincipals = elem.AllowExternalPrincipals
		} else {
			cr.Spec.ForProvider.AllowExternalPrincipals = nil
		}
		if elem.Name != nil {
			cr.Spec.ForProvider.Name = elem.Name
		} else {
			cr.Spec.ForProvider.Name = nil
		}
		if elem.Tags != nil {
			f9 := []*svcapitypes.Tag{}
			for _, f9iter := range elem.Tags {
				f9elem := &svcapitypes.Tag{}
				if f9iter.Key != nil {
					f9elem.Key = f9iter.Key
				}
				if f9iter.Value != nil {
					f9elem.Value = f9iter.Value
				}
				f9 = append(f9, f9elem)
			}
			cr.Spec.ForProvider.Tags = f9
		} else {
			cr.Spec.ForProvider.Tags = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateResourceShareInput returns a create input.
func GenerateCreateResourceShareInput(cr *svcapitypes.ResourceShare) *svcsdk.CreateResourceShareInput {
	res := &svcsdk.CreateResourceShareInput{}

	if cr.Spec.ForProvider.AllowExternalPrincipals != nil {
		res.SetAllowExternalPrincipals(*cr.Spec.ForProvider.AllowExternalPrincipals)
	}
	if cr.Spec.ForProvider.ClientToken != nil {
		res.SetClientToken(*cr.Spec.ForProvider.ClientToken)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}
	if cr.Spec.ForProvider.PermissionARNs != nil {
		f3 := []*string{}
		for _, f3iter := range cr.Spec.ForProvider.PermissionARNs {
			var f3elem string
			f3elem = *f3iter
			f3 = append(f3, &f3elem)
		}
		res.SetPermissionArns(f3)
	}
	if cr.Spec.ForProvider.Principals != nil {
		f4 := []*string{}
		for _, f4iter := range cr.Spec.ForProvider.Principals {
			var f4elem string
			f4elem = *f4iter
			f4 = append(f4, &f4elem)
		}
		res.SetPrincipals(f4)
	}
	if cr.Spec.ForProvider.ResourceARNs != nil {
		f5 := []*string{}
		for _, f5iter := range cr.Spec.ForProvider.ResourceARNs {
			var f5elem string
			f5elem = *f5iter
			f5 = append(f5, &f5elem)
		}
		res.SetResourceArns(f5)
	}
	if cr.Spec.ForProvider.Tags != nil {
		f6 := []*svcsdk.Tag{}
		for _, f6iter := range cr.Spec.ForProvider.Tags {
			f6elem := &svcsdk.Tag{}
			if f6iter.Key != nil {
				f6elem.SetKey(*f6iter.Key)
			}
			if f6iter.Value != nil {
				f6elem.SetValue(*f6iter.Value)
			}
			f6 = append(f6, f6elem)
		}
		res.SetTags(f6)
	}

	return res
}

// GenerateUpdateResourceShareInput returns an update input.
func GenerateUpdateResourceShareInput(cr *svcapitypes.ResourceShare) *svcsdk.UpdateResourceShareInput {
	res := &svcsdk.UpdateResourceShareInput{}

	if cr.Spec.ForProvider.AllowExternalPrincipals != nil {
		res.SetAllowExternalPrincipals(*cr.Spec.ForProvider.AllowExternalPrincipals)
	}
	if cr.Spec.ForProvider.ClientToken != nil {
		res.SetClientToken(*cr.Spec.ForProvider.ClientToken)
	}
	if cr.Spec.ForProvider.Name != nil {
		res.SetName(*cr.Spec.ForProvider.Name)
	}

	return res
}

// GenerateDeleteResourceShareInput returns a deletion input.
func GenerateDeleteResourceShareInput(cr *svcapitypes.ResourceShare) *svcsdk.DeleteResourceShareInput {
	res := &svcsdk.DeleteResourceShareInput{}

	if cr.Spec.ForProvider.ClientToken != nil {
		res.SetClientToken(*cr.Spec.ForProvider.ClientToken)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "UNKNOWN"
}
