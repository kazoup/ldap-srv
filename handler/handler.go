package handler

import (
	"log"

	client "github.com/jtblin/go-ldap-client"
	"github.com/kazoup/ldap-srv/proto"
	"github.com/micro/go-micro/errors"
	"golang.org/x/net/context"
)

//Config ...
type Config struct{}

//Create ...
func (c *Config) Create(
	ctx context.Context,
	req *ldap.CreateRequest,
	res *ldap.CreateResponse) error {
	//Check if req not empty
	log.Print(req)
	if req.Config == nil {
		return errors.BadRequest("com.kazoup.srv.ldap.Config.Create", "Configuration can not be blank")
	}
	//Create LDAP client
	client := &client.LDAPClient{
		Base:         req.Config.BaseDn,
		Host:         req.Config.Host,
		Port:         int(req.Config.Port),
		UseSSL:       req.Config.UseSsl,
		BindDN:       req.Config.BindDn,
		BindPassword: req.Config.BindDn,
		UserFilter:   req.Config.UserFilter,  // "(uid=%s)"
		GroupFilter:  req.Config.GroupFilter, //"(memberUid=%s)",
		Attributes:   req.Config.Attr,        //[]string{"givenName", "sn", "mail", "uid"},
	}
	//defer connection close
	defer client.Close()
	//Validate binding with AD
	err := client.Connect()
	if err != nil {
		return errors.BadRequest("com.kazoup.srv.ldap.Config.Create Can not bind to AD", err.Error())
	}
	//TODO: We succesfuly bind to AD lets save in ES es-srv

	return nil
}
