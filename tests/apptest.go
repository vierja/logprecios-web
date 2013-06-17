package tests

import "github.com/robfig/revel"

type AppTest struct {
    revel.TestSuite
}

func (t AppTest) Before() {
    println("Set up")
}

func (t AppTest) TestThatIndexPageWorks() {
    t.Get("/")
    t.AssertOk()
    t.AssertContentType("text/html")
}

func (t AppTest) After() {
    println("Tear down")
}

func (t AppTest) TestProductJson() {
    t.Get("/product/1234.json")
    t.AssertOk()
    t.AssertContentType("application/json")
}
