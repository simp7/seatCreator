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
	"github.com/simp7/seatCreator/model/pos"
)

func ArtriumSmall() model.Group {

	specific2 := eraser.Position(pos.Absolute{X: 24,Y: 11})
	integrated3 := eraser.Integrated( specific2)
	base := model.NewSeatBase(2, 5, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         4,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('C'),
	}

	rect2 := eraser.Rectangle(pos.Absolute{X: 20,Y: 12},pos.Absolute{X: 20,Y: 13})
	specific3 := eraser.Position(pos.Absolute{X: 20, Y: 15})
	integrated4 := eraser.Integrated(rect2, specific3)
	base = model.NewSeatBase(13, 5, "D석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 3, 9),
		XSize:         8,
		YSize:         5,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('D'),
	}


	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	return group.Mixed(block3, block4)
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
		XSize:           18,
		YSize:           13,
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