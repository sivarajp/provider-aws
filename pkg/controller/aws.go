/*
Copyright 2019 The Crossplane Authors.

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

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/controller"

	"github.com/crossplane-contrib/provider-aws/pkg/controller/config"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/accesskey"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/group"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/grouppolicyattachment"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/groupusermembership"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/policy"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/role"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/rolepolicyattachment"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/user"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/iam/userpolicyattachment"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/s3"
	"github.com/crossplane-contrib/provider-aws/pkg/controller/s3/bucketpolicy"
)

// Setup creates all AWS controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		s3.SetupBucket,
		bucketpolicy.SetupBucketPolicy,
		accesskey.SetupAccessKey,
		user.SetupUser,
		group.SetupGroup,
		policy.SetupPolicy,
		role.SetupRole,
		groupusermembership.SetupGroupUserMembership,
		userpolicyattachment.SetupUserPolicyAttachment,
		grouppolicyattachment.SetupGroupPolicyAttachment,
		rolepolicyattachment.SetupRolePolicyAttachment,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}

	return config.Setup(mgr, o)
}
