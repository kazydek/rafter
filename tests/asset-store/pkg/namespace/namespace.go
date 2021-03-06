package namespace

import (
	"github.com/pkg/errors"
	"k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

type Namespace struct {
	coreCli corev1.CoreV1Interface
	name    string
}

func New(coreCli corev1.CoreV1Interface, name string) *Namespace {
	return &Namespace{coreCli: coreCli, name: name}
}

func (n *Namespace) Create() error {
	_, err := n.coreCli.Namespaces().Create(&v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: n.name,
		},
	})

	if err != nil {
		if apierrors.IsAlreadyExists(err) {
			return nil
		}

		return errors.Wrapf(err, "while creating namespace %s", n.name)
	}

	return nil
}

func (n *Namespace) Delete() error {
	err := n.coreCli.Namespaces().Delete(n.name, &metav1.DeleteOptions{})
	if err != nil {
		return errors.Wrapf(err, "while deleting namespace %s", n.name)
	}

	return nil
}
