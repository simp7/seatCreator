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

	Arect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 13}, pos.Absolute{X: 4, Y: 13})
	integrated := eraser.Integrated(Arect1)
	base := model.NewSeatBase(3, 7, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         7,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 18, Y: 13}, pos.Absolute{X: 22, Y: 13})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 20, Y: 12}, pos.Absolute{X: 22, Y: 12})
	integrated2 := eraser.Integrated(Brect1, Brect2)
	base = model.NewSeatBase(13, 7, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         7,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X: 24, Y: 13}, pos.Absolute{X: 36, Y: 13})
	integrated3 := eraser.Integrated(Crect1)
	base = model.NewSeatBase(24, 7, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         13,
		YSize:         7,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}
	
	Drect1 := eraser.Rectangle(pos.Absolute{X: 38, Y: 13}, pos.Absolute{X: 42, Y: 13})
	Drect2 := eraser.Rectangle(pos.Absolute{X: 38, Y: 12}, pos.Absolute{X: 40, Y: 12})
	integrated4 := eraser.Integrated(Drect1, Drect2)
	base = model.NewSeatBase(38, 7, "D석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         7,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	Erect1 := eraser.Rectangle(pos.Absolute{X: 56, Y: 13}, pos.Absolute{X: 57, Y: 13})
	integrated5 := eraser.Integrated(Erect1)
	base = model.NewSeatBase(49, 7, "E석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         7,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	Lrect := eraser.Rectangle(pos.Absolute{X: 6, Y: 4}, pos.Absolute{X: 10, Y: 4})
	integrated6 := eraser.Integrated(Lrect)
	base = model.NewSeatBase(3, 4, "L석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         2,
		EmptyChecker:  integrated6,
		NameFormatter: nameformatter.Prefix('L'),
	}


	Rrect := eraser.Rectangle(pos.Absolute{X: 50, Y: 4}, pos.Absolute{X: 54, Y: 4})
	integrated7 := eraser.Integrated(Rrect)
	base = model.NewSeatBase(50, 4, "R석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         8,
		YSize:         2,
		EmptyChecker:  integrated7,
		NameFormatter: nameformatter.Prefix('R'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)

	return group.Mixed(block1, block2, block3, block4, block5, block6, block7)
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
		Key:             "3F",
		Seats:           seats,
		XSize:           55,
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
