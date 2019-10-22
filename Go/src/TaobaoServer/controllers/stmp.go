package controllers

import (
	md "TaobaoServer/models"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"net/smtp"

	"github.com/astaxie/beego/logs"
)

//the templace of the email that signUp and get comfirm code
var signUpMailTP = `
<div style="background-color:#4caf50;width: 400px;height30200px;padding: 10px;border-radius: 6px;font-weight: 500;margin: 20px;">
🚂  🚃  🚄  🚅  🚆  🚇  🚈  🚉  🚊  🚝  🚞  🚋 🚲 🚜<br>
你好！非常感谢你注册成为本站第%d个会员。<br>
本站服务永久免费，并将不断完善和更新,努力为你提供更好的体验。欢迎向我提出改进建议和问题反馈！  :)<br>
刚刚注册的账号：<span style=" color: #E91E63; font-weight: 600;">%s </span> <br>
验证码为：<span style=" color: #E91E63; font-weight: 600;">%s</span> <br>
(30分钟内有效,若非本人操作,请忽略此邮件)<br>
 🚌 🚍  🚎  🚏  🚐 🚑  🚒  🚓  🚔 🚕 🚖 🚗 🚘 🚚 🚛 <br>
</div>
`

var notificationTP = `
<div style="background-color:#68a8bb;width: 400px;height30200px;padding: 10px;border-radius: 6px;font-weight: 500;margin: 20px;">
🚂  🚃  🚄  🚅  🚆  🚇  🚈  🚉  🚊  🚝  🚞  🚋 🚲 🚜<br>
用户你好， 刚才用户 %s 向你发送了一条私信哦，内容如下：<br>
<pre style="color: blueviolet;font-size: 1.2em;font-weight: 600;"> %s </pre>
(若需要取消邮箱通知功能请到 "个人主页 -> 我的消息" 页面进行操作，感谢对本站的支持！) <br> 
 🚌 🚍  🚎  🚏  🚐 🚑  🚒  🚓  🚔 🚕 🚖 🚗 🚘 🚚 🚛 <br>
</div>
`

var resetPasswordTP = `
<div style="background-color:#68a8bb;width: 400px;height30200px;padding: 10px;border-radius: 6px;font-weight: 500;margin: 20px;">
🚂  🚃  🚄  🚅  🚆  🚇  🚈  🚉  🚊  🚝  🚞  🚋 🚲 🚜<br>
用户你好， 当前正在进行的重置密码操作验证码为：<span style=" color: #E91E63; font-weight: 600;">%s</span><br>
感谢你对本站的支持!<br> 
 🚌 🚍  🚎  🚏  🚐 🚑  🚒  🚓  🚔 🚕 🚖 🚗 🚘 🚚 🚛 <br>
</div>
`

//config variable can read from config file
var (
	stmpHost   = ""
	stmpPort   = 0
	myemail    = ""
	mypassword = ""
	sendEmail  = false
)

//create an auth
func createAutn() smtp.Auth {
	return smtp.PlainAuth(
		"",
		myemail,
		mypassword,
		stmpHost,
	)
}

//send the comfirm email to user after it register
//index is the rank of it new account
func SendConfrimEmail(account md.RegisterData, index int) error {
	if !sendEmail {
		return nil
	}
	toEmail, username, code := account.Email, account.Name, account.Code
	address := fmt.Sprintf("%s:%d", stmpHost, stmpPort)
	message := createEmail(toEmail, index, username, code)
	auth := createAutn()
	err := SendMailUsingTLS(address, auth, myemail, []string{toEmail}, []byte(message))
	if err != nil {
		logs.Error("Send email fall %v", err, 1)
		return err
	} else {
		logs.Warn("Send comfirm email to %s success!", toEmail)
		return nil
	}
}

//send a notice email to user after user receive a private message 🍣
func SendNotification(sender, toEmail, content string) error {
	if !sendEmail {
		return errors.New("Send email funciton is closed!")
	}
	header := make(map[string]string)
	header["From"] = "BlackCarDriver.cn" + "<" + myemail + ">"
	header["To"] = toEmail
	header["Subject"] = "收到私信提醒"
	header["Content-Type"] = "text/html; charset=UTF-8"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + fmt.Sprintf(notificationTP, sender, content)
	auth := createAutn()
	err := SendMailUsingTLS(fmt.Sprintf("%s:%d", stmpHost, stmpPort), auth, myemail, []string{toEmail}, []byte(message))
	if err != nil {
		rlog.Error("Send email fall %v", err, 1)
		return err
	} else {
		rlog.Warn("Send comfirm email to %s success!", toEmail)
		return nil
	}
}

//send a resetpassword comfirm code to user 🍥
func SendResetComfirm(toEmail, code string) error {
	if !sendEmail {
		return errors.New("Send email funciton is closed!")
	}
	header := make(map[string]string)
	header["From"] = "BlackCarDriver.cn" + "<" + myemail + ">"
	header["To"] = toEmail
	header["Subject"] = "重置密码验证"
	header["Content-Type"] = "text/html; charset=UTF-8"
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + fmt.Sprintf(resetPasswordTP, code)
	auth := createAutn()
	err := SendMailUsingTLS(fmt.Sprintf("%s:%d", stmpHost, stmpPort), auth, myemail, []string{toEmail}, []byte(message))
	if err != nil {
		rlog.Error("SendResetComfirm fall %v", err, 1)
		return err
	} else {
		rlog.Warn("SendResetComfirm  to %s success!", toEmail)
		return nil
	}
}

//create an emial by push the nessary varibale into the emil templace
func createEmail(toEmail string, num int, username string, code string) (message string) {
	header := make(map[string]string)
	header["From"] = "BlackCarDriver.cn" + "<" + myemail + ">"
	header["To"] = toEmail
	header["Subject"] = "来自blackcardriver.cn的验证码"
	header["Content-Type"] = "text/html; charset=UTF-8"
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + fmt.Sprintf(signUpMailTP, num, username, code)
	return message
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	//problem : certificate signed by unknown authority
	var tr = &tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", addr, tr)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
func SendMailUsingTLS(addr string, auth smtp.Auth, from string, to []string, msg []byte) (err error) {
	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		logs.Error("Create smpt client error %v", err)
		return err
	}
	defer c.Close()
	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}
	if err = c.Mail(from); err != nil {
		logs.Error(err)
		return err
	}
	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}
	w, err := c.Data()
	if err != nil {
		return err
	}
	if _, err = w.Write(msg); err != nil {
		return err
	}
	if err = w.Close(); err != nil {
		return err
	}
	return c.Quit()
}
