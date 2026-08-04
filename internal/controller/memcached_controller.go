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
<<<<<<< HEAD
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
=======
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"time"
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb

	cachev1alpha1 "github.com/example/memcached-operator/api/v1alpha1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

<<<<<<< HEAD
const (
	deploymentName      = "Deployment.Name"
	deploymentNamespace = "Deployment.Namespace"
)

=======
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
func (r *MemcachedReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("Starting reconciliation", "namespace", req.Namespace, "name", req.Name)
	log.Info("Cloudbuild TEST")

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
=======
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Memcached object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *MemcachedReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	//_ = log.FromContext(ctx)
	log := log.FromContext(ctx)
	// TODO(user): your logic here

	log.Info(fmt.Sprintf("START : CCY is here for %s POC", req.Namespace))

	memcached := &cachev1alpha1.Memcached{}
	//memcached.Spec.Foo
	//memcached.Spec.Installap
	//memcached.Spec.Installdb
	//memcached.Spec.Connstr

	err := r.Get(ctx, req.NamespacedName, memcached)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// If the custom resource is not found then it usually means that it was deleted or not created
			// In this way, we will stop the reconciliation
			log.Info("memcached resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "Failed to get memcached")
		return ctrl.Result{}, err
	}

	log.Info(fmt.Sprintf("[CCY]Reconcile : CCY is here for NS:%s CR (Size:%d, Foo:%s)", req.Namespace, memcached.Spec.Size, memcached.Spec.Foo))

	if memcached.Spec.Installap {

		dep := GetDeployOfApp(memcached)
		log.Info("[CCY]Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)

		if err = r.Create(ctx, dep); err != nil {
			log.Error(err, "[CCY]Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			//return ctrl.Result{}, err
		}

	} else {

		found := &appsv1.Deployment{}
		log.Info("[CCY]Deleting a new Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

		//err = r.Get(ctx, types.NamespacedName{Name: "ap", Namespace: memcached.Namespace}, found)
		err = r.Get(ctx, types.NamespacedName{Name: "client", Namespace: memcached.Namespace}, found)
		r.Delete(ctx, found)

	}

	if memcached.Spec.Installdb {

		//dep := GetDeployOfDb(memcached)
		dep := GetDeployOfSvr(memcached)
		log.Info("[CCY]Creating a new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)

		if err = r.Create(ctx, dep); err != nil {
			log.Error(err, "[CCY]Failed to create new Deployment", "Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
			//return ctrl.Result{}, err
		}

	} else {

		found := &appsv1.Deployment{}
		log.Info("[CCY]Deleting a new Deployment", "Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

		//err = r.Get(ctx, types.NamespacedName{Name: "db", Namespace: memcached.Namespace}, found)
		err = r.Get(ctx, types.NamespacedName{Name: "server", Namespace: memcached.Namespace}, found)
		r.Delete(ctx, found)

	}

	log.Info(fmt.Sprintf("END : CCY is here for %s POC", req.Namespace))
	return ctrl.Result{RequeueAfter: time.Minute}, nil
	//return ctrl.Result{Requeue: true}, nil
	//return ctrl.Result{}, nil
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
}

// SetupWithManager sets up the controller with the Manager.
func (r *MemcachedReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.Memcached{}).
		Complete(r)
}

func GetDeployOfDb(memcached *cachev1alpha1.Memcached) *appsv1.Deployment {
<<<<<<< HEAD
=======
	//replicas := memcached.Spec.Size * 2

>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "db",
			Namespace: memcached.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &memcached.Spec.Size,
			Selector: &metav1.LabelSelector{
<<<<<<< HEAD
=======
				//MatchLabels: nil,
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
				MatchLabels: map[string]string{"db": "demo-db"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
<<<<<<< HEAD
=======
					//Labels: nil,
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
						Image: "registry.access.redhat.com/rhscl/mysql-80-rhel7:latest",
=======
						Image: "registry.access.redhat.com/rhscl/mysql-80-rhel7",
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
=======
						//Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
					}},
				},
			},
		},
	}

	return dep
}

func GetDeployOfApp(memcached *cachev1alpha1.Memcached) *appsv1.Deployment {
<<<<<<< HEAD
=======
	//replicas := memcached.Spec.Size * 2

>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "client",
			Namespace: memcached.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &memcached.Spec.Size,
			Selector: &metav1.LabelSelector{
<<<<<<< HEAD
=======
				//MatchLabels: nil,
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
				MatchLabels: map[string]string{"app": "demo-client"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
<<<<<<< HEAD
=======
					//Labels: nil,
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
						Image:           "registry.access.redhat.com/rhscl/httpd-24-rhel7:latest",
=======
						Image:           "registry.access.redhat.com/rhscl/httpd-24-rhel7",
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
=======
						//Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
					}},
				},
			},
		},
	}

	return dep
}

func GetDeployOfSvr(memcached *cachev1alpha1.Memcached) *appsv1.Deployment {
<<<<<<< HEAD
=======
	//replicas := memcached.Spec.Size * 2

>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "server",
			Namespace: memcached.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &memcached.Spec.Size,
			Selector: &metav1.LabelSelector{
<<<<<<< HEAD
=======
				//MatchLabels: nil,
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
				MatchLabels: map[string]string{"app": "demo-svr"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
<<<<<<< HEAD
=======
					//Labels: nil,
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
						Image:           "registry.access.redhat.com/rhscl/httpd-24-rhel7:latest",
=======
						Image:           "registry.access.redhat.com/rhscl/httpd-24-rhel7",
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
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
<<<<<<< HEAD
=======
						//Command: []string{"memcached", "-m=64", "-o", "modern", "-v"},
>>>>>>> 75b7ee9545d0fa1562fe2e958a4a6063445d5bbb
					}},
				},
			},
		},
	}

	return dep
}
