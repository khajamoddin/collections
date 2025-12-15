package main

import (
	"fmt"
	"log"

	"github.com/khajamoddin/collections/collections"
	"k8s.io/client-go/rest"
)

type cmKey string

func makeKey(ns, name string) cmKey {
	return cmKey(ns + "/" + name)
}

// ConfigView is a minimal projection of a ConfigMap for example purposes.
type ConfigView struct {
	Namespace string
	Name      string
	Labels    map[string]string
	DataKeys  []string
}

func main() {
	// Mock implementation for demo runnability without k8s cluster
	// In-cluster config would be: cfg, err := rest.InClusterConfig()
	_, _ = rest.InClusterConfig() // unused

	// Our in-memory indexes
	configs := collections.NewOrderedMap[cmKey, *ConfigView]()
	indexByApp := collections.NewMultiMap[string, cmKey]()

	// Simulation of Add/Update/Delete events
	simulateEvents := func() {
		// Add event
		cm1 := &ConfigView{Namespace: "default", Name: "app-config", Labels: map[string]string{"app": "my-app"}, DataKeys: []string{"cfg"}}
		key1 := makeKey(cm1.Namespace, cm1.Name)
		configs.Set(key1, cm1)
		indexByApp.Add("my-app", key1)
		log.Printf("[ADD] %s/%s (app=%s)", cm1.Namespace, cm1.Name, "my-app")

		// Add another
		cm2 := &ConfigView{Namespace: "default", Name: "db-config", Labels: map[string]string{"app": "my-app"}, DataKeys: []string{"host"}}
		key2 := makeKey(cm2.Namespace, cm2.Name)
		configs.Set(key2, cm2)
		indexByApp.Add("my-app", key2)
		log.Printf("[ADD] %s/%s (app=%s)", cm2.Namespace, cm2.Name, "my-app")

		// Update event (label change)
		cm1New := &ConfigView{Namespace: "default", Name: "app-config", Labels: map[string]string{"app": "legacy"}, DataKeys: []string{"cfg"}}
		// remove old index
		indexByApp.Remove("my-app", key1)
		// update map
		configs.Set(key1, cm1New)
		// add new index
		indexByApp.Add("legacy", key1)
		log.Printf("[UPDATE] %s/%s (app check logs)", cm1New.Namespace, cm1New.Name)
	}

	log.Println("Starting k8s-config-cache example...")
	simulateEvents()

	printByApp("my-app", configs, indexByApp)
	printByApp("legacy", configs, indexByApp)
}

func printByApp(app string,
	configs *collections.OrderedMap[cmKey, *ConfigView],
	index *collections.MultiMap[string, cmKey],
) {
	// MultiMap.Get returns []V but our MultiMap is [string, cmKey], so it returns []cmKey
	keys := index.Get(app)
	if len(keys) == 0 {
		log.Printf("No ConfigMaps found for app=%s", app)
		return
	}

	log.Printf("ConfigMaps for app=%s:", app)
	for _, key := range keys {
		if cm, ok := configs.Get(key); ok {
			fmt.Printf(" - %s/%s (keys=%v)\n", cm.Namespace, cm.Name, cm.DataKeys)
		}
	}
}
