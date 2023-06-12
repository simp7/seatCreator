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
 	
	Arect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 11}, pos.Absolute{X: 3,Y: 14})
	integrated := eraser.Integrated(Arect1)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         7,
		YSize:         11,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 11,Y: 4}, pos.Absolute{X: 11,Y: 12})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 25,Y: 4}, pos.Absolute{X: 25,Y: 10})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 12,Y: 4}, pos.Absolute{X: 12,Y: 8})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 24,Y: 4}, pos.Absolute{X: 24,Y: 6})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 13,Y: 4}, pos.Absolute{X: 13,Y: 4})

	integrated2 := eraser.Integrated(Brect1,Brect2,Brect3,Brect4,Brect5)
	base = model.NewSeatBase(11, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         15,
		YSize:         11,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(11, 15, "B석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 12),
		XSize:         4,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

 	base = model.NewSeatBase(22, 15, "B석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 12, 12),
		XSize:         4,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect := eraser.Rectangle(pos.Absolute{X: 33, Y: 11}, pos.Absolute{X: 33, Y: 14})
	integrated3 := eraser.Integrated(Crect)
	base = model.NewSeatBase(27, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         7,
		YSize:         11,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 18, "A석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 14),
		XSize:         11,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(27, 18, "C석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 14),
		XSize:         7,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(1, 4, "L석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('L'),
	}

	base = model.NewSeatBase(35, 4, "R석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('R'),
	}

	Wrect := eraser.Rectangle(pos.Absolute{X: 10, Y: 15}, pos.Absolute{X: 26, Y: 15})
	integrated4 := eraser.Integrated(Wrect)
	base = model.NewSeatBase(7, 15, "휠체어석")
	blockInput10 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         23,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}
	
	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	block8 := group.VerticalBlock(blockInput8)
	block9 := group.VerticalBlock(blockInput9)
	block10 := group.HorizontalBlock(blockInput10)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8, block9, block10)
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
		XSize:           35,
		YSize:           15,
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
