/*
Copyright 2025.

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
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	deploymentName      = "Deployment.Name"
	deploymentNamespace = "Deployment.Namespace"
)

// MemcachedReconciler reconciles a Memcached object
type MemcachedReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cache.example.com,resources=memcacheds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.example.com,resources=memcacheds/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cache.example.com,resources=memcacheds/finalizers,verbs=update

// +kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.

func (r *MemcachedReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("Starting reconciliation", "namespace", req.Namespace, "name", req.Name)
	log.Info("Cloudbuild COMMIT TEST W/ cloudbuild.yaml")

	memcached := &cachev1alpha1.Memcached{}
	err := r.Get(ctx, req.NamespacedName, memcached)
	if err != nil {
		if apierrors.IsNotFound(err) {
			log.Info("Memcached resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get Memcached")
		return ctrl.Result{}, err
	}

	log.Info("Reconciling Memcached", "namespace", req.Namespace, "size", memcached.Spec.Size, "foo", memcached.Spec.Foo)

	// Reconcile App Deployment
	if memcached.Spec.Installap {
		dep := GetDeployOfApp(memcached)
		found := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{Name: dep.Name, Namespace: dep.Namespace}, found)
		if err != nil && apierrors.IsNotFound(err) {
			log.Info("Creating a new Deployment for App", deploymentNamespace, dep.Namespace, deploymentName, dep.Name)
			if err = r.Create(ctx, dep); err != nil {
				log.Error(err, "Failed to create new Deployment for App", deploymentNamespace, dep.Namespace, deploymentName, dep.Name)
				return ctrl.Result{}, err
			}
		} else if err != nil {
			log.Error(err, "Failed to get Deployment for App")
			return ctrl.Result{}, err
		}
	} else {
		found := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{Name: "client", Namespace: memcached.Namespace}, found)
		if err == nil {
			log.Info("Deleting Deployment for App", deploymentNamespace, found.Namespace, deploymentName, found.Name)
			if err = r.Delete(ctx, found); err != nil {
				log.Error(err, "Failed to delete Deployment for App", deploymentNamespace, found.Namespace, deploymentName, found.Name)
				return ctrl.Result{}, err
			}
		} else if !apierrors.IsNotFound(err) {
			log.Error(err, "Failed to get Deployment for App during deletion")
			return ctrl.Result{}, err
		}
	}

	// Reconcile Svr Deployment
	if memcached.Spec.Installdb {
		dep := GetDeployOfSvr(memcached)
		found := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{Name: dep.Name, Namespace: dep.Namespace}, found)
		if err != nil && apierrors.IsNotFound(err) {
			log.Info("Creating a new Deployment for Server", deploymentNamespace, dep.Namespace, deploymentName, dep.Name)
			if err = r.Create(ctx, dep); err != nil {
				log.Error(err, "Failed to create new Deployment for Server", deploymentNamespace, dep.Namespace, deploymentName, dep.Name)
				return ctrl.Result{}, err
			}
		} else if err != nil {
			log.Error(err, "Failed to get Deployment for Server")
			return ctrl.Result{}, err
		}
	} else {
		found := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{Name: "server", Namespace: memcached.Namespace}, found)
		if err == nil {
			log.Info("Deleting Deployment for Server", deploymentNamespace, found.Namespace, deploymentName, found.Name)
			if err = r.Delete(ctx, found); err != nil {
				log.Error(err, "Failed to delete Deployment for Server", deploymentNamespace, found.Namespace, deploymentName, found.Name)
				return ctrl.Result{}, err
			}
		} else if !apierrors.IsNotFound(err) {
			log.Error(err, "Failed to get Deployment for Server during deletion")
			return ctrl.Result{}, err
		}
	}

	log.Info("Finished reconciliation", "namespace", req.Namespace)
	return ctrl.Result{RequeueAfter: time.Minute}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MemcachedReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Memcached{}).
		Complete(r)
}

func GetDeployOfDb(memcached *cachev1alpha1.Memcached) *appsv1.Deployment {

	//replicas := memcached.Spec.Size * 2

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "db",
			Namespace: memcached.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &memcached.Spec.Size,
			Selector: &metav1.LabelSelector{
				//MatchLabels: nil,
				MatchLabels: map[string]string{"db": "demo-db"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					//Labels: nil,
					Labels: map[string]string{"db": "demo-db"},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: &[]bool{true}[0],
						SeccompProfile: &corev1.SeccompProfile{
							Type: corev1.SeccompProfileTypeRuntimeDefault,
						},
					},
					Containers: []corev1.Container{{
						Image: "registry.access.redhat.com/rhscl/mysql-80-rhel7",
						Name:  "ccydb",
						Env: []corev1.EnvVar{
							{
								Name:  "MYSQL_USER",
								Value: "wordpress",
							},
							{
								Name:  "MYSQL_PASSWORD",
								Value: "password",
							},
							{
								Name:  "MYSQL_DATABASE",
								Value: "wordpress",
							},
						},
						ImagePullPolicy: corev1.PullIfNotPresent,
						SecurityContext: &corev1.SecurityContext{
							RunAsNonRoot:             &[]bool{true}[0],
							RunAsUser:                &[]int64{1001}[0],
							AllowPrivilegeEscalation: &[]bool{false}[0],
							Capabilities: &corev1.Capabilities{
								Drop: []corev1.Capability{
									"ALL",
								},
							},
						},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 3306,
							Name:          "tcp",
						}},
						//Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
					}},
				},
			},
		},
	}

	return dep
}

func GetDeployOfApp(memcached *cachev1alpha1.Memcached) *appsv1.Deployment {

	//replicas := memcached.Spec.Size * 2

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "client",
			Namespace: memcached.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &memcached.Spec.Size,
			Selector: &metav1.LabelSelector{
				//MatchLabels: nil,
				MatchLabels: map[string]string{"app": "demo-client"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					//Labels: nil,
					Labels: map[string]string{"app": "demo-client"},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: &[]bool{true}[0],
						SeccompProfile: &corev1.SeccompProfile{
							Type: corev1.SeccompProfileTypeRuntimeDefault,
						},
					},
					Containers: []corev1.Container{{
						Image:           "registry.access.redhat.com/rhscl/httpd-24-rhel7",
						Name:            "ccyclient",
						ImagePullPolicy: corev1.PullIfNotPresent,
						SecurityContext: &corev1.SecurityContext{
							RunAsNonRoot:             &[]bool{true}[0],
							RunAsUser:                &[]int64{1001}[0],
							AllowPrivilegeEscalation: &[]bool{false}[0],
							Capabilities: &corev1.Capabilities{
								Drop: []corev1.Capability{
									"ALL",
								},
							},
						},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8080,
							Name:          "http",
						}},
						//Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
					}},
				},
			},
		},
	}

	return dep
}

func GetDeployOfSvr(memcached *cachev1alpha1.Memcached) *appsv1.Deployment {
	//replicas := memcached.Spec.Size * 2

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "server",
			Namespace: memcached.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &memcached.Spec.Size,
			Selector: &metav1.LabelSelector{
				//MatchLabels: nil,
				MatchLabels: map[string]string{"app": "demo-svr"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					//Labels: nil,
					Labels: map[string]string{"app": "demo-svr"},
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						RunAsNonRoot: &[]bool{true}[0],
						SeccompProfile: &corev1.SeccompProfile{
							Type: corev1.SeccompProfileTypeRuntimeDefault,
						},
					},
					Containers: []corev1.Container{{
						Image:           "registry.access.redhat.com/rhscl/httpd-24-rhel7",
						Name:            "ccyserver",
						ImagePullPolicy: corev1.PullIfNotPresent,
						SecurityContext: &corev1.SecurityContext{
							RunAsNonRoot:             &[]bool{true}[0],
							RunAsUser:                &[]int64{1001}[0],
							AllowPrivilegeEscalation: &[]bool{false}[0],
							Capabilities: &corev1.Capabilities{
								Drop: []corev1.Capability{
									"ALL",
								},
							},
						},
						Ports: []corev1.ContainerPort{{
							ContainerPort: 8080,
							Name:          "http",
						}},
						//Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
					}},
				},
			},
		},
	}

	return dep
}
