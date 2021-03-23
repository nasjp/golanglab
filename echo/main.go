package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GetApplicant struct {
	OrgID             uint   `params:"orgId"`
	KYCOrgApplicantID string `json:"kycOrgApplicantId"`
}

func main() {
	e := echo.New()
	e.GET("orgs/:orgId/applicants/:kycOrgApplicantId", func(c echo.Context) error {

		p := &GetApplicant{}
		if err := c.Bind(p); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, *p)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
