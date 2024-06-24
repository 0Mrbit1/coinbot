package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
)

// Define a struct to hold the parsed JSON data
type Statistics struct {
    UsdPrice           string `json:"usd_price"`
    EthPrice           string `json:"eth_price"`
    TotalSupply        string `json:"total_supply"`
    CirculatingSupply  string `json:"circulating_supply"`
    APY                string `json:"apy"`
    TotalStakedAmount  string `json:"total_staked_amount"`
    TotalStakers       string `json:"total_stakers"`
}

// Function to format the data
func formatStatistics(stats Statistics) string {
    // Convert large numbers with commas and treat them as integers
    totalSupply := addCommas(stats.TotalSupply)
    circulatingSupply := addCommas(stats.CirculatingSupply)
    totalStakedAmount := addCommas(stats.TotalStakedAmount)

    usdPrice := stats.UsdPrice
    apy := stats.APY

    formattedString := fmt.Sprintf(
        "💲price : %s\n💥 TOTAL-SUPPLY : %s\n📊 APY : %s%%\n🔒 TOTAL-STAKED-AMOUNT : %s\n🔓 circulating supply : %s\n👥 TOTAL-STAKERS : %s",
        usdPrice,
        totalSupply,
        apy,
        totalStakedAmount,
        circulatingSupply,
        stats.TotalStakers,
    )

    return formattedString
}

// Helper function to add commas to a number string
func addCommas(numStr string) string {
    parts := strings.Split(numStr, ".")
    integerPart := parts[0]
    n := len(integerPart)

    if n <= 3 {
        return integerPart
    }

    var result strings.Builder
    remainder := n % 3

    if remainder != 0 {
        result.WriteString(integerPart[:remainder])
        result.WriteString(",")
    }

    for i := remainder; i < n; i += 3 {
        result.WriteString(integerPart[i:i+3])
        if i+3 < n {
            result.WriteString(",")
        }
    }

    return result.String()
}

func fetchStatistics() (string, error) {
    var data Statistics
    url := "https://feyorra.com/data/statistics"

    // Perform HTTP GET request
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    // Parse the JSON data
    err = json.Unmarshal(body, &data)
    if err != nil {
        return "", err
    }

    formatted_data := formatStatistics(data)

    // Return the formatted string
    return formatted_data, nil
}



