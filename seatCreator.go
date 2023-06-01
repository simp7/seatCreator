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

	rect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 4},pos.Absolute{X: 5,Y: 4})
	rect2 := eraser.Rectangle(pos.Absolute{X: 3,Y: 5},pos.Absolute{X: 4,Y: 6})
	rect3 := eraser.Rectangle(pos.Absolute{X: 3,Y: 7},pos.Absolute{X: 3,Y: 8})
	integrated := eraser.Integrated(rect1, rect2, rect3)
	base := model.NewSeatBase(3, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('A'),
	}

	rect4 := eraser.Rectangle(pos.Absolute{X: 22,Y: 4}, pos.Absolute{X: 22,Y: 6})
	integrated2 := eraser.Integrated(rect4)
	base = model.NewSeatBase(9, 4, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         9,
		EmptyChecker:  integrated2,
		NameFormatter: nameformatter.Prefix('B'),
	}

	rect5 := eraser.Rectangle(pos.Absolute{X: 26,Y: 4},pos.Absolute{X: 28,Y: 4})
	rect6 := eraser.Rectangle(pos.Absolute{X: 27,Y: 5},pos.Absolute{X: 28,Y: 6})
	rect7 := eraser.Rectangle(pos.Absolute{X: 28,Y: 7},pos.Absolute{X: 28,Y: 8})
	integrated3 := eraser.Integrated(rect5,rect6,rect7)
	base = model.NewSeatBase(24, 4, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         9,
		EmptyChecker:  integrated3,
		NameFormatter: nameformatter.Prefix('C'),
	}

	rect8 := eraser.Rectangle(pos.Absolute{X: 3,Y: 22}, pos.Absolute{X: 3,Y: 23})
	integrated4 := eraser.Integrated(rect8)
	base = model.NewSeatBase(3, 14, "D석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         10,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('D'),
	}
	base = model.NewSeatBase(9, 14, "E석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         9,
		EmptyChecker:  integrated,
		NameFormatter: nameformatter.Prefix('E'),
	}

	rect9 := eraser.Rectangle(pos.Absolute{X: 28, Y: 22}, pos.Absolute{X: 28, Y: 23})
	integrated5 := eraser.Integrated(rect9)
	base = model.NewSeatBase(24, 14, "F석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         5,
		YSize:         10,
		EmptyChecker:  integrated5,
		NameFormatter: nameformatter.Prefix('F'),
	}

	specific1 := eraser.Position(pos.Absolute{X: 8,Y: 2}, pos.Absolute{X: 22,Y: 2}, pos.Absolute{X: 23,Y: 2})
	integrated6 := eraser.Integrated(specific1)
	base = model.NewSeatBase(6, 2, "O석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         20,
		YSize:         1,
		EmptyChecker:  integrated6,
		NameFormatter: nameformatter.Prefix('O'),
	}

	rect10 := eraser.Rectangle(pos.Absolute{X: 8,Y: 24}, pos.Absolute{X: 23,Y: 24})
	integrated7 := eraser.Integrated(rect10)
	base = model.NewSeatBase(4, 24, "휠체어석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         24,
		YSize:         1,
		EmptyChecker:  integrated7,
		NameFormatter: nameformatter.Prefix('W'),
	}

	rect11 := eraser.Rectangle(pos.Absolute{X: 11,Y: 23},pos.Absolute{X: 20,Y: 23})
	integrated8 := eraser.Integrated(rect11)
	base = model.NewSeatBase(9, 23, "휠체어석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 9, 23),
		XSize:         14,
		YSize:         1,
		EmptyChecker:  integrated8,
		NameFormatter: nameformatter.Prefix('W'),
	}

	block1 := group.HorizontalBlock(blockInput)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.HorizontalBlock(blockInput5)
	block6 := group.HorizontalBlock(blockInput6)
	block7 := group.HorizontalBlock(blockInput7)
	block8 := group.HorizontalBlock(blockInput8)
	block9 := group.HorizontalBlock(blockInput9)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8, block9)
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
		XSize:           26,
		YSize:           20,
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