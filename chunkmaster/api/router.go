package api

import (
	"net/http"
)

var RouteMap map[string]map[string]http.HandlerFunc

func init() {
	RouteMap = map[string]map[string]http.HandlerFunc {
		"POST": {
			//"/v1/chunkserver/push"  : chunkserverPushHandler,
			"/v1/chunkserver/initserver" : initChunkserverHandler,
			//for monitor
			//"/v1/chunkserver/alarm" : monitorReportHandler,
			"/v1/chunkserver/reloadinfo" : loadChunkserverInfoHandler,
			"/v1/chunkserver/reportinfo" : reportChunkserverInfoHandler,
		},
		"GET": {
			"/v1/chunkmaster/route" : chunkmasterRouteHandler,
			"/v1/chunkmaster/fid"	: chunkmasterFidHandler,

			"/v1/chunkserver/checkerror" : chunkserverCheckError,
		},
	}
}
