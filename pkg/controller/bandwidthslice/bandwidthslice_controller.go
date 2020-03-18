package bandwidthslice

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	bansv1alpha1 "github.com/stevenchiu30801/onos-bandwidth-operator/pkg/apis/bans/v1alpha1"
	helm "github.com/stevenchiu30801/onos-bandwidth-operator/pkg/helm"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var reqLogger = logf.Log.WithName("controller_bandwidthslice")

const (
	ONOS_USERNAME string = "onos"
	ONOS_PASSWORD string = "rocks"
)

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new BandwidthSlice Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileBandwidthSlice{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("bandwidthslice-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource BandwidthSlice
	err = c.Watch(&source.Kind{Type: &bansv1alpha1.BandwidthSlice{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner BandwidthSlice
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &bansv1alpha1.BandwidthSlice{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileBandwidthSlice implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileBandwidthSlice{}

// ReconcileBandwidthSlice reconciles a BandwidthSlice object
type ReconcileBandwidthSlice struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a BandwidthSlice object and makes changes based on the state read
// and what is in the BandwidthSlice.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileBandwidthSlice) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger.Info("Reconciling BandwidthSlice")

	// Fetch the BandwidthSlice instance
	instance := &bansv1alpha1.BandwidthSlice{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Check if ONOS already exists, if not create a new one
	onos := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "onos", Namespace: instance.Namespace}, onos)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating ONOS", "Namespace", instance.Namespace, "Name", "onos")

		// Create ONOS bandwidth management application Helm values
		vals := map[string]interface{}{
			"bandwidthManagement": true,
			"env": []map[string]interface{}{
				{
					"name":  "ONOS_APPS",
					"value": "bandwidth-management,drivers,drivers.bmv2,fwd,hostprovider,proxyarp",
				},
			},
		}

		err = helm.InstallHelmChart(instance.Namespace, "onos", "onos", vals)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
		// ONOS already exists
		reqLogger.Info("ONOS already exists", "Namespace", onos.Namespace, "Name", onos.Name)
	}

	// Wait for ONOS to be ready
	client := &http.Client{}
	for {
		// Verfiy state of org.onosproject.drivers.bmv2 application
		req, err := http.NewRequest("GET",
			"http://onos-gui.default.svc.cluster.local:8181/onos/v1/applications/org.onosproject.drivers.bmv2",
			nil)
		if err != nil {
			return reconcile.Result{}, err
		}
		req.SetBasicAuth(ONOS_USERNAME, ONOS_PASSWORD)
		resp, err := client.Do(req)

		if resp.StatusCode == http.StatusOK {
			// Read response from ONOS
			defer resp.Body.Close()
			buf, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return reconcile.Result{}, err
			}
			var decoded map[string]interface{}
			err = json.Unmarshal(buf, &decoded)
			if err != nil {
				return reconcile.Result{}, err
			}
			if decoded["state"] == "ACTIVE" {
				break
			}
		}
		reqLogger.Info("Waiting 3 seconds for ONOS to be ready", "Namespace", onos.Namespace, "Name", onos.Name)
		time.Sleep(3 * time.Second)
	}

	// Check if Mininet already exists, if not create a new one
	mininet := &appsv1.Deployment{}
	err = r.client.Get(context.TODO(), types.NamespacedName{Name: "mininet", Namespace: instance.Namespace}, mininet)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating Mininet", "Namespace", instance.Namespace, "Name", "mininet")

		err = helm.InstallHelmChart(instance.Namespace, "mininet", "mininet", nil)
		if err != nil {
			return reconcile.Result{}, err
		}
	} else if err != nil {
		return reconcile.Result{}, err
	} else {
		// Mininet already exists
		reqLogger.Info("Mininet already exists", "Namespace", mininet.Namespace, "Name", mininet.Name)
	}

	// Send bandwidth slice request to ONOS
	buf, err := json.Marshal(instance.Spec)
	if err != nil {
		// Malformed BandwidthSliceSpec
		// Return and don't requeue
		reqLogger.Error(err, "Cannot marshal BandwidthSliceSpec to JSON format")
		return reconcile.Result{}, nil
	}
	reqLogger.Info("Sending bandwidth slice request", "Body", string(buf))
	req, err := http.NewRequest("POST",
		"http://onos-gui.default.svc.cluster.local:8181/onos/bandwidth-management/slices",
		bytes.NewReader(buf))
	if err != nil {
		return reconcile.Result{}, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(ONOS_USERNAME, ONOS_PASSWORD)
	resp, err := client.Do(req)

	// Read response from ONOS
	defer resp.Body.Close()
	buf, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return reconcile.Result{}, err
	}
	reqLogger.Info("Successfully sent bandwidth slice request", "Response", string(buf))

	return reconcile.Result{}, nil
}
