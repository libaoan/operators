package samplectd_test

import (
	"github.com/operators/demo4crd/pkg/generated/clientset/versioned/typed/samplecrd/v1/fake"
	"github.com/operators/demo4crd/pkg/generated/informers/externalversions"
	"testing"
)

func SampleCrdWithFakeTest(t testing.T) {
	fakeClient := fake.FakeCrdV1{}
	factory := externalversions.NewSharedInformerFactory(fakeClient, 0)
	fooInformer := factory.Crd().V1().Foos()
	foo, err := fooInformer.Lister().Foos("default").Get("test")

}
