package controllers

import (
    //"github.com/bradfitz/gomemcache/memcache"
    "github.com/robfig/revel"
    "time"
)

type ProductController struct {
    *revel.Controller
}

type Product struct {
    Id         int     `json:"id"`
    Name       string  `json:"name"`
    Url        string  `json:"url"`
    LastPrice  float64 `json:"last_price"`
    CurrencyId string  `json:"currency_id"`
}

type Price struct {
    Price      float64   `json:"price"`
    CurrencyId string    `json:"currency_id"`
    FromDate   time.Time `json:"from_date"`
    UntilDate  time.Time `json:"until_date"`
}

func (c ProductController) Show(id int) revel.Result {
    return c.Render()
}

func (c ProductController) PriceJson(id int) revel.Result {
    prices := []Price{
        {10.0, "UYU", time.Now().UTC(), time.Now().UTC()},
        {11.0, "UYU", time.Now().UTC(), time.Now().UTC()},
        {12.0, "UYU", time.Now().UTC(), time.Now().UTC()},
        {13.0, "UYU", time.Now().UTC(), time.Now().UTC()},
        {14.0, "UYU", time.Now().UTC(), time.Now().UTC()},
        {15.0, "UYU", time.Now().UTC(), time.Now().UTC()},
    }
    return c.RenderJson(prices)
}

func (c ProductController) ShowJson(id int) revel.Result {
    product := Product{
        Id:         1,
        Name:       "ARROZ SAMAN PARBOILED 1KG",
        Url:        "http://www.tinglesa.com.uy/producto.php?idarticulo=86",
        LastPrice:  33.00,
        CurrencyId: "UYU",
    }
    return c.RenderJson(product)
}
