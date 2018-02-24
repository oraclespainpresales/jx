// Package target provides the Chrome Debugging Protocol
// commands, types, and events for the Target domain.
//
// Supports additional targets discovery and allows to attach to them.
//
// Generated by the cdproto-gen command.
package target

// Code generated by cdproto-gen. DO NOT EDIT.

import (
	"context"

	"github.com/chromedp/cdproto/cdp"
)

// ActivateTargetParams activates (focuses) the target.
type ActivateTargetParams struct {
	TargetID ID `json:"targetId"`
}

// ActivateTarget activates (focuses) the target.
//
// parameters:
//   targetID
func ActivateTarget(targetID ID) *ActivateTargetParams {
	return &ActivateTargetParams{
		TargetID: targetID,
	}
}

// Do executes Target.activateTarget against the provided context.
func (p *ActivateTargetParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandActivateTarget, p, nil)
}

// AttachToTargetParams attaches to the target with given id.
type AttachToTargetParams struct {
	TargetID ID `json:"targetId"`
}

// AttachToTarget attaches to the target with given id.
//
// parameters:
//   targetID
func AttachToTarget(targetID ID) *AttachToTargetParams {
	return &AttachToTargetParams{
		TargetID: targetID,
	}
}

// AttachToTargetReturns return values.
type AttachToTargetReturns struct {
	SessionID SessionID `json:"sessionId,omitempty"` // Id assigned to the session.
}

// Do executes Target.attachToTarget against the provided context.
//
// returns:
//   sessionID - Id assigned to the session.
func (p *AttachToTargetParams) Do(ctxt context.Context, h cdp.Executor) (sessionID SessionID, err error) {
	// execute
	var res AttachToTargetReturns
	err = h.Execute(ctxt, CommandAttachToTarget, p, &res)
	if err != nil {
		return "", err
	}

	return res.SessionID, nil
}

// CloseTargetParams closes the target. If the target is a page that gets
// closed too.
type CloseTargetParams struct {
	TargetID ID `json:"targetId"`
}

// CloseTarget closes the target. If the target is a page that gets closed
// too.
//
// parameters:
//   targetID
func CloseTarget(targetID ID) *CloseTargetParams {
	return &CloseTargetParams{
		TargetID: targetID,
	}
}

// CloseTargetReturns return values.
type CloseTargetReturns struct {
	Success bool `json:"success,omitempty"`
}

// Do executes Target.closeTarget against the provided context.
//
// returns:
//   success
func (p *CloseTargetParams) Do(ctxt context.Context, h cdp.Executor) (success bool, err error) {
	// execute
	var res CloseTargetReturns
	err = h.Execute(ctxt, CommandCloseTarget, p, &res)
	if err != nil {
		return false, err
	}

	return res.Success, nil
}

// CreateBrowserContextParams creates a new empty BrowserContext. Similar to
// an incognito profile but you can have more than one.
type CreateBrowserContextParams struct{}

// CreateBrowserContext creates a new empty BrowserContext. Similar to an
// incognito profile but you can have more than one.
func CreateBrowserContext() *CreateBrowserContextParams {
	return &CreateBrowserContextParams{}
}

// CreateBrowserContextReturns return values.
type CreateBrowserContextReturns struct {
	BrowserContextID BrowserContextID `json:"browserContextId,omitempty"` // The id of the context created.
}

// Do executes Target.createBrowserContext against the provided context.
//
// returns:
//   browserContextID - The id of the context created.
func (p *CreateBrowserContextParams) Do(ctxt context.Context, h cdp.Executor) (browserContextID BrowserContextID, err error) {
	// execute
	var res CreateBrowserContextReturns
	err = h.Execute(ctxt, CommandCreateBrowserContext, nil, &res)
	if err != nil {
		return "", err
	}

	return res.BrowserContextID, nil
}

// CreateTargetParams creates a new page.
type CreateTargetParams struct {
	URL                     string           `json:"url"`                               // The initial URL the page will be navigated to.
	Width                   int64            `json:"width,omitempty"`                   // Frame width in DIP (headless chrome only).
	Height                  int64            `json:"height,omitempty"`                  // Frame height in DIP (headless chrome only).
	BrowserContextID        BrowserContextID `json:"browserContextId,omitempty"`        // The browser context to create the page in (headless chrome only).
	EnableBeginFrameControl bool             `json:"enableBeginFrameControl,omitempty"` // Whether BeginFrames for this target will be controlled via DevTools (headless chrome only, not supported on MacOS yet, false by default).
}

// CreateTarget creates a new page.
//
// parameters:
//   url - The initial URL the page will be navigated to.
func CreateTarget(url string) *CreateTargetParams {
	return &CreateTargetParams{
		URL: url,
	}
}

