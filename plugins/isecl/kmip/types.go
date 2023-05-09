/*
Copyright 2021-2022.
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

package kmip

import (
	"github.com/gemalto/kmip-go"
	"github.com/gemalto/kmip-go/kmip20"
	"github.com/gemalto/kmip-go/ttlv"
)

type ObjectInfo struct {
	ID    string
	Label string
	Type  string
}

type GetAttributesRequestPayload struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
	Attributes       ttlv.Values
}
type GetAttributes struct {
	Name       kmip.Name
	ObjectType kmip20.ObjectType
}
type GetAttributesResponsePayload struct {
	UniqueIdentifier string
	Attribute        []kmip.Attribute // KMIP 1.4
	Attributes       *GetAttributes   // KMIP 2.0
}

type RegisterRequestPayload struct {
	ObjectType        kmip20.ObjectType
	TemplateAttribute []kmip.TemplateAttribute // KMIP 1.4
	Attributes        ttlv.Value               // KMIP 2.0
	Certificate       *kmip.Certificate
	PrivateKey        *kmip.PrivateKey
}

type GetResponsePayload struct {
	ObjectType       kmip20.ObjectType
	UniqueIdentifier string
	Certificate      kmip.Certificate
	PublicKey        kmip.PublicKey
	PrivateKey       kmip.PrivateKey
}

type LocateAttributes struct {
	ObjectType string
	Name       *kmip.Name
}
type LocateRequestPayload struct {
	Attribute  []kmip.Attribute // KMIP 1.4
	Attributes ttlv.Value       // KMIP 2.0
}
type LocateResponsePayload struct {
	UniqueIdentifier []string
}

type DeleteRequestPayload struct {
	UniqueIdentifier kmip20.UniqueIdentifierValue
}
