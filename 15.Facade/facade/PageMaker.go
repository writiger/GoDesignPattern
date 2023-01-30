package facade

import "fmt"

func MakeWelcomePage(mail, filename string) error {
	d := &database{}
	username, err := d.GetProperty(mail)
	if err != nil {
		return err
	}
	hm := &htmlMaker{}
	hm.Title("Welcome to " + username + "'s Page!")
	hm.Paragraph("欢迎来到" + username + "的主页")
	hm.Paragraph("等着你的邮件哦 ！")
	hm.MailTo(mail, username)
	hm.Close()
	if err = hm.MakeHtml(filename); err != nil {
		return err
	}
	fmt.Printf("%s is created for %s ( %s )", filename, mail, username)
	return nil
}
