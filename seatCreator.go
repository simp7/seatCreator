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

	integrated := eraser.Integrated()

	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         7,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('A'),
	}

	rect1 := eraser.Rectangle(pos.Absolute{X: 20,Y: 4},pos.Absolute{X: 20,Y: 5})
	specific := eraser.Position(pos.Absolute{X: 20,Y: 7})
	integrated2 := eraser.Integrated(rect1, specific)
	base = model.NewSeatBase(13, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 3, 8),
		XSize:         8,
		YSize:         7,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('B'),
	}


	// specific2 := eraser.Position(pos.Absolute{X: 24,Y: 11})
	// integrated3 := eraser.Integrated( specific2)
	// base = model.NewSeatBase(2, 12, "A석")
	// blockInput3 := group.BlockInput{
	// 	Criteria:      model.NewSeat(base, 1, 1),
	// 	XSize:         10,
	// 	YSize:         4,
	// 	EmptyChecker:  integrated3,
	// 	NameFormatter: nameformatter.Floor(nameformatter.Prefix('C'),2),
	// }

	// rect2 := eraser.Rectangle(pos.Absolute{X: 20,Y: 12},pos.Absolute{X: 20,Y: 13})
	// specific3 := eraser.Position(pos.Absolute{X: 20, Y: 15})
	// integrated4 := eraser.Integrated(rect2, specific3)
	// base = model.NewSeatBase(13, 12, "A석")
	// blockInput4 := group.BlockInput{
	// 	Criteria:      model.NewSeat(base, 3, 9),
	// 	XSize:         8,
	// 	YSize:         5,
	// 	EmptyChecker:  integrated4,
	// 	NameFormatter: nameformatter.Floor(nameformatter.Prefix('D'),2),
	// }


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	// block3 := group.HorizontalBlock(blockInput3)
	// block4 := group.HorizontalBlock(blockInput4)
	return group.Mixed(block1, block2)
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
		Key:             "1F",
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