package service

import (
	"context"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

/**
DIT IS TEST CODE
*/

func RunClientFunction() {
	// Setup client
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// `clientset` can now be used to create, edit and read resources

	printNodes(*clientset)
}

/*
*
Chain of operations to deploy simple html page
*/
func deployHtml(clientset kubernetes.Clientset) {
	fmt.Println("Deploying 'Hello world' page...")

	createHtmlConfigMap(clientset)
	createDeployment(clientset)
	createService(clientset)

	fmt.Println("Deployed 'Hello world' page.")
}

/*
*
Test function to create a service
*/
func createService(clientset kubernetes.Clientset) {
	fmt.Println("Creating service for 'Hello world' page...")
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "html-service",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "nginx",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       80,
					TargetPort: intstr.FromInt(80),
				},
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
	clientset.CoreV1().Services("default").Create(context.Background(), service, metav1.CreateOptions{})
}

/*
*
Test function to create a deployment
*/
func createDeployment(clientset kubernetes.Clientset) {
	fmt.Println("Deploying 'Hello world' html page...")

	deployment := appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "html-deployment",
			Namespace: "default",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(5),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "nginx",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "nginx",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "nginx",
							Image: "nginx:latest",
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "nginx-index-file",
									MountPath: "/usr/share/nginx/html/",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "nginx-index-file",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "html-test-config",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	clientset.AppsV1().Deployments("default").Create(context.Background(), &deployment, metav1.CreateOptions{})
}

/*
*
Test function to create a config map containing a simple 'Hello world' html page
*/
func createHtmlConfigMap(clientset kubernetes.Clientset) {
	fmt.Println("Creating 'Hello world' html config...")
	configMap := corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "html-test-config",
			Namespace: "default",
		},
		Data: map[string]string{
			"index.html": `<html><h1>Hello world!</h1></html`,
		},
	}
	clientset.CoreV1().ConfigMaps("default").Create(context.Background(), &configMap, metav1.CreateOptions{})
}

/*
*
Test function to print running nodes
*/
func printNodes(clientset kubernetes.Clientset) {
	fmt.Println("Fetching current nodes...")
	nodes, _ := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	fmt.Println("Nodes:\n")
	for _, node := range nodes.Items {
		fmt.Printf("\t%s\n", node.Name)
	}
}

/*
*
Parse function
*/
func int32Ptr(i int32) *int32 {
	return &i
}
