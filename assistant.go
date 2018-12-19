package main

import (
	"encoding/json"
	"fmt"
	color "github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type DelayedQuoteResponse struct {
	Symbol           string  `json:"symbol"`
	DelayedPrice     float64 `json:"delayedPrice"`
	High             float64 `json:"high"`
	Low              float64 `json:"low"`
	DelayedSize      int     `json:"delayedSize"`
	DelayedPriceTime int64   `json:"delayedPriceTime"`
	ProcessedTime    int64   `json:"processedTime"`
}

type SectorPerformanceResponse []struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Performance float64 `json:"performance"`
	LastUpdated int64   `json:"lastUpdated"`
}

type DarkSkyWeatherResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timezone  string  `json:"timezone"`
	Currently struct {
		Time                 int     `json:"time"`
		Summary              string  `json:"summary"`
		Icon                 string  `json:"icon"`
		NearestStormDistance int     `json:"nearestStormDistance"`
		NearestStormBearing  int     `json:"nearestStormBearing"`
		PrecipIntensity      float64 `json:"precipIntensity"`
		PrecipProbability    float64 `json:"precipProbability"`
		Temperature          float64 `json:"temperature"`
		ApparentTemperature  float64 `json:"apparentTemperature"`
		DewPoint             float64 `json:"dewPoint"`
		Humidity             float64 `json:"humidity"`
		Pressure             float64 `json:"pressure"`
		WindSpeed            float64 `json:"windSpeed"`
		WindGust             float64 `json:"windGust"`
		WindBearing          int     `json:"windBearing"`
		CloudCover           float64 `json:"cloudCover"`
		UvIndex              int     `json:"uvIndex"`
		Visibility           float64 `json:"visibility"`
		Ozone                float64 `json:"ozone"`
	} `json:"currently"`
	Minutely struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                 int     `json:"time"`
			PrecipIntensity      float64 `json:"precipIntensity"`
			PrecipProbability    float64 `json:"precipProbability"`
			PrecipIntensityError float64 `json:"precipIntensityError,omitempty"`
			PrecipAccumulation   float64 `json:"precipAccumulation,omitempty"`
			PrecipType           string  `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"minutely"`
	Hourly struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Pressure            float64 `json:"pressure"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         int     `json:"windBearing"`
			CloudCover          float64 `json:"cloudCover"`
			UvIndex             int     `json:"uvIndex"`
			Visibility          float64 `json:"visibility"`
			Ozone               float64 `json:"ozone"`
			PrecipAccumulation  float64 `json:"precipAccumulation,omitempty"`
			PrecipType          string  `json:"precipType,omitempty"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                        int     `json:"time"`
			Summary                     string  `json:"summary"`
			Icon                        string  `json:"icon"`
			SunriseTime                 int     `json:"sunriseTime"`
			SunsetTime                  int     `json:"sunsetTime"`
			MoonPhase                   float64 `json:"moonPhase"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipAccumulation          float64 `json:"precipAccumulation,omitempty"`
			PrecipType                  string  `json:"precipType"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int     `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int     `json:"temperatureLowTime"`
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Pressure                    float64 `json:"pressure"`
			WindSpeed                   float64 `json:"windSpeed"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int     `json:"windGustTime"`
			WindBearing                 int     `json:"windBearing"`
			CloudCover                  float64 `json:"cloudCover"`
			UvIndex                     int     `json:"uvIndex"`
			UvIndexTime                 int     `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			Ozone                       float64 `json:"ozone"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int     `json:"temperatureMinTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
		} `json:"data"`
	} `json:"daily"`
	Alerts []struct {
		Title       string   `json:"title"`
		Regions     []string `json:"regions"`
		Severity    string   `json:"severity"`
		Time        int      `json:"time"`
		Expires     int      `json:"expires"`
		Description string   `json:"description"`
		URI         string   `json:"uri"`
	} `json:"alerts"`
	Flags struct {
		Sources        []string `json:"sources"`
		NearestStation float64  `json:"nearest-station"`
		Units          string   `json:"units"`
	} `json:"flags"`
	Offset int `json:"offset"`
}

var prefix = "https://api.iextrading.com/1.0"
var darkSkyKey = "4fcd5742700dd8303d37becb1a54db67"

func main() {
	vgtQuote := getDelayedQuote("vgt")
	vgtQuote.printDelayedQuote()
	sectorPerformance := getSectorPerformance()
	sectorPerformance.printSectorPerformance()
	weatherResponse := getWeatherResponse("41.829624", "-87.653075") //Bridgeport, Chicago
	weatherResponse.printWeatherResponse()
}

func getDelayedQuote(stock string) DelayedQuoteResponse {
	var delayedQuoteResponse DelayedQuoteResponse
	resp, err := http.Get(fmt.Sprintf("%s/stock/%s/delayed-quote", prefix, stock))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &delayedQuoteResponse)
	if err != nil {
		fmt.Println(err)
	}
	return delayedQuoteResponse
}

func getSectorPerformance() SectorPerformanceResponse {
	var sectorPerformanceResponse SectorPerformanceResponse
	resp, err := http.Get(fmt.Sprintf("%s/stock/market/sector-performance", prefix))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &sectorPerformanceResponse)
	if err != nil {
		fmt.Println(err)
	}
	return sectorPerformanceResponse
}

func getWeatherResponse(latitude, longitude string) DarkSkyWeatherResponse {
	var darkSkyWeatherResponse DarkSkyWeatherResponse
	resp, err := http.Get(fmt.Sprintf("https://api.darksky.net/forecast/%s/%s,%s", darkSkyKey, latitude, longitude))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &darkSkyWeatherResponse)
	if err != nil {
		fmt.Println(err)
	}
	return darkSkyWeatherResponse
}

func (delayedQuote DelayedQuoteResponse) printDelayedQuote() {
	color.Cyan(fmt.Sprintf("\n%v Summary:\n", delayedQuote.Symbol))
	fmt.Printf("\tDelayedPrice: %v\n\tHigh: %v\n\tLow: %v\n", delayedQuote.DelayedPrice, delayedQuote.High, delayedQuote.Low)
	color.Cyan("\nFinancial Sector Summary:\n")
}

func (s SectorPerformanceResponse) printSectorPerformance() {
	for _, v := range s {
		if v.Performance > 0 {
			fmt.Printf("\t%v Performance: ", v.Name)
			color.Green(fmt.Sprintf("%v", v.Performance))
		} else {
			fmt.Printf("\t%v Performance: ", v.Name)
			color.Red(fmt.Sprintf("%v", v.Performance))
		}
	}
}

func (weatherResponse DarkSkyWeatherResponse) printWeatherResponse() {
	color.Cyan("\nBridgeport Weather: ")
	fmt.Printf("Summary: %v\n", weatherResponse.Daily.Summary)
	sunriseTime, err := stringToTime(strconv.Itoa(weatherResponse.Daily.Data[0].SunriseTime))
	if err != nil {
		fmt.Println(err)
	}
	sunsetTime, err := stringToTime(strconv.Itoa(weatherResponse.Daily.Data[0].SunsetTime))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nCurrently:\n\tSummary: %v\n\tTemperature: %v\n\tApparent Temperature: %v\n\tPrecipProbability: %v\n\tNearest Storm Distance: %v\n\tUV Index: %v \n\tWind Gust: %v\n\tSunrise Time: %v\n\tSunset Time: %v\n", weatherResponse.Currently.Summary, weatherResponse.Currently.Temperature, weatherResponse.Currently.ApparentTemperature, weatherResponse.Currently.PrecipProbability, weatherResponse.Currently.NearestStormDistance, weatherResponse.Currently.UvIndex, weatherResponse.Currently.WindGust, sunriseTime, sunsetTime)
	sunriseTime, err = stringToTime(strconv.Itoa(weatherResponse.Daily.Data[1].SunriseTime))
	if err != nil {
		fmt.Println(err)
	}
	sunsetTime, err = stringToTime(strconv.Itoa(weatherResponse.Daily.Data[1].SunsetTime))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("\nTomorrow:\n\tSummary: %v\n\tTemperature High: %v\n\tTemperature Low: %v\n\tApparent Temperature Low: %v\n\tPrecipProbability: %v\n\tUV Index: %v \n\tWind Gust: %v\n\tSunrise Time: %v\n\tSunset Time: %v\n", weatherResponse.Daily.Data[0].Summary, weatherResponse.Daily.Data[0].TemperatureMax, weatherResponse.Daily.Data[0].TemperatureMin, weatherResponse.Daily.Data[0].ApparentTemperatureMin, weatherResponse.Daily.Data[0].PrecipProbability, weatherResponse.Daily.Data[0].UvIndex, weatherResponse.Daily.Data[0].WindGust, sunriseTime, sunsetTime)
}

func stringToTime(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0), nil
}
