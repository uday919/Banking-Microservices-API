package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "uday", City: "Hyderabad", Zipcode: "500037", DateofBirth: "01-01-2000", Status: "Active"},
		{Id: "1002", Name: "venkat", City: "chennai", Zipcode: "503037", DateofBirth: "01-08-2001", Status: "Active"},
		{Id: "1003", Name: "durga", City: "delhi", Zipcode: "500837", DateofBirth: "01-02-2004", Status: "Active"},
		{Id: "1004", Name: "manu", City: "bangalore", Zipcode: "555037", DateofBirth: "01-03-2003", Status: "Active"},
	}
	return CustomerRepositoryStub{customers}
}
