/*
Copyright 2023 The AAQ Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubevirt.io/application-aware-quota/staging/src/kubevirt.io/application-aware-quota-api/pkg/apis/core/v1alpha1"
)

// AAQJobQueueConfigLister helps list AAQJobQueueConfigs.
// All objects returned here must be treated as read-only.
type AAQJobQueueConfigLister interface {
	// List lists all AAQJobQueueConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AAQJobQueueConfig, err error)
	// AAQJobQueueConfigs returns an object that can list and get AAQJobQueueConfigs.
	AAQJobQueueConfigs(namespace string) AAQJobQueueConfigNamespaceLister
	AAQJobQueueConfigListerExpansion
}

// aAQJobQueueConfigLister implements the AAQJobQueueConfigLister interface.
type aAQJobQueueConfigLister struct {
	indexer cache.Indexer
}

// NewAAQJobQueueConfigLister returns a new AAQJobQueueConfigLister.
func NewAAQJobQueueConfigLister(indexer cache.Indexer) AAQJobQueueConfigLister {
	return &aAQJobQueueConfigLister{indexer: indexer}
}

// List lists all AAQJobQueueConfigs in the indexer.
func (s *aAQJobQueueConfigLister) List(selector labels.Selector) (ret []*v1alpha1.AAQJobQueueConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AAQJobQueueConfig))
	})
	return ret, err
}

// AAQJobQueueConfigs returns an object that can list and get AAQJobQueueConfigs.
func (s *aAQJobQueueConfigLister) AAQJobQueueConfigs(namespace string) AAQJobQueueConfigNamespaceLister {
	return aAQJobQueueConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AAQJobQueueConfigNamespaceLister helps list and get AAQJobQueueConfigs.
// All objects returned here must be treated as read-only.
type AAQJobQueueConfigNamespaceLister interface {
	// List lists all AAQJobQueueConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AAQJobQueueConfig, err error)
	// Get retrieves the AAQJobQueueConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AAQJobQueueConfig, error)
	AAQJobQueueConfigNamespaceListerExpansion
}

// aAQJobQueueConfigNamespaceLister implements the AAQJobQueueConfigNamespaceLister
// interface.
type aAQJobQueueConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AAQJobQueueConfigs in the indexer for a given namespace.
func (s aAQJobQueueConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AAQJobQueueConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AAQJobQueueConfig))
	})
	return ret, err
}

// Get retrieves the AAQJobQueueConfig from the indexer for a given namespace and name.
func (s aAQJobQueueConfigNamespaceLister) Get(name string) (*v1alpha1.AAQJobQueueConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("aaqjobqueueconfig"), name)
	}
	return obj.(*v1alpha1.AAQJobQueueConfig), nil
}
