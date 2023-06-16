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

	hallA := eraser.VerticalHallway(15,23)
	integrated := eraser.Integrated(hallA)
	base := model.NewSeatBase(10, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         19,
		YSize:         6,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('A'),
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 8, Y: 11}, pos.Absolute{X: 13, Y:  11})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 26, Y: 11}, pos.Absolute{X: 31, Y:  11})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 8, Y: 12}, pos.Absolute{X: 11, Y:  12})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 28, Y: 12}, pos.Absolute{X: 31, Y:  12})
	Bspecific := eraser.Position(pos.Absolute{X: 17, Y: 11}, pos.Absolute{X: 22,Y: 11},
								pos.Absolute{X: 16, Y: 12}, pos.Absolute{X: 23,Y: 12},
								pos.Absolute{X: 13, Y: 13}, pos.Absolute{X: 26, Y: 13},
							)
	integrated2 := eraser.Integrated(Brect1, Brect2, Brect3, Brect4, Bspecific)
	base = model.NewSeatBase(8, 11, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         24,
		YSize:         3,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('B'),
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X:11, Y: 15 }, pos.Absolute{X: 11, Y: 16})
	Crect2 := eraser.Rectangle(pos.Absolute{X:28, Y: 15 }, pos.Absolute{X: 28, Y: 16})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 1, Y: 15}, pos.Absolute{X: 2, Y: 15})
	Crect4 := eraser.Rectangle(pos.Absolute{X: 37, Y: 15}, pos.Absolute{X: 38, Y: 15})
	integrated3 := eraser.Integrated(Crect1,Crect2, Crect3,Crect4)
	base = model.NewSeatBase(1, 15, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         38,
		YSize:         2,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Standard(),
	}

	integrated4 := eraser.Integrated()
	base = model.NewSeatBase(8, 4, "L석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('L'),
	}

	integrated5 := eraser.Integrated()
	base = model.NewSeatBase(30, 4, "R석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Prefix('R'),
	}

	Wrect1 := eraser.Rectangle(pos.Absolute{X: 9, Y: 9}, pos.Absolute{X: 29, Y: 10})
	integrated6 := eraser.Integrated(Wrect1)
	base = model.NewSeatBase(8, 9, "휠체어석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         23,
		YSize:         2,
		EmptyChecker:  integrated6,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)

	return group.Mixed(block1, block2, block3, block4,block5, block6)
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
		XSize:           36,
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
