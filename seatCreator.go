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
	nameFormatter := nameformatter.Standard()

	Especific := eraser.Position(pos.Absolute{X: 13, Y: 31})
	integrated := eraser.Integrated(Especific)
	base := model.NewSeatBase(3, 28, "E석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         5,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	integrated2 := eraser.Integrated()
	base = model.NewSeatBase(15, 28, "F석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         5,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Gspecific := eraser.Position(pos.Absolute{X: 30, Y: 32})
	integrated3 := eraser.Integrated(Gspecific)
	base = model.NewSeatBase(30, 28, "G석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         5,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 3, "D석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         3,
		YSize:         24,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
		Reverse: true,
	}

	pos := eraser.Position(pos.Absolute{X: 39, Y: 3})
	integrated4 := eraser.Integrated(pos)
	base = model.NewSeatBase(38, 3, "H석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         3,
		YSize:         24,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}


	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.VerticalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)

	return group.Mixed(block1, block2, block3, block4, block5)
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
		XSize:           48,
		YSize:           10,
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
