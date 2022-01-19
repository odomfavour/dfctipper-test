package app

import (
	"fmt"
	"net/http"

	"github.com/ademuanthony/dfctipper/web"
)

func (a app) homePage(w http.ResponseWriter, r *http.Request) {

	data := struct {
		*web.CommonPageData
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: a.server.CommonData(w, r),
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Home",
				Active:    true,
			},
		},
	}
	web.RenderHTML("home", w, r, data, a.server)
}

func (a app) advertiser(w http.ResponseWriter, r *http.Request) {

	data := struct {
		*web.CommonPageData
		BreadcrumbItems []web.BreadcrumbItem
	}{
		CommonPageData: a.server.CommonData(w, r),
		BreadcrumbItems: []web.BreadcrumbItem{
			{
				HyperText: "Home",
				Active:    true,
			},
		},
	}
	web.RenderHTML("advertiser", w, r, data, a.server)
}

func (a app) contactPostBack(w http.ResponseWriter, r *http.Request) {
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	twitterHandle := r.FormValue("twitter_handle")

	message := `
	Name: %s %s
	Email: %s
	Phone: %s
	Twitter: %s
	`

	title := "New Message From DFCTipper"
	sender := "dfctipper@club250cent.com"

	if err := a.SendEmail(r.Context(), sender, "info@deficonnect.tech", title,
		fmt.Sprintf(message, firstname, lastname, email, phone, twitterHandle)); err != nil {
		log.Errorf("SendEmail", err)
		web.NotifyError(w, "Error is sending message. Please contact marketing@deficonnect.tech")
	} else {
		web.NotifySuccess(w, "Message sent")
	}
	http.Redirect(w, r, "/advertiser", http.StatusSeeOther)
}
