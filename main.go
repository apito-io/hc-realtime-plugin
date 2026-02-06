package main

import (
	"context"
	"log"

	sdk "github.com/apito-io/go-apito-plugin-sdk"
)

func main() {
	log.Printf("🎯 [hc-realtime-plugin] Starting Realtime plugin...")
	plugin := sdk.Init("hc-realtime-plugin", "1.0.0", "apito-plugin-key")
	statusType := sdk.NewObjectType("RealtimeStatus", "Realtime plugin status").
		AddStringField("status", "Plugin status", false).
		AddStringField("version", "Plugin version", false).
		Build()
	plugin.RegisterQuery("getRealtimeStatus",
		sdk.ComplexObjectField("Get realtime plugin status", statusType),
		func(ctx context.Context, rawArgs map[string]interface{}) (interface{}, error) {
			return map[string]interface{}{"status": "ready", "version": "1.0.0"}, nil
		})
	log.Printf("🚀 [hc-realtime-plugin] Plugin ready")
	plugin.Serve()
}
