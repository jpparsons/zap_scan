package main

// https://github.com/sendgrid/sendgrid-go/blob/main/use-cases/attachments-with-mailer-helper.md

import (
  "fmt"
  "log"
  "os"
  "encoding/base64"
  "io/ioutil"
  "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
  // create new *SGMailV3
  m := mail.NewV3Mail()

  from := mail.NewEmail("ProductPlan", os.Getenv("FROM_EMAIL"))
  to := mail.NewEmail("ProductPlan", os.Getenv("TO_EMAIL"))

  // create new *Personalization
  personalization := mail.NewPersonalization()
  personalization.AddTos(to)
  personalization.Subject = "Weekly Security Report"

  // add `personalization` to `m`
  m.AddPersonalizations(personalization)

  // read/attach the report html file
  attachment := mail.NewAttachment()
  dat, err := ioutil.ReadFile("/home/circleci/project/security_report.html")
  if err != nil {
    fmt.Println(err)
  }
  encoded := base64.StdEncoding.EncodeToString([]byte(dat))
  attachment.SetContent(encoded)
  attachment.SetType("text/plain")
  attachment.SetFilename("security_report.html")
  attachment.SetDisposition("attachment")

  // also include report in email body
  content := mail.NewContent("text/html", string(dat))
  m.AddContent(content)

  m.SetFrom(from)
  m.AddAttachment(attachment)

  request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
  request.Method = "POST"
  request.Body = mail.GetRequestBody(m)
  response, err := sendgrid.API(request)
  if err != nil {
    log.Println(err)
  } else {
    fmt.Println(response.StatusCode)
    fmt.Println(response.Body)
    fmt.Println(response.Headers)
  }
}
