package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
)

func main() {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "test.exampledomain.com", 389))
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind("cn=test user,ou=test,dc=exampledomain,dc=com", "Exampledomain")
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		"dc=exampledomain,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(objectClass=organizationalPerson))",
		[]string{"dn", "cn", "sAMAccountName", "mail", "department", "manager", "memberOf"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("CN: %s\n", entry.GetAttributeValue("cn"))
		fmt.Printf("sAMAccountName: %s\n", entry.GetAttributeValue("sAMAccountName"))
		fmt.Printf("mail: %s\n", entry.GetAttributeValue("mail"))
		fmt.Printf("department: %s\n", entry.GetAttributeValue("department"))
		fmt.Printf("manager: %s\n", entry.GetAttributeValue("manager"))
		fmt.Printf("memberOf: %v\n", entry.GetAttributeValues("memberOf"))
		fmt.Println("---")
	}
}
