// generated by your friendly code generator. DO NOT EDIT.
// to refresh this file, run `go generate` in your shell.

package updater

import (
	"fmt"

	iam_cnrm_cloud_google_com_v1alpha1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1alpha1"
	networking_istio_io_v1alpha3 "github.com/nais/naiserator/pkg/apis/networking.istio.io/v1alpha3"
	"github.com/nais/naiserator/pkg/apis/rbac.istio.io/v1alpha1"
	storage_cnrm_cloud_google_com_v1alpha2 "github.com/nais/naiserator/pkg/apis/storage.cnrm.cloud.google.com/v1alpha2"
	clientV1Alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned"
	typed_iam_cnrm_cloud_google_com_v1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/iam.cnrm.cloud.google.com/v1alpha1"
	typed_networking_istio_io_v1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/networking.istio.io/v1alpha3"
	istio_v1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/rbac.istio.io/v1alpha1"
	typed_storage_cnrm_cloud_google_com_v1alpha2 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/storage.cnrm.cloud.google.com/v1alpha2"
	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	networkingv1 "k8s.io/api/networking/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	typed_apps_v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	typed_autoscaling_v1 "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	typed_core_v1 "k8s.io/client-go/kubernetes/typed/core/v1"
	typed_extensions_v1beta1 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	typed_networking_v1 "k8s.io/client-go/kubernetes/typed/networking/v1"
	typed_rbac_v1 "k8s.io/client-go/kubernetes/typed/rbac/v1"
)

