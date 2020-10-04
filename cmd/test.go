package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/jax-b/iconfinderapi"
)

func main() {
	iconfinder := iconfinderapi.NewIconFinder("your api key")

	color.Set(color.FgYellow)
	fmt.Print("User 14: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetUserIDDetails(14))

	color.Set(color.FgYellow)
	fmt.Print("Style Flat: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetStyleDetails("flat"))

	color.Set(color.FgYellow)
	fmt.Print("Avalible Styles: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListAllStylesFast())

	color.Set(color.FgYellow)
	fmt.Print("License1 Details: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetLicenseDetails(1))

	color.Set(color.FgYellow)
	fmt.Print("Halloween Catagorie Details: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetCategoryDetails("halloween"))

	color.Set(color.FgYellow)
	fmt.Print("Avalible Catagories: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListAllCatagoriesFast())

	color.Set(color.FgYellow)
	fmt.Print("Author 100: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetAuthorDetails(100))

	color.Set(color.FgYellow)
	fmt.Print("Style outline iconsets: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListIconSetsOfStyle("outline", 3, -1, -1, -1, ""))

	color.Set(color.FgYellow)
	fmt.Print("Author 100 iconsets: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListIconSetsOfAuthor(100, 3, -1, -1, -1, ""))

	color.Set(color.FgYellow)
	fmt.Print("User 14 iconsets: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListIconSetsOfUser("14", 3, -1, -1, -1, ""))

	color.Set(color.FgYellow)
	fmt.Print("Category arrow iconsets: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListIconSetsOfCategory("arrow", 3, -1, -1, -1, ""))

	color.Set(color.FgYellow)
	fmt.Print("Icon Set 1761: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetIconSetDetails(1761))

	color.Set(color.FgYellow)
	fmt.Print("PublicIconSets: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.ListPublicIconSets(3, -1, -1, -1, ""))

	color.Set(color.FgYellow)
	fmt.Print("IconSet 1761: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.IconsInSet(1761, "", 3, -1, -1))

	color.Set(color.FgYellow)
	fmt.Print("Icon 182504: ")
	color.Set(color.FgWhite)
	fmt.Println(iconfinder.GetIconDetails(182504))

	color.Set(color.FgYellow)
	fmt.Print("Seach firefox: ")
	color.Set(color.FgWhite)
	searched, _ := iconfinder.SearchIcons("firefox", 1, -1, 0, 0, "", "", "flat")
	fmt.Println(searched)

	iconPath := searched.Icons[0].Rasters[0].Formats[0]
	img := iconfinder.DownloadIcon(iconPath)
	f, err := os.Create("test.png")
	defer f.Close()
	err = png.Encode(f, img)

	if err != nil {
		log.Fatal(err)
	}
}
