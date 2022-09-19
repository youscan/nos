package v1alpha1

import (
	"context"
	"fmt"
	"github.com/nebuly-ai/nebulnetes/pkg/constant"
	"github.com/nebuly-ai/nebulnetes/pkg/util"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var ceqLog = logf.Log.WithName("compositeelasticquota-resource")

func (r *CompositeElasticQuota) SetupWebhookWithManager(mgr ctrl.Manager) error {
	if client == nil {
		client = mgr.GetClient()
	}
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/validate-n8s-nebuly-ai-v1alpha1-compositeelasticquota,mutating=false,failurePolicy=fail,sideEffects=None,groups=n8s.nebuly.ai,resources=compositeelasticquotas,verbs=create;update,versions=v1alpha1,name=vcompositeelasticquota.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &CompositeElasticQuota{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CompositeElasticQuota) ValidateCreate() error {
	ceqLog.V(1).Info("validate create", "name", r.Name)
	return validateCompositeElasticQuotaNamespaces(r)
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CompositeElasticQuota) ValidateUpdate(old runtime.Object) error {
	ceqLog.V(1).Info("validate update", "name", r.Name)
	return validateCompositeElasticQuotaNamespaces(r)
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CompositeElasticQuota) ValidateDelete() error {
	return nil
}

// validateCompositeElasticQuotaNamespaces checks if the specified namespaces are subject to
// any other CompositeElasticQuota: if so it returns an error
func validateCompositeElasticQuotaNamespaces(instance *CompositeElasticQuota) error {
	var ceqList CompositeElasticQuotaList
	if err := client.List(context.Background(), &ceqList); err != nil {
		eqlog.Error(err, "unable to list composite elastic quotas")
		return fmt.Errorf(constant.InternalErrorMsg)
	}
	for _, ceq := range ceqList.Items {
		for _, ns := range instance.Spec.Namespaces {
			if util.InSlice(ns, ceq.Spec.Namespaces) {
				return fmt.Errorf(
					"a namespace can belong to only 1 CompositeElasticQuota: "+
						"namespace %q already belongs to CompositeElasticQuota \"%s/%s\"",
					ns,
					ceq.Namespace,
					ceq.Name,
				)
			}
		}
	}
	return nil
}