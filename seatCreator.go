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

	Arect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 5}, pos.Absolute{X: 3, Y: 13})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 6}, pos.Absolute{X: 4, Y: 13})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 5, Y: 8}, pos.Absolute{X: 5, Y: 13})
	Arect4 := eraser.Rectangle(pos.Absolute{X: 6, Y: 10}, pos.Absolute{X: 6, Y: 13})
	Arect5 := eraser.Rectangle(pos.Absolute{X: 7, Y: 11}, pos.Absolute{X: 7, Y: 13})
	Arect6 := eraser.Rectangle(pos.Absolute{X: 8, Y: 13}, pos.Absolute{X: 8, Y: 13})
	integrated := eraser.Integrated(Arect1, Arect2, Arect3,Arect4,Arect5,Arect6)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         10,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 20, Y: 4}, pos.Absolute{X: 20, Y: 12})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 45, Y: 4}, pos.Absolute{X: 45, Y: 11})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 21, Y: 4}, pos.Absolute{X: 21, Y: 10})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 44, Y: 4}, pos.Absolute{X: 44, Y: 9})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 22, Y: 4}, pos.Absolute{X: 22, Y: 8})
	Brect6 := eraser.Rectangle(pos.Absolute{X: 43, Y: 4}, pos.Absolute{X: 43, Y: 7})
	Brect7 := eraser.Rectangle(pos.Absolute{X: 23, Y: 4}, pos.Absolute{X: 23, Y: 6})
	Brect8 := eraser.Rectangle(pos.Absolute{X: 42, Y: 4}, pos.Absolute{X: 42, Y: 5})
	Brect9 := eraser.Rectangle(pos.Absolute{X: 24, Y: 4}, pos.Absolute{X: 24, Y: 5})
	Bspecific := eraser.Position(pos.Absolute{X: 41, Y: 4})
	integrated2 := eraser.Integrated(Brect1,Brect2,Brect3,Brect4,Brect5,Brect6,Brect7,Brect8,Brect9, Bspecific)
	base = model.NewSeatBase(20, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         26,
		YSize:         10,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X: 62, Y: 5}, pos.Absolute{X: 62, Y: 13})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 61, Y: 6}, pos.Absolute{X: 61, Y: 13})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 60, Y: 8}, pos.Absolute{X: 60, Y: 13})
	Crect4 := eraser.Rectangle(pos.Absolute{X: 59, Y: 10}, pos.Absolute{X: 59, Y: 13})
	Crect5 := eraser.Rectangle(pos.Absolute{X: 58, Y: 11}, pos.Absolute{X: 58, Y: 13})
	Crect6 := eraser.Rectangle(pos.Absolute{X: 57, Y: 13}, pos.Absolute{X: 57, Y: 13})
	integrated3 := eraser.Integrated(Crect1,Crect2,Crect3,Crect4,Crect5,Crect6)
	base = model.NewSeatBase(47, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         10,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
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
