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
	hall := eraser.HorizontalHallway(11)
	hall2 := eraser.VerticalHallway(17,31)

	rectLeft1 := eraser.Rectangle(pos.Absolute{X: 3, Y: 4}, pos.Absolute{X: 3, Y: 8})
	rectLeft2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 4}, pos.Absolute{X: 4, Y: 6})
	rectLeft3 := eraser.Rectangle(pos.Absolute{X: 5, Y: 4}, pos.Absolute{X: 5, Y: 5})
	specificRight1 := eraser.Position(pos.Absolute{X: 6,Y: 4})

	rectLeft4 := eraser.Rectangle(pos.Absolute{X: 3,Y: 12}, pos.Absolute{X: 3,Y: 18})
	rectLeft5 := eraser.Rectangle(pos.Absolute{X: 4,Y: 12}, pos.Absolute{X: 4,Y: 14})
	specificRight2 := eraser.Position(pos.Absolute{X: 5,Y: 12})

	rectRight1 := eraser.Rectangle(pos.Absolute{X: 45,Y: 4},pos.Absolute{X: 45,Y: 8})
	rectRight2 := eraser.Rectangle(pos.Absolute{X: 44,Y: 4},pos.Absolute{X: 44,Y: 6})
	rectRight3 := eraser.Rectangle(pos.Absolute{X: 42,Y: 4},pos.Absolute{X: 43,Y: 4})

	rectRight4 := eraser.Rectangle(pos.Absolute{X: 44,Y: 12},pos.Absolute{X: 45,Y: 18})
	rectRight5 := eraser.Rectangle(pos.Absolute{X: 43,Y: 12},pos.Absolute{X: 44,Y: 13})

	rectBottom1 := eraser.Rectangle(pos.Absolute{X: 17,Y: 17}, pos.Absolute{X: 30, Y: 18})
	rectBottom2 := eraser.Rectangle(pos.Absolute{X: 31,Y: 18}, pos.Absolute{X: 38,Y: 18})


	integrated := eraser.Integrated(hall, hall2, rectLeft1, rectLeft2,rectLeft3,specificRight1,
									rectLeft4, rectLeft5, specificRight2, rectRight1, rectRight2,rectRight3,
									rectRight4,rectRight5, rectBottom1, rectBottom2)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         43,
		YSize:         15,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	block1 := group.HorizontalBlock(blockInput)
	return group.Mixed(block1)
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
		XSize:           43,
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
