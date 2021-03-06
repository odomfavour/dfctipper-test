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

	totalAccounts, _ := a.db.TotalAccounts(r.Context())
	verifiedAccounts, _ := a.db.VerifiedAccounts(r.Context())

	data := struct {
		*web.CommonPageData
		BreadcrumbItems  []web.BreadcrumbItem
		TotalAccounts    int64
		VerifiedAccounts int64
	}{
		CommonPageData:   a.server.CommonData(w, r),
		TotalAccounts:    totalAccounts,
		VerifiedAccounts: verifiedAccounts,
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

	if err := a.SendEmail(r.Context(), sender, "marketing@deficonnect.tech", title,
		fmt.Sprintf(message, firstname, lastname, email, phone, twitterHandle)); err != nil {
		log.Errorf("SendEmail", err)
		web.NotifyError(w, "Error is sending message. Please contact marketing@deficonnect.tech")
	} else {
		http.Redirect(w, r, "/thankyou", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/advertiserThankyou", http.StatusSeeOther)
}

func (a app) advertiserThankyou(w http.ResponseWriter, r *http.Request) {

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
	web.RenderHTML("advertiser-thankyou", w, r, data, a.server)
}
