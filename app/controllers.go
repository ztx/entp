// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/ztx/entp/designsvc
// --out=$(GOPATH)/src/github.com/ztx/entp
// --version=v1.1.0-dirty
//
// API "entp": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// PrLineController is the controller interface for the PrLine actions.
type PrLineController interface {
	goa.Muxer
	Aprove(*AprovePrLineContext) error
	Create(*CreatePrLineContext) error
	List(*ListPrLineContext) error
	Show(*ShowPrLineContext) error
}

// MountPrLineController "mounts" a PrLine resource controller on the given service.
func MountPrLineController(service *goa.Service, ctrl PrLineController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/entp/prs/:prNum/prlines/aprove/:qty", ctrl.MuxHandler("preflight", handlePrLineOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/entp/prs/:prNum/prlines", ctrl.MuxHandler("preflight", handlePrLineOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/entp/prs/:prNum/prlines/:prLineNum", ctrl.MuxHandler("preflight", handlePrLineOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAprovePrLineContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AprovePrLinePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Aprove(rctx)
	}
	h = handlePrLineOrigin(h)
	service.Mux.Handle("POST", "/entp/prs/:prNum/prlines/aprove/:qty", ctrl.MuxHandler("Aprove", h, unmarshalAprovePrLinePayload))
	service.LogInfo("mount", "ctrl", "PrLine", "action", "Aprove", "route", "POST /entp/prs/:prNum/prlines/aprove/:qty")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePrLineContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreatePrLinePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handlePrLineOrigin(h)
	service.Mux.Handle("POST", "/entp/prs/:prNum/prlines", ctrl.MuxHandler("Create", h, unmarshalCreatePrLinePayload))
	service.LogInfo("mount", "ctrl", "PrLine", "action", "Create", "route", "POST /entp/prs/:prNum/prlines")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPrLineContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handlePrLineOrigin(h)
	service.Mux.Handle("GET", "/entp/prs/:prNum/prlines", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "PrLine", "action", "List", "route", "GET /entp/prs/:prNum/prlines")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPrLineContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handlePrLineOrigin(h)
	service.Mux.Handle("GET", "/entp/prs/:prNum/prlines/:prLineNum", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "PrLine", "action", "Show", "route", "GET /entp/prs/:prNum/prlines/:prLineNum")
}

// handlePrLineOrigin applies the CORS response headers corresponding to the origin.
func handlePrLineOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAprovePrLinePayload unmarshals the request body into the context request data Payload field.
func unmarshalAprovePrLinePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &aprovePrLinePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreatePrLinePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePrLinePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createPrLinePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// ItemController is the controller interface for the Item actions.
type ItemController interface {
	goa.Muxer
	Create(*CreateItemContext) error
	List(*ListItemContext) error
	Show(*ShowItemContext) error
}

// MountItemController "mounts" a Item resource controller on the given service.
func MountItemController(service *goa.Service, ctrl ItemController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/entp/items", ctrl.MuxHandler("preflight", handleItemOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/entp/items/:itemCode", ctrl.MuxHandler("preflight", handleItemOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateItemContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateItemPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleItemOrigin(h)
	service.Mux.Handle("POST", "/entp/items", ctrl.MuxHandler("Create", h, unmarshalCreateItemPayload))
	service.LogInfo("mount", "ctrl", "Item", "action", "Create", "route", "POST /entp/items")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListItemContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleItemOrigin(h)
	service.Mux.Handle("GET", "/entp/items", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Item", "action", "List", "route", "GET /entp/items")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowItemContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handleItemOrigin(h)
	service.Mux.Handle("GET", "/entp/items/:itemCode", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Item", "action", "Show", "route", "GET /entp/items/:itemCode")
}

// handleItemOrigin applies the CORS response headers corresponding to the origin.
func handleItemOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateItemPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateItemPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createItemPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PrController is the controller interface for the Pr actions.
type PrController interface {
	goa.Muxer
	AddLine(*AddLinePrContext) error
	Create(*CreatePrContext) error
	List(*ListPrContext) error
	Show(*ShowPrContext) error
}

// MountPrController "mounts" a Pr resource controller on the given service.
func MountPrController(service *goa.Service, ctrl PrController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/entp/prs/addline", ctrl.MuxHandler("preflight", handlePrOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/entp/prs", ctrl.MuxHandler("preflight", handlePrOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/entp/prs/:prNum", ctrl.MuxHandler("preflight", handlePrOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAddLinePrContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*AddLinePrPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.AddLine(rctx)
	}
	h = handlePrOrigin(h)
	service.Mux.Handle("POST", "/entp/prs/addline", ctrl.MuxHandler("AddLine", h, unmarshalAddLinePrPayload))
	service.LogInfo("mount", "ctrl", "Pr", "action", "AddLine", "route", "POST /entp/prs/addline")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreatePrContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreatePrPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handlePrOrigin(h)
	service.Mux.Handle("POST", "/entp/prs", ctrl.MuxHandler("Create", h, unmarshalCreatePrPayload))
	service.LogInfo("mount", "ctrl", "Pr", "action", "Create", "route", "POST /entp/prs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListPrContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handlePrOrigin(h)
	service.Mux.Handle("GET", "/entp/prs", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Pr", "action", "List", "route", "GET /entp/prs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowPrContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	h = handlePrOrigin(h)
	service.Mux.Handle("GET", "/entp/prs/:prNum", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Pr", "action", "Show", "route", "GET /entp/prs/:prNum")
}

// handlePrOrigin applies the CORS response headers corresponding to the origin.
func handlePrOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalAddLinePrPayload unmarshals the request body into the context request data Payload field.
func unmarshalAddLinePrPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &addLinePrPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreatePrPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreatePrPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createPrPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
