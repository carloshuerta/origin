/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BrokerLister helps list Brokers.
type BrokerLister interface {
	// List lists all Brokers in the indexer.
	List(selector labels.Selector) (ret []*servicecatalog.Broker, err error)
	// Get retrieves the Broker from the index for a given name.
	Get(name string) (*servicecatalog.Broker, error)
	BrokerListerExpansion
}

// brokerLister implements the BrokerLister interface.
type brokerLister struct {
	indexer cache.Indexer
}

// NewBrokerLister returns a new BrokerLister.
func NewBrokerLister(indexer cache.Indexer) BrokerLister {
	return &brokerLister{indexer: indexer}
}

// List lists all Brokers in the indexer.
func (s *brokerLister) List(selector labels.Selector) (ret []*servicecatalog.Broker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.Broker))
	})
	return ret, err
}

// Get retrieves the Broker from the index for a given name.
func (s *brokerLister) Get(name string) (*servicecatalog.Broker, error) {
	key := &servicecatalog.Broker{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("broker"), name)
	}
	return obj.(*servicecatalog.Broker), nil
}