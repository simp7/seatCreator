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

	Hhall := eraser.HorizontalHallway(16)
	Arect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 7}, pos.Absolute{X: 6, Y: 10})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 3, Y: 10}, pos.Absolute{X: 5, Y: 13})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 13}, pos.Absolute{X: 4, Y: 15})

	Arect4 := eraser.Rectangle(pos.Absolute{X: 3,Y: 17}, pos.Absolute{X: 4, Y: 17})
	Arect5 := eraser.Rectangle(pos.Absolute{X: 3,Y: 18}, pos.Absolute{X: 3, Y: 21})

	Ahall := eraser.Rectangle(pos.Absolute{X: 3, Y: 22}, pos.Absolute{X: 13, Y: 23})

	Arect6 := eraser.Rectangle(pos.Absolute{X: 3, Y: 25}, pos.Absolute{X: 4, Y: 25})
	Arect7 := eraser.Rectangle(pos.Absolute{X: 3, Y: 26}, pos.Absolute{X: 8, Y: 26})

	integrated1 := eraser.Integrated(Hhall, Arect1,Arect2,Arect3, Arect4, Arect5,
									Ahall,Arect6,Arect7 )
	base := model.NewSeatBase(3, 7, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         20,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 15, Y: 7}, pos.Absolute{X: 15, Y: 20})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 32, Y: 7}, pos.Absolute{X: 32, Y: 18})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 16, Y: 7}, pos.Absolute{X: 16, Y: 16})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 31, Y: 7}, pos.Absolute{X: 31, Y: 14})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 17, Y: 7}, pos.Absolute{X: 17, Y: 12})
	Brect6 := eraser.Rectangle(pos.Absolute{X: 30, Y: 7}, pos.Absolute{X: 30, Y: 10})
	Brect7 := eraser.Rectangle(pos.Absolute{X: 18, Y: 7}, pos.Absolute{X: 18, Y: 8})

	Brect8 := eraser.Rectangle(pos.Absolute{X: 15, Y: 22}, pos.Absolute{X: 16, Y: 25})
	Brect9 := eraser.Rectangle(pos.Absolute{X: 31, Y: 22}, pos.Absolute{X: 32, Y: 25})
	
	Brect10 := eraser.Rectangle(pos.Absolute{X: 18, Y: 26}, pos.Absolute{X: 32, Y: 26})

	integrated2 := eraser.Integrated(Hhall, Brect1,Brect2,Brect3,Brect4,Brect5,Brect6,Brect7,
									Brect8,Brect9,Brect10)
	base = model.NewSeatBase(15, 7, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         18,
		YSize:         20,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	integrated4 := eraser.Integrated()
	base = model.NewSeatBase(30, 26, "B석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 12, 1),
		XSize:         3,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}


	Crect1 := eraser.Rectangle(pos.Absolute{X: 41, Y: 7}, pos.Absolute{X: 44, Y: 10})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 42, Y: 11}, pos.Absolute{X: 44, Y: 13})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 43, Y: 14}, pos.Absolute{X: 44, Y: 17})
	Crect4 := eraser.Rectangle(pos.Absolute{X: 44, Y: 18}, pos.Absolute{X: 44, Y: 21})

	Chall := eraser.Rectangle(pos.Absolute{X: 24, Y: 22}, pos.Absolute{X: 44, Y: 23})

	Crect5 := eraser.Rectangle(pos.Absolute{X: 43, Y: 25}, pos.Absolute{X: 44, Y: 25})
	Crect6 := eraser.Rectangle(pos.Absolute{X: 39, Y: 26}, pos.Absolute{X: 44, Y: 26})


	integrated3 := eraser.Integrated(Hhall, Crect1,Crect2,Crect3,Crect4,Chall, Crect5, Crect6)
	base = model.NewSeatBase(34, 7, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         11,
		YSize:         20,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	rect := eraser.Rectangle(pos.Absolute{X: 14, Y: 22}, pos.Absolute{X: 33, Y: 22})
	integrated5 := eraser.Integrated(rect)
	base = model.NewSeatBase(8, 22, "휠체어석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         32,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Prefix('W'),
	}

	OPspecific := eraser.Position(pos.Absolute{X: 14, Y: 3}, pos.Absolute{X: 34,Y: 3})
	integrated6 := eraser.Integrated(OPspecific)
	base = model.NewSeatBase(14, 3, "OP석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         21,
		YSize:         3,
		EmptyChecker:  integrated6,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)

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
		Key:             "1F",
		Seats:           seats,
		XSize:           40,
		YSize:           24,
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
