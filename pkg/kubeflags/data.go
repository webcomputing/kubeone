/*
Copyright 2019 The KubeOne Authors.

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

package kubeflags

var (
	defaultAdmissionControllersv114x = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"PersistentVolumeClaimResize",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}

	defaultAdmissionControllersv115x = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}

	defaultAdmissionControllersv116xv117x = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionWebhook",
		"RuntimeClass",
		"ResourceQuota",
	}

	defaultAdmissionControllersv118xv121x = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"RuntimeClass",
		"CertificateApproval",
		"CertificateSigning",
		"CertificateSubjectRestriction",
		"DefaultIngressClass",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}

	defaultAdmissionControllersv122x = []string{
		"NamespaceLifecycle",
		"LimitRanger",
		"ServiceAccount",
		"TaintNodesByCondition",
		"PodSecurity",
		"Priority",
		"DefaultTolerationSeconds",
		"DefaultStorageClass",
		"StorageObjectInUseProtection",
		"PersistentVolumeClaimResize",
		"RuntimeClass",
		"CertificateApproval",
		"CertificateSigning",
		"CertificateSubjectRestriction",
		"DefaultIngressClass",
		"MutatingAdmissionWebhook",
		"ValidatingAdmissionWebhook",
		"ResourceQuota",
	}
)