// WithWidth frame width in DIP (headless chrome only).
func (p CreateTargetParams) WithWidth(width int64) *CreateTargetParams {
	p.Width = width
	return &p
}

// WithHeight frame height in DIP (headless chrome only).
func (p CreateTargetParams) WithHeight(height int64) *CreateTargetParams {
	p.Height = height
	return &p
}

// WithBrowserContextID the browser context to create the page in (headless
// chrome only).
func (p CreateTargetParams) WithBrowserContextID(browserContextID BrowserContextID) *CreateTargetParams {
	p.BrowserContextID = browserContextID
	return &p
}

// WithEnableBeginFrameControl whether BeginFrames for this target will be
// controlled via DevTools (headless chrome only, not supported on MacOS yet,
// false by default).
func (p CreateTargetParams) WithEnableBeginFrameControl(enableBeginFrameControl bool) *CreateTargetParams {
	p.EnableBeginFrameControl = enableBeginFrameControl
	return &p
}

// CreateTargetReturns return values.
type CreateTargetReturns struct {
	TargetID ID `json:"targetId,omitempty"` // The id of the page opened.
}

// Do executes Target.createTarget against the provided context.
//
// returns:
//   targetID - The id of the page opened.
func (p *CreateTargetParams) Do(ctxt context.Context, h cdp.Executor) (targetID ID, err error) {
	// execute
	var res CreateTargetReturns
	err = h.Execute(ctxt, CommandCreateTarget, p, &res)
	if err != nil {
		return "", err
	}

	return res.TargetID, nil
}

// DetachFromTargetParams detaches session with given id.
type DetachFromTargetParams struct {
	SessionID SessionID `json:"sessionId,omitempty"` // Session to detach.
}

// DetachFromTarget detaches session with given id.
//
// parameters:
func DetachFromTarget() *DetachFromTargetParams {
	return &DetachFromTargetParams{}
}

// WithSessionID session to detach.
func (p DetachFromTargetParams) WithSessionID(sessionID SessionID) *DetachFromTargetParams {
	p.SessionID = sessionID
	return &p
}

// Do executes Target.detachFromTarget against the provided context.
func (p *DetachFromTargetParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDetachFromTarget, p, nil)
}

// DisposeBrowserContextParams deletes a BrowserContext, will fail of any
// open page uses it.
type DisposeBrowserContextParams struct {
	BrowserContextID BrowserContextID `json:"browserContextId"`
}

// DisposeBrowserContext deletes a BrowserContext, will fail of any open page
// uses it.
//
// parameters:
//   browserContextID
func DisposeBrowserContext(browserContextID BrowserContextID) *DisposeBrowserContextParams {
	return &DisposeBrowserContextParams{
		BrowserContextID: browserContextID,
	}
}

// DisposeBrowserContextReturns return values.
type DisposeBrowserContextReturns struct {
	Success bool `json:"success,omitempty"`
}

// Do executes Target.disposeBrowserContext against the provided context.
//
// returns:
//   success
func (p *DisposeBrowserContextParams) Do(ctxt context.Context, h cdp.Executor) (success bool, err error) {
	// execute
	var res DisposeBrowserContextReturns
	err = h.Execute(ctxt, CommandDisposeBrowserContext, p, &res)
	if err != nil {
		return false, err
	}

	return res.Success, nil
}

// GetTargetInfoParams returns information about a target.
type GetTargetInfoParams struct {
	TargetID ID `json:"targetId"`
}

// GetTargetInfo returns information about a target.
//
// parameters:
//   targetID
func GetTargetInfo(targetID ID) *GetTargetInfoParams {
	return &GetTargetInfoParams{
		TargetID: targetID,
	}
}

// GetTargetInfoReturns return values.
type GetTargetInfoReturns struct {
	TargetInfo *Info `json:"targetInfo,omitempty"`
}

// Do executes Target.getTargetInfo against the provided context.
//
// returns:
//   targetInfo
func (p *GetTargetInfoParams) Do(ctxt context.Context, h cdp.Executor) (targetInfo *Info, err error) {
	// execute
	var res GetTargetInfoReturns
	err = h.Execute(ctxt, CommandGetTargetInfo, p, &res)
	if err != nil {
		return nil, err
	}

	return res.TargetInfo, nil
}

// GetTargetsParams retrieves a list of available targets.
type GetTargetsParams struct{}

// GetTargets retrieves a list of available targets.
func GetTargets() *GetTargetsParams {
	return &GetTargetsParams{}
}

// GetTargetsReturns return values.
type GetTargetsReturns struct {
	TargetInfos []*Info `json:"targetInfos,omitempty"` // The list of targets.
}

