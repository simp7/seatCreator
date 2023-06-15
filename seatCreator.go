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
	hallH := eraser.HorizontalHallway(18)
	

	Arect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 7}, pos.Absolute{X: 3, Y: 15})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 7}, pos.Absolute{X: 4, Y: 13})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 5, Y: 7}, pos.Absolute{X: 5, Y: 10})
	Arect4 := eraser.Rectangle(pos.Absolute{X: 6, Y: 7}, pos.Absolute{X: 6, Y: 7})
	integrated := eraser.Integrated(hallH, Arect1,Arect2,Arect3,Arect4)
	base := model.NewSeatBase(3, 7, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         25,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 13, Y: 7}, pos.Absolute{X: 13, Y: 16})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 14, Y: 7}, pos.Absolute{X: 14, Y: 14})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 15, Y: 7}, pos.Absolute{X: 15, Y: 12})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 16, Y: 7}, pos.Absolute{X: 16, Y: 10})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 17, Y: 7}, pos.Absolute{X: 17, Y: 8})
	integrated2 := eraser.Integrated(hallH, Brect1,Brect2,Brect3,Brect4,Brect5)
	base = model.NewSeatBase(13, 7, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         25,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect := eraser.Rectangle(pos.Absolute{X: 24, Y: 31}, pos.Absolute{X: 36, Y: 31})
	integrated3 := eraser.Integrated(hallH, Crect)
	base = model.NewSeatBase(24, 7, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         13,
		YSize:         25,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}
	
	Drect1 := eraser.Rectangle(pos.Absolute{X: 47, Y: 7}, pos.Absolute{X: 47, Y: 16})
	Drect2 := eraser.Rectangle(pos.Absolute{X: 46, Y: 7}, pos.Absolute{X: 46, Y: 14})
	Drect3 := eraser.Rectangle(pos.Absolute{X: 45, Y: 7}, pos.Absolute{X: 45, Y: 12})
	Drect4 := eraser.Rectangle(pos.Absolute{X: 44, Y: 7}, pos.Absolute{X: 44, Y: 10})
	Drect5 := eraser.Rectangle(pos.Absolute{X: 43, Y: 7}, pos.Absolute{X: 43, Y: 8})
	integrated4 := eraser.Integrated(hallH,Drect1,Drect2,Drect3,Drect4,Drect5)
	base = model.NewSeatBase(38, 7, "D석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         10,
		YSize:         25,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	Erect1 := eraser.Rectangle(pos.Absolute{X: 57, Y: 7},pos.Absolute{X: 57,Y: 15})
	Erect2 := eraser.Rectangle(pos.Absolute{X: 56, Y: 7},pos.Absolute{X: 56,Y: 13})
	Erect3 := eraser.Rectangle(pos.Absolute{X: 55, Y: 7},pos.Absolute{X: 55,Y: 10})
	Erect4 := eraser.Rectangle(pos.Absolute{X: 54, Y: 7},pos.Absolute{X: 54,Y: 7})
	integrated5 := eraser.Integrated(hallH, Erect1,Erect2,Erect3,Erect4)
	base = model.NewSeatBase(49, 7, "E석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         25,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	hallV := eraser.VerticalHallway(24,36)
	Orect1 := eraser.Rectangle(pos.Absolute{X: 14, Y: 2}, pos.Absolute{X: 14, Y: 4})
	Orect2 := eraser.Rectangle(pos.Absolute{X: 15, Y: 2}, pos.Absolute{X: 15, Y: 3})
	Orect3 := eraser.Rectangle(pos.Absolute{X: 16, Y: 2}, pos.Absolute{X: 16, Y: 2})

	Orect4 := eraser.Rectangle(pos.Absolute{X: 46, Y: 2}, pos.Absolute{X: 46, Y: 4})
	Orect5 := eraser.Rectangle(pos.Absolute{X: 45, Y: 2}, pos.Absolute{X: 45, Y: 3})
	Orect6 := eraser.Rectangle(pos.Absolute{X: 44, Y: 2}, pos.Absolute{X: 44, Y: 2})

	integrated6 := eraser.Integrated(hallV, Orect1,Orect2,Orect3,Orect4,Orect5,Orect6)
	base = model.NewSeatBase(14, 2, "O석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         33,
		YSize:         4,
		EmptyChecker:  integrated6,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput)
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
		XSize:           55,
		YSize:           30,
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
