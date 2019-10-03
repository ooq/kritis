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

package testutil

import (
	"fmt"

	kritisv1beta1 "github.com/grafeas/kritis/pkg/kritis/apis/kritis/v1beta1"
	"github.com/grafeas/kritis/pkg/kritis/metadata"
	"github.com/grafeas/kritis/pkg/kritis/secrets"
	"google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/grafeas"
)

type MockMetadataClient struct {
	Vulnz           []metadata.Vulnerability
	PGPAttestations []metadata.PGPAttestation
	Occ             map[string]string
	Err             error
}

func (m *MockMetadataClient) SetError(err error) {
	m.Err = err
}

// Close does not do anything for MockMetadataClient
func (m *MockMetadataClient) Close() {
	// No ops
}

func (m *MockMetadataClient) Vulnerabilities(containerImage string) ([]metadata.Vulnerability, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.Vulnz, nil
}

func (m *MockMetadataClient) CreateAttestationOccurence(n *grafeas.Note, image string,
	s *secrets.PGPSigningSecret) (*grafeas.Occurrence, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if m.Occ == nil {
		m.Occ = map[string]string{}
	}
	m.Occ[fmt.Sprintf("%s-%s", image, n.Name)] = s.SecretName
	return nil, nil
}

func (m *MockMetadataClient) AttestationNote(aa *kritisv1beta1.Attestor) (*grafeas.Note, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	if aa == nil {
		return nil, fmt.Errorf("could not get note")
	}
	return &grafeas.Note{
		Name: aa.Spec.NoteReference,
	}, nil
}

func (m *MockMetadataClient) CreateAttestationNote(aa *kritisv1beta1.Attestor) (*grafeas.Note, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return &grafeas.Note{
		Name: aa.Spec.NoteReference,
	}, nil
}

func (m *MockMetadataClient) Attestations(containerImage string) ([]metadata.PGPAttestation, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return m.PGPAttestations, nil
}

func NilFetcher() func() (metadata.Fetcher, error) {
	return func() (metadata.Fetcher, error) {
		return &MockMetadataClient{
			Vulnz:           []metadata.Vulnerability{},
			PGPAttestations: []metadata.PGPAttestation{},
		}, nil
	}
}