// Do executes Target.getTargets against the provided context.
//
// returns:
//   targetInfos - The list of targets.
func (p *GetTargetsParams) Do(ctxt context.Context, h cdp.Executor) (targetInfos []*Info, err error) {
	// execute
	var res GetTargetsReturns
	err = h.Execute(ctxt, CommandGetTargets, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.TargetInfos, nil
}

// SendMessageToTargetParams sends protocol message over session with given
// id.
type SendMessageToTargetParams struct {
	Message   string    `json:"message"`
	SessionID SessionID `json:"sessionId,omitempty"` // Identifier of the session.
}

// SendMessageToTarget sends protocol message over session with given id.
//
// parameters:
//   message
func SendMessageToTarget(message string) *SendMessageToTargetParams {
	return &SendMessageToTargetParams{
		Message: message,
	}
}

// WithSessionID identifier of the session.
func (p SendMessageToTargetParams) WithSessionID(sessionID SessionID) *SendMessageToTargetParams {
	p.SessionID = sessionID
	return &p
}

// Do executes Target.sendMessageToTarget against the provided context.
func (p *SendMessageToTargetParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSendMessageToTarget, p, nil)
}

// SetAutoAttachParams controls whether to automatically attach to new
// targets which are considered to be related to this one. When turned on,
// attaches to all existing related targets as well. When turned off,
// automatically detaches from all currently attached targets.
type SetAutoAttachParams struct {
	AutoAttach             bool `json:"autoAttach"`             // Whether to auto-attach to related targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"` // Whether to pause new targets when attaching to them. Use Runtime.runIfWaitingForDebugger to run paused targets.
}

// SetAutoAttach controls whether to automatically attach to new targets
// which are considered to be related to this one. When turned on, attaches to
// all existing related targets as well. When turned off, automatically detaches
// from all currently attached targets.
//
// parameters:
//   autoAttach - Whether to auto-attach to related targets.
//   waitForDebuggerOnStart - Whether to pause new targets when attaching to them. Use Runtime.runIfWaitingForDebugger to run paused targets.
func SetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool) *SetAutoAttachParams {
	return &SetAutoAttachParams{
		AutoAttach:             autoAttach,
		WaitForDebuggerOnStart: waitForDebuggerOnStart,
	}
}

// Do executes Target.setAutoAttach against the provided context.
func (p *SetAutoAttachParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetAutoAttach, p, nil)
}

// SetDiscoverTargetsParams controls whether to discover available targets
// and notify via targetCreated/targetInfoChanged/targetDestroyed events.
type SetDiscoverTargetsParams struct {
	Discover bool `json:"discover"` // Whether to discover available targets.
}

// SetDiscoverTargets controls whether to discover available targets and
// notify via targetCreated/targetInfoChanged/targetDestroyed events.
//
// parameters:
//   discover - Whether to discover available targets.
func SetDiscoverTargets(discover bool) *SetDiscoverTargetsParams {
	return &SetDiscoverTargetsParams{
		Discover: discover,
	}
}

// Do executes Target.setDiscoverTargets against the provided context.
func (p *SetDiscoverTargetsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetDiscoverTargets, p, nil)
}

// SetRemoteLocationsParams enables target discovery for the specified
// locations, when setDiscoverTargets was set to true.
type SetRemoteLocationsParams struct {
	Locations []*RemoteLocation `json:"locations"` // List of remote locations.
}

// SetRemoteLocations enables target discovery for the specified locations,
// when setDiscoverTargets was set to true.
//
// parameters:
//   locations - List of remote locations.
func SetRemoteLocations(locations []*RemoteLocation) *SetRemoteLocationsParams {
	return &SetRemoteLocationsParams{
		Locations: locations,
	}
}

// Do executes Target.setRemoteLocations against the provided context.
func (p *SetRemoteLocationsParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetRemoteLocations, p, nil)
}

// Command names.
const (
	CommandActivateTarget        = "Target.activateTarget"
	CommandAttachToTarget        = "Target.attachToTarget"
	CommandCloseTarget           = "Target.closeTarget"
	CommandCreateBrowserContext  = "Target.createBrowserContext"
	CommandCreateTarget          = "Target.createTarget"
	CommandDetachFromTarget      = "Target.detachFromTarget"
	CommandDisposeBrowserContext = "Target.disposeBrowserContext"
	CommandGetTargetInfo         = "Target.getTargetInfo"
	CommandGetTargets            = "Target.getTargets"
	CommandSendMessageToTarget   = "Target.sendMessageToTarget"
	CommandSetAutoAttach         = "Target.setAutoAttach"
	CommandSetDiscoverTargets    = "Target.setDiscoverTargets"
	CommandSetRemoteLocations    = "Target.setRemoteLocations"
)
