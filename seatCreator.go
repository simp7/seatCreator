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


	Arect1 := eraser.Rectangle(pos.Absolute{X: 5, Y: 16}, pos.Absolute{X: 5, Y: 17})
	Brect1 := eraser.Rectangle(pos.Absolute{X: 31, Y: 13}, pos.Absolute{X: 31, Y: 15})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 11, Y: 17}, pos.Absolute{X: 14, Y: 17})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 27, Y: 17}, pos.Absolute{X: 31, Y: 17})
	Crect1 := eraser.Rectangle(pos.Absolute{X: 38, Y: 16}, pos.Absolute{X: 38, Y: 17})
	integrated := eraser.Integrated(Arect1, Brect1, Crect1, Brect2, Brect3)

	base := model.NewSeatBase(5, 13, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         5,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(12, 13, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         5,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	

	base = model.NewSeatBase(33, 13, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         6,
		YSize:         5,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}
	
	rect6 := eraser.Rectangle(pos.Absolute{X: 15, Y: 17}, pos.Absolute{X: 26, Y: 17})
	integrated2 := eraser.Integrated(rect6)
	base = model.NewSeatBase(13, 17, "휠체어석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         16,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('W'),
	}

	rect7 := eraser.Rectangle(pos.Absolute{X: 4, Y: 12}, pos.Absolute{X: 4, Y: 15})
	rect8 := eraser.Rectangle(pos.Absolute{X: 39, Y: 12}, pos.Absolute{X: 39, Y: 15})
	integrated3 := eraser.Integrated(rect7, rect8)
	base = model.NewSeatBase(3, 3, "BOX석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         2,
		YSize:         13,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(39, 3, "BOX석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 3),
		XSize:         2,
		YSize:         13,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}
	
	
	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.VerticalBlock(blockInput6)
	return group.Mixed(block1, block2, block3, block4, block5, block6)
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
		XSize:           38,
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
