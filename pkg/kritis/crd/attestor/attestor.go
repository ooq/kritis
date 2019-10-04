/*
Copyright 2018 Google LLC

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

package attestor

import (
	"fmt"

	"github.com/grafeas/kritis/pkg/kritis/apis/kritis/v1beta1"
	clientset "github.com/grafeas/kritis/pkg/kritis/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

type Lister func(namespace string) ([]v1beta1.Attestor, error)
type Fetcher func(namespace string, name string) (*v1beta1.Attestor, error)

// Attestors returns all Attestors in the specified namespaces
// Pass in an empty string to get all Attestors in all namespaces
func Attestors(namespace string) ([]v1beta1.Attestor, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("error building config: %v", err)
	}

	client, err := clientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building clientset: %v", err)
	}
	list, err := client.KritisV1beta1().Attestors(namespace).List(metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("error listing all attestors: %v", err)
	}
	return list.Items, nil
}

// Attestor returns the Attestor in the specified namespace and with the given name
// Returns error if Attestor is not found
func Attestor(namespace string, name string) (*v1beta1.Attestor, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("error building config: %v", err)
	}

	client, err := clientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("error building clientset: %v", err)
	}
	return client.KritisV1beta1().Attestors(namespace).Get(name, metav1.GetOptions{})
}
