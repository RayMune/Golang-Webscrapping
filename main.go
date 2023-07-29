package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "strings"
)

func main() {
  resp, err := http.Get("https://www.jumia.co.ke/catalog/?q=playstation")
  if err != nil {
    fmt.Println(err)
    return
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println(err)
    return
  }

  // Find all the product cards on the page
  productCards := strings.Split(string(body), "<div class=\"product-card\">")[1:]

  // For each product card, extract the product name, price, and image URL
  for _, productCard := range productCards {
    // Extract the product name
    productName := strings.Split(productCard, "<h2>")[1]
    productName = strings.Split(productName, "</h2>")[0]

    // Extract the price
    price := strings.Split(productCard, "<span class=\"price\">")[1]
    price = strings.Split(price, "</span>")[0]

    // Extract the image URL
    imageURL := strings.Split(productCard, "<img src=\"")[1]
    imageURL = strings.Split(imageURL, "\" alt=\"")[0]

    // Print the product information
    fmt.Println("Product name:", productName)
    fmt.Println("Price:", price)
    fmt.Println("Image URL:", imageURL)
  }
}

