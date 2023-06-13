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
	base := model.NewSeatBase(3, 4, "C석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         4,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('C'),
		Reverse: true,
	}

	specific := eraser.Position(pos.Absolute{X: 8, Y: 14}, pos.Absolute{X: 19, Y: 14})
	integrated3 := eraser.Integrated(specific)
	base = model.NewSeatBase(8, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         11,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('B'),
		Reverse: true,
	}

	rect := eraser.Rectangle(pos.Absolute{X: 21, Y: 8}, pos.Absolute{X: 26, Y: 8})
	integrated2 := eraser.Integrated(rect)
	base = model.NewSeatBase(21, 4, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         10,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('A'),
		Reverse: true,
	}
	
	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1, block2, block3)
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
		XSize:           34,
		YSize:           11,
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
