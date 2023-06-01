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
	hall := eraser.HorizontalHallway(14)
	rectAsection := eraser.Rectangle(pos.Absolute{X: 3,Y: 4}, pos.Absolute{X: 3,Y: 9})
	rectAsection2 := eraser.Rectangle(pos.Absolute{X: 4,Y: 4}, pos.Absolute{X: 4,Y: 6})
	rectAsection3 := eraser.Rectangle(pos.Absolute{X: 3,Y: 15}, pos.Absolute{X: 9,Y: 16})
	rectAsection4 := eraser.Rectangle(pos.Absolute{X: 10,Y: 21}, pos.Absolute{X: 11,Y: 21})

	specificA := eraser.Position(pos.Absolute{X: 5, Y: 4})

	integrated := eraser.Integrated(hall, rectAsection, rectAsection2, rectAsection3, rectAsection4, specificA)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         18,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	specificB := eraser.Position(pos.Absolute{X: 27, Y: 5},
								 pos.Absolute{X: 27,Y: 7},
								 pos.Absolute{X: 27,Y: 9},
								 pos.Absolute{X: 27,Y: 11},
								 pos.Absolute{X: 27,Y: 13},
								 pos.Absolute{X: 27,Y: 16},
								 pos.Absolute{X: 27,Y: 18},
								 pos.Absolute{X: 27,Y: 20})
	rectBsection1 := eraser.Rectangle(pos.Absolute{X: 25,Y: 21}, pos.Absolute{X: 27,Y: 21})

	integrated2 := eraser.Integrated(hall, specificB, rectBsection1)
	base = model.NewSeatBase(13, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         15,
		YSize:         18,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}


	rectCsection1 := eraser.Rectangle(pos.Absolute{X: 37, Y: 4}, pos.Absolute{X: 37, Y: 9})
	rectCsection2 := eraser.Rectangle(pos.Absolute{X: 36, Y: 4}, pos.Absolute{X: 36, Y: 6})
	rectCsection3 := eraser.Rectangle(pos.Absolute{X: 31, Y: 15}, pos.Absolute{X: 37, Y: 16})
	rectCsection4 := eraser.Rectangle(pos.Absolute{X: 29, Y: 21}, pos.Absolute{X: 30, Y: 21})
	
	specificC := eraser.Position(pos.Absolute{X: 35,Y: 4})

	integrated3 := eraser.Integrated(hall, rectCsection1, rectCsection2, rectCsection3,rectCsection4, specificC)
	base = model.NewSeatBase(29, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         9,
		YSize:         18,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	emptySection := eraser.VerticalHallway(4,6,8, 32,34,36)
	rect1 := eraser.Rectangle(pos.Absolute{X: 10,Y: 16}, pos.Absolute{X: 30,Y: 16})
	integrated4 := eraser.Integrated(emptySection, rect1)
	base = model.NewSeatBase(3, 16, "휠체어석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         35,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	return group.Mixed(block1, block2, block3, block4)
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
