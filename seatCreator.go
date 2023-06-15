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

	Arect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 3, Y: 10})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 4}, pos.Absolute{X: 4, Y: 6})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 17}, pos.Absolute{X: 3, Y: 21})
	integrated := eraser.Integrated(Arect1,Arect2,Arect3)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         18,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 18, Y: 4}, pos.Absolute{X: 18, Y: 18})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 35, Y: 4}, pos.Absolute{X: 35, Y: 15})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 19, Y: 4}, pos.Absolute{X: 19, Y: 13})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 34, Y: 4}, pos.Absolute{X: 34, Y: 11})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 20, Y: 4}, pos.Absolute{X: 20, Y: 9})
	Brect6 := eraser.Rectangle(pos.Absolute{X: 33, Y: 4}, pos.Absolute{X: 33, Y: 7})
	Brect7 := eraser.Rectangle(pos.Absolute{X: 21, Y: 4}, pos.Absolute{X: 21, Y: 4})
	integrated2 := eraser.Integrated(Brect1,Brect2,Brect3,Brect4,Brect5,Brect6,Brect7)
	base = model.NewSeatBase(18, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         18,
		YSize:         17,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X: 50, Y: 4}, pos.Absolute{X: 50, Y: 10})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 49, Y: 4}, pos.Absolute{X: 49, Y: 6})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 50, Y: 17}, pos.Absolute{X: 50, Y: 21})
	integrated3 := eraser.Integrated(Crect1,Crect2,Crect3)
	base = model.NewSeatBase(37, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         18,
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
		Key:             "1F",
		Seats:           seats,
		XSize:           48,
		YSize:           18,
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
