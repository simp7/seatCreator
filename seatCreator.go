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

	Drect1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 3}, pos.Absolute{X: 4, Y: 3})
	Drect2 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 3, Y: 5})
	Drect3 := eraser.Rectangle(pos.Absolute{X: 3, Y: 10}, pos.Absolute{X: 4, Y: 10})
	Drect4 := eraser.Rectangle(pos.Absolute{X: 3, Y: 11}, pos.Absolute{X: 6, Y: 11})
	Drect5 := eraser.Rectangle(pos.Absolute{X: 3, Y: 12}, pos.Absolute{X: 6, Y: 12})

	Drect6 := eraser.Rectangle(pos.Absolute{X: 13, Y: 11}, pos.Absolute{X: 14, Y: 11})
	Drect7 := eraser.Rectangle(pos.Absolute{X: 14, Y: 12}, pos.Absolute{X: 14, Y: 12})
	integrated1 := eraser.Integrated(Drect1, Drect2,Drect3,Drect4,Drect5,Drect6,Drect7)
	base := model.NewSeatBase(3, 3, "D석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         10,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	Erect1 := eraser.Rectangle(pos.Absolute{X: 16,Y: 3}, pos.Absolute{X: 16, Y: 8})
	Erect2 := eraser.Rectangle(pos.Absolute{X: 35,Y: 3}, pos.Absolute{X: 35, Y: 6})
	Erect3 := eraser.Rectangle(pos.Absolute{X: 17,Y: 3}, pos.Absolute{X: 17, Y: 4})

	Erect4 := eraser.Rectangle(pos.Absolute{X: 16,Y: 11}, pos.Absolute{X: 18, Y:11})
	Erect5 := eraser.Rectangle(pos.Absolute{X: 32,Y: 11}, pos.Absolute{X: 35, Y: 11})

	Erect6 := eraser.Rectangle(pos.Absolute{X: 16,Y: 12}, pos.Absolute{X: 17, Y: 12})
	Erect7 := eraser.Rectangle(pos.Absolute{X: 33,Y: 12}, pos.Absolute{X: 35, Y: 12})
	integrated2 := eraser.Integrated(Erect1,Erect2,Erect3,Erect4,Erect5,Erect6,Erect7)
	base = model.NewSeatBase(16, 3, "E석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         10,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Frect1 := eraser.Rectangle(pos.Absolute{X: 47, Y: 3}, pos.Absolute{X: 49,Y: 3})
	Frect2 := eraser.Rectangle(pos.Absolute{X: 48, Y: 4}, pos.Absolute{X: 49,Y: 5})
	Frect3 := eraser.Rectangle(pos.Absolute{X: 47, Y: 10}, pos.Absolute{X: 49,Y: 10})

	Frect4 := eraser.Rectangle(pos.Absolute{X: 37, Y: 11}, pos.Absolute{X: 38,Y: 11})
	Frect5 := eraser.Rectangle(pos.Absolute{X: 45, Y: 11}, pos.Absolute{X: 49,Y: 11})

	Frect6 := eraser.Rectangle(pos.Absolute{X: 37, Y: 12}, pos.Absolute{X: 37,Y: 12})
	Frect7 := eraser.Rectangle(pos.Absolute{X: 45, Y: 12}, pos.Absolute{X: 49,Y: 12})
	integrated3 := eraser.Integrated(Frect1,Frect2,Frect3,Frect4,Frect5,Frect6,Frect7)
	base = model.NewSeatBase(37, 3, "F석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         12,
		YSize:         10,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput1)
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
		XSize:           42,
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
