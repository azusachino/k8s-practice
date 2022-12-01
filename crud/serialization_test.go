package crud

import (
	"encoding/json"
	"fmt"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	jsonserializer "k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
)

func TestJson(t *testing.T) {
	j_son()
}

func TestYaml(t *testing.T) {
	y_aml()
}

func j_son() {
	obj := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo-bar-config",
			Namespace: "default",
		},
		Data: map[string]string{"foo": "bar"},
	}

	encoder := jsonserializer.NewSerializerWithOptions(
		nil, // jsonserializer.MetaFactory
		nil, // runtime.ObjectCreater
		nil, // runtime.ObjectTyper
		jsonserializer.SerializerOptions{
			Yaml:   false,
			Pretty: false,
			Strict: false,
		},
	)

	// k8s 运行时提供的序列化方法
	encoded, err := runtime.Encode(encoder, &obj)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Serialized (option I)", string(encoded))

	// 直接使用 官方json库
	encoded2, err := json.Marshal(obj)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Serialized (option II)", string(encoded2))

	decoder := jsonserializer.NewSerializerWithOptions(
		jsonserializer.DefaultMetaFactory, // jsonserializer.MetaFactory
		scheme.Scheme,                     // runtime.Scheme implements runtime.ObjectCreater
		scheme.Scheme,                     // runtime.Scheme implements runtime.ObjectTyper
		jsonserializer.SerializerOptions{
			Yaml:   false,
			Pretty: false,
			Strict: false,
		},
	)

	// The actual decoding is much like stdlib encoding/json.Unmarshal but with some
	// minor tweaks - see https://github.com/kubernetes-sigs/json for more.
	// 运行时提供的反序列方法
	decoded, err := runtime.Decode(decoder, encoded)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Deserialized %#v\n", decoded)
}

func y_aml() {
	obj := corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		Data: map[string]string{"foo": "bar"},
	}
	obj.Namespace = "default"
	obj.Name = "my-configmap"

	// Serializer = Decoder + Encoder.
	serializer := jsonserializer.NewSerializerWithOptions(
		jsonserializer.DefaultMetaFactory, // jsonserializer.MetaFactory
		scheme.Scheme,                     // runtime.Scheme implements runtime.ObjectCreater
		scheme.Scheme,                     // runtime.Scheme implements runtime.ObjectTyper
		jsonserializer.SerializerOptions{
			Yaml:   true,
			Pretty: false,
			Strict: false,
		},
	)

	// Typed -> YAML
	// Runtime.Encode() is just a helper function to invoke Encoder.Encode()
	yaml, err := runtime.Encode(serializer, &obj)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Serialized:\n%s", string(yaml))

	// YAML -> Typed (through JSON, actually)
	decoded, err := runtime.Decode(serializer, yaml)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Deserialized: %#v\n", decoded)
}
