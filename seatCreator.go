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
	hall := eraser.HorizontalHallway(4,14)
	rect1 := eraser.Rectangle(pos.Absolute{X: 10, Y: 4}, pos.Absolute{X: 12, Y: 13})
	rect2 := eraser.Rectangle(pos.Absolute{X: 19, Y: 4}, pos.Absolute{X: 21, Y: 13})
	rect3 := eraser.Rectangle(pos.Absolute{X: 9, Y: 15}, pos.Absolute{X: 9, Y: 24})
	rect4 := eraser.Rectangle(pos.Absolute{X: 22, Y: 15}, pos.Absolute{X: 22, Y: 24})

	specific1 := eraser.Position(pos.Absolute{X: 8, Y: 15})
	specific2 := eraser.Position(pos.Absolute{X: 23, Y: 15})

	rect5 := eraser.Rectangle(pos.Absolute{X: 20, Y: 15},pos.Absolute{X:21,Y: 15})
	specific3 := eraser.Position(pos.Absolute{X: 21, Y: 17})
	specific4 := eraser.Position(pos.Absolute{X: 21, Y: 19})
	specific5 := eraser.Position(pos.Absolute{X: 21, Y: 21})

	rect6 := eraser.Rectangle(pos.Absolute{X:10 , Y: 23},pos.Absolute{X:21,Y: 24})
	rect7 := eraser.Rectangle(pos.Absolute{X:3 , Y: 12},pos.Absolute{X:9,Y: 13})
	rect8 := eraser.Rectangle(pos.Absolute{X:3 , Y: 24},pos.Absolute{X:10,Y: 24})

	integrated := eraser.Integrated(hall, rect1, rect2, rect3, rect4, specific1, specific2,rect5,specific3,specific4,specific5, rect6, rect7, rect8)
	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         26,
		YSize:         21,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	extraRect2 := eraser.Rectangle(pos.Absolute{X: 1,Y: 1},pos.Absolute{X: 1,Y: 1})
	integrated2 := eraser.Integrated(extraRect2)

	base = model.NewSeatBase(23, 23, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 18, 18),
		XSize:         6,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	extraRect3 := eraser.Rectangle(pos.Absolute{X: 1,Y: 1},pos.Absolute{X: 1,Y: 1})
	integrated3 := eraser.Integrated(extraRect3)

	base = model.NewSeatBase(23, 24, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 19, 19),
		XSize:         6,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}


	extraRect4 := eraser.Rectangle(pos.Absolute{X: 1,Y: 1},pos.Absolute{X: 1,Y: 1})
	integrated4 := eraser.Integrated(extraRect4)

	base = model.NewSeatBase(25, 25, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 21, 20),
		XSize:         4,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	extraRect5 := eraser.Rectangle(pos.Absolute{X: 1,Y: 1},pos.Absolute{X: 1,Y: 1})
	integrated5 := eraser.Integrated(extraRect5)

	base = model.NewSeatBase(3, 12, "휠체어석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 23),
		XSize:         5,
		YSize:         1,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	return group.Mixed(block1,block2,block3,block4, block5)
}

func ConcertHall1F() model.Group {
	hall1 := eraser.VerticalHallway(14,35)
	hall2 := eraser.HorizontalHallway(16,17,18)

	rect1 := eraser.Rectangle(pos.Absolute{X: 2, Y: 25}, pos.Absolute{X: 14, Y: 26})
	rect2 := eraser.Rectangle(pos.Absolute{X: 35, Y: 25}, pos.Absolute{X: 47, Y: 26})

	integrated := eraser.Integrated(hall1,hall2,rect1,rect2)

	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(2, 8, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         46,
		YSize:         17,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	integrated2 := eraser.Integrated()

	base = model.NewSeatBase(15, 25, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 13, 15),
		XSize:         20,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	rect3 := eraser.Rectangle(pos.Absolute{X: 8, Y: 25}, pos.Absolute{X: 35, Y: 26})
	rect4 := eraser.Rectangle(pos.Absolute{X: 41, Y: 25}, pos.Absolute{X: 48, Y: 26})
	integrated3 := eraser.Integrated(rect3,rect4)

	base = model.NewSeatBase(2, 25, "휠체어석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 22),
		XSize:         46,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Floor(nameformatter.Prefix('W'), 2),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	return group.Mixed(block1, block2, block3)
}

func ConcertHall2F() model.Group {
	vHall := eraser.VerticalHallway(7,27)
	integrated := eraser.Integrated(vHall)
	nameFormatter := nameformatter.Standard()

	base := model.NewSeatBase(2, 8, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         31,
		YSize:         7,
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
	seats := ConcertHall2F() // Put Seating Here
	target := area.Area{
		Key:             "2F",
		Seats:           seats,
		XSize:           20,
		YSize:           12,
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
