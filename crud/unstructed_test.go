package crud

import (
	"fmt"
	"reflect"
	"testing"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func TestUnstructedJson(t *testing.T) {
	u_json()
}

func TestUnstructedYaml(t *testing.T) {
	u_yaml()
}

func u_json() {
	uConfigMap := unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "v1",
			"kind":       "ConfigMap",
			"metadata": map[string]interface{}{
				"creationTimestamp": nil,
				"namespace":         "default",
				"name":              "my-configmap",
			},
			"data": map[string]interface{}{
				"foo": "bar",
			},
		},
	}
	bytes, err := runtime.Encode(unstructured.UnstructuredJSONScheme, &uConfigMap)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Serialized (option I)", string(bytes))

	// Unstructured -> JSON (Option II)
	//   - This is just a handy shortcut for the above code.
	bytes, err = uConfigMap.MarshalJSON()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Serialized (option II)", string(bytes))

	// JSON -> Unstructured (Option I)
	//   - Usage example: dynamic client (client-go/dynamic.Interface)
	obj1, err := runtime.Decode(unstructured.UnstructuredJSONScheme, bytes)
	if err != nil {
		panic(err.Error())
	}

	// JSON -> Unstructured (Option II)
	//   - This is just a handy shortcut for the above code.
	obj2 := &unstructured.Unstructured{}
	err = obj2.UnmarshalJSON(bytes)
	if err != nil {
		panic(err.Error())
	}
	if !reflect.DeepEqual(obj1, obj2) {
		panic("Unexpected configmap data")
	}
}

func u_yaml() {
	yConfigMap := `---
	apiVersion: v1
	data:
	  foo: bar
	kind: ConfigMap
	metadata:
	  creationTimestamp:
	  name: my-configmap
	  namespace: default
	`

	// YAML -> Unstructured (through JSON)
	jConfigMap, err := yaml.ToJSON([]byte(yConfigMap))
	if err != nil {
		panic(err.Error())
	}

	object, err := runtime.Decode(unstructured.UnstructuredJSONScheme, jConfigMap)
	if err != nil {
		panic(err.Error())
	}

	uConfigMap, ok := object.(*unstructured.Unstructured)
	if !ok {
		panic("unstructured.Unstructured expected")
	}

	if uConfigMap.GetName() != "my-configmap" {
		panic("Unexpected configmap data")
	}
}
