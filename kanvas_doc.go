// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
// DO NOT EDIT: this file is automatically generated by docgen
package kanvas

import (
	"github.com/projectdiscovery/yamldoc-go/encoder"
)

var (
	ComponentDoc  encoder.Doc
	DockerDoc     encoder.Doc
	TerraformDoc  encoder.Doc
	VarDoc        encoder.Doc
	KubernetesDoc encoder.Doc
)

func init() {
	ComponentDoc.Type = "Component"
	ComponentDoc.Comments[encoder.LineComment] = "Component is a component of the application"
	ComponentDoc.Description = "Component is a component of the application"
	ComponentDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Component",
			FieldName: "components",
		},
	}
	ComponentDoc.Fields = make([]encoder.Doc, 6)
	ComponentDoc.Fields[0].Name = "dir"
	ComponentDoc.Fields[0].Type = "string"
	ComponentDoc.Fields[0].Note = ""
	ComponentDoc.Fields[0].Description = "Dir is the directory to be chdir'ed before running the commands\nIf empty, this defaults to the base dir, which is where kanvas.yaml is located.\n"
	ComponentDoc.Fields[0].Comments[encoder.LineComment] = "Dir is the directory to be chdir'ed before running the commands"
	ComponentDoc.Fields[1].Name = "components"
	ComponentDoc.Fields[1].Type = "map[string]Component"
	ComponentDoc.Fields[1].Note = ""
	ComponentDoc.Fields[1].Description = "Components is a map of sub-components"
	ComponentDoc.Fields[1].Comments[encoder.LineComment] = "Components is a map of sub-components"
	ComponentDoc.Fields[2].Name = "needs"
	ComponentDoc.Fields[2].Type = "[]string"
	ComponentDoc.Fields[2].Note = ""
	ComponentDoc.Fields[2].Description = "Needs is a list of components that this component depends on"
	ComponentDoc.Fields[2].Comments[encoder.LineComment] = "Needs is a list of components that this component depends on"
	ComponentDoc.Fields[3].Name = "docker"
	ComponentDoc.Fields[3].Type = "Docker"
	ComponentDoc.Fields[3].Note = ""
	ComponentDoc.Fields[3].Description = "Docker is a docker-specific configuration"
	ComponentDoc.Fields[3].Comments[encoder.LineComment] = "Docker is a docker-specific configuration"
	ComponentDoc.Fields[4].Name = "terraform"
	ComponentDoc.Fields[4].Type = "Terraform"
	ComponentDoc.Fields[4].Note = ""
	ComponentDoc.Fields[4].Description = "Terraform is a terraform-specific configuration"
	ComponentDoc.Fields[4].Comments[encoder.LineComment] = "Terraform is a terraform-specific configuration"
	ComponentDoc.Fields[5].Name = "kubernetes"
	ComponentDoc.Fields[5].Type = "Kubernetes"
	ComponentDoc.Fields[5].Note = ""
	ComponentDoc.Fields[5].Description = "Kubernetes is a kubernetes-specific configuration"
	ComponentDoc.Fields[5].Comments[encoder.LineComment] = "Kubernetes is a kubernetes-specific configuration"

	DockerDoc.Type = "Docker"
	DockerDoc.Comments[encoder.LineComment] = "Docker is a docker-specific configuration"
	DockerDoc.Description = "Docker is a docker-specific configuration"
	DockerDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Component",
			FieldName: "docker",
		},
	}
	DockerDoc.Fields = make([]encoder.Doc, 2)
	DockerDoc.Fields[0].Name = "image"
	DockerDoc.Fields[0].Type = "string"
	DockerDoc.Fields[0].Note = ""
	DockerDoc.Fields[0].Description = "Image is the name of the image to be built"
	DockerDoc.Fields[0].Comments[encoder.LineComment] = "Image is the name of the image to be built"
	DockerDoc.Fields[1].Name = "file"
	DockerDoc.Fields[1].Type = "string"
	DockerDoc.Fields[1].Note = ""
	DockerDoc.Fields[1].Description = "File is the path to the Dockerfile"
	DockerDoc.Fields[1].Comments[encoder.LineComment] = "File is the path to the Dockerfile"

	TerraformDoc.Type = "Terraform"
	TerraformDoc.Comments[encoder.LineComment] = "Terraform is a terraform-specific configuration"
	TerraformDoc.Description = "Terraform is a terraform-specific configuration"
	TerraformDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Component",
			FieldName: "terraform",
		},
	}
	TerraformDoc.Fields = make([]encoder.Doc, 2)
	TerraformDoc.Fields[0].Name = "target"
	TerraformDoc.Fields[0].Type = "string"
	TerraformDoc.Fields[0].Note = ""
	TerraformDoc.Fields[0].Description = "Target is the target resource to be deployed"
	TerraformDoc.Fields[0].Comments[encoder.LineComment] = "Target is the target resource to be deployed"
	TerraformDoc.Fields[1].Name = "vars"
	TerraformDoc.Fields[1].Type = "[]Var"
	TerraformDoc.Fields[1].Note = ""
	TerraformDoc.Fields[1].Description = "Vars is a list of variables to be passed to terraform"
	TerraformDoc.Fields[1].Comments[encoder.LineComment] = "Vars is a list of variables to be passed to terraform"

	VarDoc.Type = "Var"
	VarDoc.Comments[encoder.LineComment] = "Var is a variable to be passed to terraform"
	VarDoc.Description = "Var is a variable to be passed to terraform"
	VarDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Terraform",
			FieldName: "vars",
		},
	}
	VarDoc.Fields = make([]encoder.Doc, 3)
	VarDoc.Fields[0].Name = "name"
	VarDoc.Fields[0].Type = "string"
	VarDoc.Fields[0].Note = ""
	VarDoc.Fields[0].Description = "Name is the name of the variable"
	VarDoc.Fields[0].Comments[encoder.LineComment] = "Name is the name of the variable"
	VarDoc.Fields[1].Name = "valueFrom"
	VarDoc.Fields[1].Type = "string"
	VarDoc.Fields[1].Note = ""
	VarDoc.Fields[1].Description = "ValueFrom is the source of the value of the variable"
	VarDoc.Fields[1].Comments[encoder.LineComment] = "ValueFrom is the source of the value of the variable"
	VarDoc.Fields[2].Name = "value"
	VarDoc.Fields[2].Type = "string"
	VarDoc.Fields[2].Note = ""
	VarDoc.Fields[2].Description = "Value is the value of the variable"
	VarDoc.Fields[2].Comments[encoder.LineComment] = "Value is the value of the variable"

	KubernetesDoc.Type = "Kubernetes"
	KubernetesDoc.Comments[encoder.LineComment] = "Kubernetes is a kubernetes-specific configuration"
	KubernetesDoc.Description = "Kubernetes is a kubernetes-specific configuration"
	KubernetesDoc.AppearsIn = []encoder.Appearance{
		{
			TypeName:  "Component",
			FieldName: "kubernetes",
		},
	}
	KubernetesDoc.Fields = make([]encoder.Doc, 0)
}

func (_ Component) Doc() *encoder.Doc {
	return &ComponentDoc
}

func (_ Docker) Doc() *encoder.Doc {
	return &DockerDoc
}

func (_ Terraform) Doc() *encoder.Doc {
	return &TerraformDoc
}

func (_ Var) Doc() *encoder.Doc {
	return &VarDoc
}

func (_ Kubernetes) Doc() *encoder.Doc {
	return &KubernetesDoc
}

// GetComponentDoc returns documentation for the file kanvas_doc.go.
func GetComponentDoc() *encoder.FileDoc {
	return &encoder.FileDoc{
		Name:        "Component",
		Description: "",
		Structs: []*encoder.Doc{
			&ComponentDoc,
			&DockerDoc,
			&TerraformDoc,
			&VarDoc,
			&KubernetesDoc,
		},
	}
}