func service(client typed_core_v1.ServiceInterface, old, new *corev1.Service) func() error {
	log.Infof("creating or updating *corev1.Service for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	CopyService(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func serviceAccount(client typed_core_v1.ServiceAccountInterface, old, new *corev1.ServiceAccount) func() error {
	log.Infof("creating or updating *corev1.ServiceAccount for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func deployment(client typed_apps_v1.DeploymentInterface, old, new *appsv1.Deployment) func() error {
	log.Infof("creating or updating *appsv1.Deployment for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func ingress(client typed_extensions_v1beta1.IngressInterface, old, new *extensionsv1beta1.Ingress) func() error {
	log.Infof("creating or updating *extensionsv1beta1.Ingress for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func horizontalPodAutoscaler(client typed_autoscaling_v1.HorizontalPodAutoscalerInterface, old, new *autoscalingv1.HorizontalPodAutoscaler) func() error {
	log.Infof("creating or updating *autoscalingv1.HorizontalPodAutoscaler for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func networkPolicy(client typed_networking_v1.NetworkPolicyInterface, old, new *networkingv1.NetworkPolicy) func() error {
	log.Infof("creating or updating *networkingv1.NetworkPolicy for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func serviceRole(client istio_v1alpha1.ServiceRoleInterface, old, new *v1alpha1.ServiceRole) func() error {
	log.Infof("creating or updating *v1alpha1.ServiceRole for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func serviceRoleBinding(client istio_v1alpha1.ServiceRoleBindingInterface, old, new *v1alpha1.ServiceRoleBinding) func() error {
	log.Infof("creating or updating *v1alpha1.ServiceRoleBinding for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func virtualService(client typed_networking_istio_io_v1alpha3.VirtualServiceInterface, old, new *networking_istio_io_v1alpha3.VirtualService) func() error {
	log.Infof("creating or updating *networking_istio_io_v1alpha3.VirtualService for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func ServiceEntry(client typed_networking_istio_io_v1alpha3.ServiceEntryInterface, old, new *networking_istio_io_v1alpha3.ServiceEntry) func() error {
	log.Infof("creating or updating *networking_istio_io_v1alpha3.ServiceEntry for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func Role(client typed_rbac_v1.RoleInterface, old, new *rbacv1.Role) func() error {
	log.Infof("creating or updating *rbacv1.Role for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func RoleBinding(client typed_rbac_v1.RoleBindingInterface, old, new *rbacv1.RoleBinding) func() error {
	log.Infof("creating or updating *rbacv1.RoleBinding for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func iamServiceAccount(client typed_iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccountInterface, old, new *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount) func() error {
	log.Infof("creating or updating *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func iamPolicy(client typed_iam_cnrm_cloud_google_com_v1alpha1.IAMPolicyInterface, old, new *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy) func() error {
	log.Infof("creating or updating *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func googleStorageBucket(client typed_storage_cnrm_cloud_google_com_v1alpha2.StorageBucketInterface, old, new *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket) func() error {
	log.Infof("creating or updating *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func googleStorageBucketAccessControl(client typed_storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControlInterface, old, new *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl) func() error {
	log.Infof("creating or updating *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl for %s", new.Name)
	if old == nil {
		return func() error {
			_, err := client.Create(new)
			return err
		}
	}

	CopyMeta(old, new)

	return func() error {
		_, err := client.Update(new)
		return err
	}
}

func CreateOrUpdate(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {

	case *corev1.Service:
		c := clientSet.CoreV1().Services(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return service(c, nil, new)
		}
		return service(c, old, new)

	case *corev1.ServiceAccount:
		c := clientSet.CoreV1().ServiceAccounts(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return serviceAccount(c, nil, new)
		}
		return serviceAccount(c, old, new)

	case *appsv1.Deployment:
		c := clientSet.AppsV1().Deployments(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return deployment(c, nil, new)
		}
		return deployment(c, old, new)

	case *extensionsv1beta1.Ingress:
		c := clientSet.ExtensionsV1beta1().Ingresses(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return ingress(c, nil, new)
		}
		return ingress(c, old, new)

	case *autoscalingv1.HorizontalPodAutoscaler:
		c := clientSet.AutoscalingV1().HorizontalPodAutoscalers(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return horizontalPodAutoscaler(c, nil, new)
		}
		return horizontalPodAutoscaler(c, old, new)

	case *networkingv1.NetworkPolicy:
		c := clientSet.NetworkingV1().NetworkPolicies(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return networkPolicy(c, nil, new)
		}
		return networkPolicy(c, old, new)

	case *v1alpha1.ServiceRole:
		c := customClient.RbacV1alpha1().ServiceRoles(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return serviceRole(c, nil, new)
		}
		return serviceRole(c, old, new)

	case *v1alpha1.ServiceRoleBinding:
		c := customClient.RbacV1alpha1().ServiceRoleBindings(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return serviceRoleBinding(c, nil, new)
		}
		return serviceRoleBinding(c, old, new)

	case *networking_istio_io_v1alpha3.VirtualService:
		c := customClient.NetworkingV1alpha3().VirtualServices(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return virtualService(c, nil, new)
		}
		return virtualService(c, old, new)

	case *networking_istio_io_v1alpha3.ServiceEntry:
		c := customClient.NetworkingV1alpha3().ServiceEntries(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return ServiceEntry(c, nil, new)
		}
		return ServiceEntry(c, old, new)

	case *rbacv1.Role:
		c := clientSet.RbacV1().Roles(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return Role(c, nil, new)
		}
		return Role(c, old, new)

	case *rbacv1.RoleBinding:
		c := clientSet.RbacV1().RoleBindings(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return RoleBinding(c, nil, new)
		}
		return RoleBinding(c, old, new)

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount:
		c := customClient.IamV1alpha1().IAMServiceAccounts(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return iamServiceAccount(c, nil, new)
		}
		return iamServiceAccount(c, old, new)

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy:
		c := customClient.IamV1alpha1().IAMPolicies(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return iamPolicy(c, nil, new)
		}
		return iamPolicy(c, old, new)

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket:
		c := customClient.StorageV1alpha2().StorageBuckets(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return googleStorageBucket(c, nil, new)
		}
		return googleStorageBucket(c, old, new)

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl:
		c := customClient.StorageV1alpha2().StorageBucketAccessControls(new.Namespace)
		old, err := c.Get(new.Name, metav1.GetOptions{})
		if err != nil {
			if !errors.IsNotFound(err) {
				return func() error { return err }
			}
			return googleStorageBucketAccessControl(c, nil, new)
		}
		return googleStorageBucketAccessControl(c, old, new)

	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}

func CreateOrRecreate(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {

	case *corev1.Service:
		c := clientSet.CoreV1().Services(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *corev1.Service for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *corev1.Service for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *corev1.ServiceAccount:
		c := clientSet.CoreV1().ServiceAccounts(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *corev1.ServiceAccount for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *corev1.ServiceAccount for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *appsv1.Deployment:
		c := clientSet.AppsV1().Deployments(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *appsv1.Deployment for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *appsv1.Deployment for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *extensionsv1beta1.Ingress:
		c := clientSet.ExtensionsV1beta1().Ingresses(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *extensionsv1beta1.Ingress for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *extensionsv1beta1.Ingress for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *autoscalingv1.HorizontalPodAutoscaler:
		c := clientSet.AutoscalingV1().HorizontalPodAutoscalers(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *autoscalingv1.HorizontalPodAutoscaler for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *autoscalingv1.HorizontalPodAutoscaler for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *networkingv1.NetworkPolicy:
		c := clientSet.NetworkingV1().NetworkPolicies(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *networkingv1.NetworkPolicy for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *networkingv1.NetworkPolicy for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *v1alpha1.ServiceRole:
		c := customClient.RbacV1alpha1().ServiceRoles(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *v1alpha1.ServiceRole for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *v1alpha1.ServiceRole for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *v1alpha1.ServiceRoleBinding:
		c := customClient.RbacV1alpha1().ServiceRoleBindings(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *v1alpha1.ServiceRoleBinding for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *v1alpha1.ServiceRoleBinding for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *networking_istio_io_v1alpha3.VirtualService:
		c := customClient.NetworkingV1alpha3().VirtualServices(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *networking_istio_io_v1alpha3.VirtualService for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *networking_istio_io_v1alpha3.VirtualService for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *networking_istio_io_v1alpha3.ServiceEntry:
		c := customClient.NetworkingV1alpha3().ServiceEntries(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *networking_istio_io_v1alpha3.ServiceEntry for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *networking_istio_io_v1alpha3.ServiceEntry for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *rbacv1.Role:
		c := clientSet.RbacV1().Roles(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *rbacv1.Role for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *rbacv1.Role for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *rbacv1.RoleBinding:
		c := clientSet.RbacV1().RoleBindings(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *rbacv1.RoleBinding for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *rbacv1.RoleBinding for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount:
		c := customClient.IamV1alpha1().IAMServiceAccounts(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy:
		c := customClient.IamV1alpha1().IAMPolicies(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket:
		c := customClient.StorageV1alpha2().StorageBuckets(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl:
		c := customClient.StorageV1alpha2().StorageBucketAccessControls(new.Namespace)
		return func() error {
			log.Infof("pre-deleting *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && !errors.IsNotFound(err) {
				return err
			}
			log.Infof("creating new *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl for %s", new.Name)
			_, err = c.Create(new)
			return err
		}

	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}

func CreateIfNotExists(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {

	case *corev1.Service:
		c := clientSet.CoreV1().Services(new.Namespace)
		return func() error {
			log.Infof("creating new *corev1.Service for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *corev1.ServiceAccount:
		c := clientSet.CoreV1().ServiceAccounts(new.Namespace)
		return func() error {
			log.Infof("creating new *corev1.ServiceAccount for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *appsv1.Deployment:
		c := clientSet.AppsV1().Deployments(new.Namespace)
		return func() error {
			log.Infof("creating new *appsv1.Deployment for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *extensionsv1beta1.Ingress:
		c := clientSet.ExtensionsV1beta1().Ingresses(new.Namespace)
		return func() error {
			log.Infof("creating new *extensionsv1beta1.Ingress for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *autoscalingv1.HorizontalPodAutoscaler:
		c := clientSet.AutoscalingV1().HorizontalPodAutoscalers(new.Namespace)
		return func() error {
			log.Infof("creating new *autoscalingv1.HorizontalPodAutoscaler for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *networkingv1.NetworkPolicy:
		c := clientSet.NetworkingV1().NetworkPolicies(new.Namespace)
		return func() error {
			log.Infof("creating new *networkingv1.NetworkPolicy for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *v1alpha1.ServiceRole:
		c := customClient.RbacV1alpha1().ServiceRoles(new.Namespace)
		return func() error {
			log.Infof("creating new *v1alpha1.ServiceRole for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *v1alpha1.ServiceRoleBinding:
		c := customClient.RbacV1alpha1().ServiceRoleBindings(new.Namespace)
		return func() error {
			log.Infof("creating new *v1alpha1.ServiceRoleBinding for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *networking_istio_io_v1alpha3.VirtualService:
		c := customClient.NetworkingV1alpha3().VirtualServices(new.Namespace)
		return func() error {
			log.Infof("creating new *networking_istio_io_v1alpha3.VirtualService for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *networking_istio_io_v1alpha3.ServiceEntry:
		c := customClient.NetworkingV1alpha3().ServiceEntries(new.Namespace)
		return func() error {
			log.Infof("creating new *networking_istio_io_v1alpha3.ServiceEntry for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *rbacv1.Role:
		c := clientSet.RbacV1().Roles(new.Namespace)
		return func() error {
			log.Infof("creating new *rbacv1.Role for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *rbacv1.RoleBinding:
		c := clientSet.RbacV1().RoleBindings(new.Namespace)
		return func() error {
			log.Infof("creating new *rbacv1.RoleBinding for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount:
		c := customClient.IamV1alpha1().IAMServiceAccounts(new.Namespace)
		return func() error {
			log.Infof("creating new *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy:
		c := customClient.IamV1alpha1().IAMPolicies(new.Namespace)
		return func() error {
			log.Infof("creating new *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket:
		c := customClient.StorageV1alpha2().StorageBuckets(new.Namespace)
		return func() error {
			log.Infof("creating new *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl:
		c := customClient.StorageV1alpha2().StorageBucketAccessControls(new.Namespace)
		return func() error {
			log.Infof("creating new *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl for %s", new.Name)
			_, err := c.Create(new)
			if err != nil && !errors.IsAlreadyExists(err) {
				return err
			}
			return nil
		}

	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}

func DeleteIfExists(clientSet kubernetes.Interface, customClient clientV1Alpha1.Interface, resource runtime.Object) func() error {
	switch new := resource.(type) {

	case *corev1.Service:
		c := clientSet.CoreV1().Services(new.Namespace)
		return func() error {
			log.Infof("deleting *corev1.Service for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *corev1.ServiceAccount:
		c := clientSet.CoreV1().ServiceAccounts(new.Namespace)
		return func() error {
			log.Infof("deleting *corev1.ServiceAccount for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *appsv1.Deployment:
		c := clientSet.AppsV1().Deployments(new.Namespace)
		return func() error {
			log.Infof("deleting *appsv1.Deployment for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *extensionsv1beta1.Ingress:
		c := clientSet.ExtensionsV1beta1().Ingresses(new.Namespace)
		return func() error {
			log.Infof("deleting *extensionsv1beta1.Ingress for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *autoscalingv1.HorizontalPodAutoscaler:
		c := clientSet.AutoscalingV1().HorizontalPodAutoscalers(new.Namespace)
		return func() error {
			log.Infof("deleting *autoscalingv1.HorizontalPodAutoscaler for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *networkingv1.NetworkPolicy:
		c := clientSet.NetworkingV1().NetworkPolicies(new.Namespace)
		return func() error {
			log.Infof("deleting *networkingv1.NetworkPolicy for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *v1alpha1.ServiceRole:
		c := customClient.RbacV1alpha1().ServiceRoles(new.Namespace)
		return func() error {
			log.Infof("deleting *v1alpha1.ServiceRole for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *v1alpha1.ServiceRoleBinding:
		c := customClient.RbacV1alpha1().ServiceRoleBindings(new.Namespace)
		return func() error {
			log.Infof("deleting *v1alpha1.ServiceRoleBinding for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *networking_istio_io_v1alpha3.VirtualService:
		c := customClient.NetworkingV1alpha3().VirtualServices(new.Namespace)
		return func() error {
			log.Infof("deleting *networking_istio_io_v1alpha3.VirtualService for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *networking_istio_io_v1alpha3.ServiceEntry:
		c := customClient.NetworkingV1alpha3().ServiceEntries(new.Namespace)
		return func() error {
			log.Infof("deleting *networking_istio_io_v1alpha3.ServiceEntry for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *rbacv1.Role:
		c := clientSet.RbacV1().Roles(new.Namespace)
		return func() error {
			log.Infof("deleting *rbacv1.Role for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *rbacv1.RoleBinding:
		c := clientSet.RbacV1().RoleBindings(new.Namespace)
		return func() error {
			log.Infof("deleting *rbacv1.RoleBinding for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount:
		c := customClient.IamV1alpha1().IAMServiceAccounts(new.Namespace)
		return func() error {
			log.Infof("deleting *iam_cnrm_cloud_google_com_v1alpha1.IAMServiceAccount for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy:
		c := customClient.IamV1alpha1().IAMPolicies(new.Namespace)
		return func() error {
			log.Infof("deleting *iam_cnrm_cloud_google_com_v1alpha1.IAMPolicy for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket:
		c := customClient.StorageV1alpha2().StorageBuckets(new.Namespace)
		return func() error {
			log.Infof("deleting *storage_cnrm_cloud_google_com_v1alpha2.StorageBucket for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	case *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl:
		c := customClient.StorageV1alpha2().StorageBucketAccessControls(new.Namespace)
		return func() error {
			log.Infof("deleting *storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl for %s", new.Name)
			err := c.Delete(new.Name, &metav1.DeleteOptions{})
			if err != nil && errors.IsNotFound(err) {
				return nil
			}
			return err
		}

	default:
		panic(fmt.Errorf("BUG! You didn't specify a case for type '%T' in the file hack/generator/updater.go", new))
	}
}
