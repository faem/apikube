package cgo

import (
	"github.com/appscode/go/log"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kutilAppsV1 "github.com/appscode/kutil/apps/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)
/*
func CreateDeployment() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	kc := kubernetes.NewForConfigOrDie(config)
	varAppsV1 := kc.AppsV1()

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "apiserver",
			Labels: map[string]string{
				"app": "apiserver",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "apiserver",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "apiserver",
					Labels: map[string]string{
						"app": "apiserver",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "apiserver",
							Image:           "fahimabrar/api",
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "api-port",
									ContainerPort: 8080,
									Protocol:      "TCP",
								},
							},
						},
					},
					RestartPolicy: "Always",
				},
			},
		},
	}

	_, err = varAppsV1.Deployments("default").Create(deployment)

	if err != nil {
		panic(err)
	}
}*/

func CreateDeploymentKutil(){
	log.Debug("Creating Deployment. . . . .")
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	kc := kubernetes.NewForConfigOrDie(config)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "apiserver",
			Namespace: "default",
			Labels: map[string]string{
				"app": "apiserver",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(50),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "apiserver",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name: "apiserver",
					Labels: map[string]string{
						"app": "apiserver",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "apiserver",
							Image:           "fahimabrar/api",
							ImagePullPolicy: "IfNotPresent",
							Ports: []corev1.ContainerPort{
								{
									Name:          "api-port",
									ContainerPort: 8080,
									Protocol:      "TCP",
								},
							},
						},
					},
					RestartPolicy: "Always",
				},
			},
		},
	}

	_ , _, err = kutilAppsV1.CreateOrPatchDeployment(
		kc,
		deployment.ObjectMeta,
		func(d *appsv1.Deployment) *appsv1.Deployment {
			d = deployment
			return d
		},
	)

	if err != nil {
		panic(err)
	}
	log.Debug("Deployment Created!")
}
/*func DeleteDeployment() {
	log.Debug("Deleting Deployment, Service and Ingress")

	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	kc := kubernetes.NewForConfigOrDie(config)
	varAppsV1 := kc.AppsV1()
	varCoreV1 := kc.CoreV1()
	varExtensionsV1Beta1 := kc.ExtensionsV1beta1()
	deletePolicy := metav1.DeletePropagationForeground
	if err := varAppsV1.Deployments("default").Delete("apiserver", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}

	if err = varCoreV1.Services("default").Delete("apiserver", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}

	if err = varExtensionsV1Beta1.Ingresses("default").Delete("ingress-apiserver", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	log.Debug("Deletion Competed Successfully!")

}*/

func DeleteDeploymentKutil(){
	log.Debug("Deleting Deployment, Service and Ingress")
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	kc := kubernetes.NewForConfigOrDie(config)
	varAppsV1 := kc.AppsV1()
	varCoreV1 := kc.CoreV1()
	varExtensionsV1Beta1 := kc.ExtensionsV1beta1()
	deletePolicy := metav1.DeletePropagationForeground

	deployment, err := varAppsV1.Deployments("default").Get("apiserver", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	if err := kutilAppsV1.DeleteDeployment(
		kc,
		deployment.ObjectMeta,
		); err !=nil{
		panic(err)
	}

	if err = varCoreV1.Services("default").Delete("apiserver", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}

	if err = varExtensionsV1Beta1.Ingresses("default").Delete("ingress-apiserver", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	log.Debug("Deletion Competed Successfully!")

}

func CreateService() {
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	kc := kubernetes.NewForConfigOrDie(config)
	varCoreV1 := kc.CoreV1()

	service := corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:   "apiserver",
			Labels: map[string]string{"app": "apiserver"},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:     "api-port",
					Protocol: "TCP",
					Port:     8080,
				},
			},
			Selector: map[string]string{
				"app": "apiserver",
			},
			Type: "NodePort",
		},
	}

	_, err = varCoreV1.Services("default").Create(&service)
	if err != nil {
		panic(err)
	}
}

func UpdateDeployment() {
	log.Debug("Scaling Up deployment from 1 to 5. . . . .")
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		log.Fatalf("Could not get Kubernetes config: %s", err)
	}
	kc := kubernetes.NewForConfigOrDie(config)
	varAppsV1 := kc.AppsV1()

	deployment, err := varAppsV1.Deployments("default").Get("apiserver", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	deployment.Spec.Replicas = int32Ptr(5)
	_, err = varAppsV1.Deployments("default").Update(deployment)

	if err != nil {
		panic(err)
	}
	log.Debug("Scaling Completed!")
}

func CreateIngress() {
	log.Debug("Creating Ingress. . . . .")
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}

	kc := kubernetes.NewForConfigOrDie(config)

	ingress := extv1beta1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: "ingress-apiserver",
		},
		Spec: extv1beta1.IngressSpec{
			Rules: []extv1beta1.IngressRule{
				{
					Host: "linkedin.local",
					IngressRuleValue: extv1beta1.IngressRuleValue{
						HTTP: &extv1beta1.HTTPIngressRuleValue{
							Paths: []extv1beta1.HTTPIngressPath{
								{
									Path: "/",
									Backend: extv1beta1.IngressBackend{
										ServiceName: "apiserver",
										ServicePort: intstr.FromString("api-port"),
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err = kc.ExtensionsV1beta1().Ingresses("default").Create(&ingress)

	if err != nil {
		panic(err)
	}
	log.Debug("Ingress Created!")

}

/*func GetURL(){
	kubeconfigPath := filepath.Join(os.Getenv("HOME"), ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		panic(err)
	}
	kc := kubernetes.NewForConfigOrDie(config)
	varCoreV1 := kc.CoreV1()

	node, err := varCoreV1.Nodes().Get("minikube", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", node)
}
*/

func int32Ptr(i int32) *int32 {
	return &i
}
