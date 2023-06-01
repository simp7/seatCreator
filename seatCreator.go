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
	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 11},pos.Absolute{X: 4,Y: 11})
	rect2 := eraser.Rectangle(pos.Absolute{X: 5,Y: 4},pos.Absolute{X: 6,Y: 4})
	integrated := eraser.Integrated(rect1, rect2)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         4,
		YSize:         8,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('A'),
	}

	specific := eraser.Position(pos.Absolute{X: 8,Y: 5}, pos.Absolute{X: 8, Y: 7}, pos.Absolute{X: 8, Y: 9}, pos.Absolute{X: 8, Y: 11})
	integrated2 := eraser.Integrated(specific)
	base = model.NewSeatBase(8, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         9,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('B'),
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 25,Y: 4},pos.Absolute{X: 25, Y: 11})
	rect4 := eraser.Rectangle(pos.Absolute{X: 21,Y: 4}, pos.Absolute{X: 22,Y: 4})
	integrated3 := eraser.Integrated(rect3, rect4)
	base = model.NewSeatBase(21, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('C'),
	}

	rect5 := eraser.Rectangle(pos.Absolute{X: 7,Y: 4}, pos.Absolute{X: 20,Y: 4})
	integrated4 := eraser.Integrated(rect5)
	base = model.NewSeatBase(5, 4, "휠체어석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         18,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
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
		Key:             "1F",
		Seats:           seats,
		XSize:           24,
		YSize:           9,
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