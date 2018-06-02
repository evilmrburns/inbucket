// Package webui powers Inbucket's web GUI
package webui

import (
	"github.com/gorilla/mux"
	"github.com/jhillyerd/inbucket/pkg/server/web"
)

// NewRouter returns a router configured for the web UI.
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.Path("/serve/").Handler(
		web.Handler(RootIndex)).Name("RootIndex").Methods("GET")
	r.Path("/serve/monitor").Handler(
		web.Handler(RootMonitor)).Name("RootMonitor").Methods("GET")
	r.Path("/serve/monitor/{name}").Handler(
		web.Handler(RootMonitorMailbox)).Name("RootMonitorMailbox").Methods("GET")
	r.Path("/serve/status").Handler(
		web.Handler(RootStatus)).Name("RootStatus").Methods("GET")
	r.Path("/serve/link/{name}/{id}").Handler(
		web.Handler(MailboxLink)).Name("MailboxLink").Methods("GET")
	r.Path("/serve/mailbox").Handler(
		web.Handler(MailboxIndex)).Name("MailboxIndex").Methods("GET")
	r.Path("/serve/mailbox/{name}").Handler(
		web.Handler(MailboxList)).Name("MailboxList").Methods("GET")
	r.Path("/serve/mailbox/{name}/{id}").Handler(
		web.Handler(MailboxShow)).Name("MailboxShow").Methods("GET")
	r.Path("/serve/mailbox/{name}/{id}/html").Handler(
		web.Handler(MailboxHTML)).Name("MailboxHtml").Methods("GET")
	r.Path("/serve/mailbox/{name}/{id}/source").Handler(
		web.Handler(MailboxSource)).Name("MailboxSource").Methods("GET")
	r.Path("/serve/mailbox/dattach/{name}/{id}/{num}/{file}").Handler(
		web.Handler(MailboxDownloadAttach)).Name("MailboxDownloadAttach").Methods("GET")
	r.Path("/serve/mailbox/vattach/{name}/{id}/{num}/{file}").Handler(
		web.Handler(MailboxViewAttach)).Name("MailboxViewAttach").Methods("GET")
	r.Path("/serve/{name}").Handler(
		web.Handler(MailboxIndexFriendly)).Name("MailboxListFriendly").Methods("GET")
	return r
}
