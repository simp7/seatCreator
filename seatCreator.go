package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/gin-gonic/gin"
	"github.com/simp7/seatCreator/model"
	"github.com/simp7/seatCreator/model/area"
	"github.com/simp7/seatCreator/model/eraser"
	"github.com/simp7/seatCreator/model/group"
	"github.com/simp7/seatCreator/model/nameformatter"
)

func ArtriumSmall() model.Group {
	vertical := eraser.VerticalHallway(11,28)

	nameFormatter := nameformatter.Standard()
	integrated := eraser.Integrated(vertical)

	base := model.NewSeatBase(3, 5, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 16),
		XSize:         34,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(6, 3, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 4, 14),
		XSize:         28,
		YSize:         2,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 14, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 25),
		XSize:         4,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(13, 14, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 10, 25),
		XSize:         24,
		YSize:         1,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	return group.Mixed(block1, block2, block3, block4)
}


func copy(target model.Group) {
	err := clipboard.WriteAll(target.String())
	if err != nil {
		fmt.Println("Error occured when copying")
		return
	}
	fmt.Println("Copied!")
}
func handler(c *gin.Context) {
	seats := ArtriumSmall() // Put Seating Here
	target := area.Area{
		Key:             "2F",
		Seats:           seats,
		XSize:           34,
		YSize:           12,
		BackgroundImage: "",
		Color:           "#ff9f00",
	}
  
	fmt.Fprintf(c.Writer, "<h1>Seats Map</h1><body>\n%s\n</body>", target.Html())
	copy(target)
}
func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", handler)
	r.Run()
}