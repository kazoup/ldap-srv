package main

import (
	"fmt"

	"github.com/kazoup/ldap-srv/handler"
	"github.com/kazoup/ldap-srv/proto"
	"github.com/micro/go-micro"
)

func main() {
	//Crete new service
	service := micro.NewService(
		micro.Name("kazoup.com.srv.ldap"),
		micro.Version("v1"),
		micro.Metadata(map[string]string{
			"type": "template",
		}),
	)
	service.Init()
	ldap.RegisterLDAPHandler(service.Server(), new(handler.LDAP))
	//Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
