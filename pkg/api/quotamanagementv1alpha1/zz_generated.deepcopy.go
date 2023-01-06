//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package quotamanagementv1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationConstantStrategy) DeepCopyInto(out *AllocationConstantStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationConstantStrategy.
func (in *AllocationConstantStrategy) DeepCopy() *AllocationConstantStrategy {
	if in == nil {
		return nil
	}
	out := new(AllocationConstantStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationDelayedStrategy) DeepCopyInto(out *AllocationDelayedStrategy) {
	*out = *in
	in.ExpectedUsageDate.DeepCopyInto(&out.ExpectedUsageDate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationDelayedStrategy.
func (in *AllocationDelayedStrategy) DeepCopy() *AllocationDelayedStrategy {
	if in == nil {
		return nil
	}
	out := new(AllocationDelayedStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationDisasterRecoveryStrategy) DeepCopyInto(out *AllocationDisasterRecoveryStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationDisasterRecoveryStrategy.
func (in *AllocationDisasterRecoveryStrategy) DeepCopy() *AllocationDisasterRecoveryStrategy {
	if in == nil {
		return nil
	}
	out := new(AllocationDisasterRecoveryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationPeriodicStrategy) DeepCopyInto(out *AllocationPeriodicStrategy) {
	*out = *in
	out.Interval = in.Interval
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationPeriodicStrategy.
func (in *AllocationPeriodicStrategy) DeepCopy() *AllocationPeriodicStrategy {
	if in == nil {
		return nil
	}
	out := new(AllocationPeriodicStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllocationStrategy) DeepCopyInto(out *AllocationStrategy) {
	*out = *in
	if in.Delayed != nil {
		in, out := &in.Delayed, &out.Delayed
		*out = new(AllocationDelayedStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.DisasterRecovery != nil {
		in, out := &in.DisasterRecovery, &out.DisasterRecovery
		*out = new(AllocationDisasterRecoveryStrategy)
		**out = **in
	}
	if in.Periodic != nil {
		in, out := &in.Periodic, &out.Periodic
		*out = new(AllocationPeriodicStrategy)
		**out = **in
	}
	if in.Constant != nil {
		in, out := &in.Constant, &out.Constant
		*out = new(AllocationConstantStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllocationStrategy.
func (in *AllocationStrategy) DeepCopy() *AllocationStrategy {
	if in == nil {
		return nil
	}
	out := new(AllocationStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaDescriptor) DeepCopyInto(out *ResourceQuotaDescriptor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaDescriptor.
func (in *ResourceQuotaDescriptor) DeepCopy() *ResourceQuotaDescriptor {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceQuotaDescriptor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaDescriptorCondition) DeepCopyInto(out *ResourceQuotaDescriptorCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaDescriptorCondition.
func (in *ResourceQuotaDescriptorCondition) DeepCopy() *ResourceQuotaDescriptorCondition {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaDescriptorCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaDescriptorList) DeepCopyInto(out *ResourceQuotaDescriptorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceQuotaDescriptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaDescriptorList.
func (in *ResourceQuotaDescriptorList) DeepCopy() *ResourceQuotaDescriptorList {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaDescriptorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceQuotaDescriptorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaDescriptorSpec) DeepCopyInto(out *ResourceQuotaDescriptorSpec) {
	*out = *in
	out.ResourceQuotaRef = in.ResourceQuotaRef
	in.AllocationStrategy.DeepCopyInto(&out.AllocationStrategy)
	if in.TargetAllocationsPolicy != nil {
		in, out := &in.TargetAllocationsPolicy, &out.TargetAllocationsPolicy
		*out = new(TargetAllocationsPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaDescriptorSpec.
func (in *ResourceQuotaDescriptorSpec) DeepCopy() *ResourceQuotaDescriptorSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaDescriptorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaDescriptorStatus) DeepCopyInto(out *ResourceQuotaDescriptorStatus) {
	*out = *in
	if in.ObservationWindowStart != nil {
		in, out := &in.ObservationWindowStart, &out.ObservationWindowStart
		*out = (*in).DeepCopy()
	}
	if in.ObservationWindowEnd != nil {
		in, out := &in.ObservationWindowEnd, &out.ObservationWindowEnd
		*out = (*in).DeepCopy()
	}
	if in.ObservationWindowLength != nil {
		in, out := &in.ObservationWindowLength, &out.ObservationWindowLength
		*out = new(v1.Duration)
		**out = **in
	}
	if in.QuotaRightsizingEnabled != nil {
		in, out := &in.QuotaRightsizingEnabled, &out.QuotaRightsizingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.GracePeriodLength != nil {
		in, out := &in.GracePeriodLength, &out.GracePeriodLength
		*out = new(v1.Duration)
		**out = **in
	}
	if in.ScheduledRightsizeDate != nil {
		in, out := &in.ScheduledRightsizeDate, &out.ScheduledRightsizeDate
		*out = (*in).DeepCopy()
	}
	if in.ActualQuota != nil {
		in, out := &in.ActualQuota, &out.ActualQuota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.ProposedQuota != nil {
		in, out := &in.ProposedQuota, &out.ProposedQuota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.PreRightsizeQuota != nil {
		in, out := &in.PreRightsizeQuota, &out.PreRightsizeQuota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.MaxUsedQuota != nil {
		in, out := &in.MaxUsedQuota, &out.MaxUsedQuota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.P95UsedQuota != nil {
		in, out := &in.P95UsedQuota, &out.P95UsedQuota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.AvgUsedQuota != nil {
		in, out := &in.AvgUsedQuota, &out.AvgUsedQuota
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.LastRightsizeTime != nil {
		in, out := &in.LastRightsizeTime, &out.LastRightsizeTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaDescriptorStatus.
func (in *ResourceQuotaDescriptorStatus) DeepCopy() *ResourceQuotaDescriptorStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaDescriptorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetAllocationsPolicy) DeepCopyInto(out *TargetAllocationsPolicy) {
	*out = *in
	if in.LimitsTargetPercent != nil {
		in, out := &in.LimitsTargetPercent, &out.LimitsTargetPercent
		*out = make(map[corev1.ResourceName]AllocationPercent, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RequestsTargetPercent != nil {
		in, out := &in.RequestsTargetPercent, &out.RequestsTargetPercent
		*out = make(map[corev1.ResourceName]AllocationPercent, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetAllocationsPolicy.
func (in *TargetAllocationsPolicy) DeepCopy() *TargetAllocationsPolicy {
	if in == nil {
		return nil
	}
	out := new(TargetAllocationsPolicy)
	in.DeepCopyInto(out)
	return out
}
