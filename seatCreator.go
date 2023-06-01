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

	rect4 := eraser.Rectangle(pos.Absolute{X: 11, Y: 4}, pos.Absolute{X: 15, Y: 4})
	rect5 := eraser.Rectangle(pos.Absolute{X: 27, Y: 4}, pos.Absolute{X: 31, Y: 4})

	rect6 := eraser.Rectangle(pos.Absolute{X: 11, Y: 5}, pos.Absolute{X: 15, Y: 5})
	rect7 := eraser.Rectangle(pos.Absolute{X: 28, Y: 5}, pos.Absolute{X: 31, Y: 5})

	rect8 := eraser.Rectangle(pos.Absolute{X: 11, Y: 6}, pos.Absolute{X: 14, Y: 6})
	rect9 := eraser.Rectangle(pos.Absolute{X: 28, Y: 6}, pos.Absolute{X: 31, Y: 6})

	rect10 := eraser.Rectangle(pos.Absolute{X: 11, Y: 7}, pos.Absolute{X: 14, Y: 7})
	rect11 := eraser.Rectangle(pos.Absolute{X: 29, Y: 7}, pos.Absolute{X: 31, Y: 7})

	rect12 := eraser.Rectangle(pos.Absolute{X: 11, Y: 8}, pos.Absolute{X: 13, Y: 8})
	rect13 := eraser.Rectangle(pos.Absolute{X: 29, Y: 8}, pos.Absolute{X: 31, Y: 8})

	rect14 := eraser.Rectangle(pos.Absolute{X: 11, Y: 9}, pos.Absolute{X: 13, Y: 9})
	rect15 := eraser.Rectangle(pos.Absolute{X: 30, Y: 9}, pos.Absolute{X: 31, Y: 9})

	rect16 := eraser.Rectangle(pos.Absolute{X: 11, Y: 10}, pos.Absolute{X: 12, Y: 10})
	rect17 := eraser.Rectangle(pos.Absolute{X: 30, Y: 10}, pos.Absolute{X: 31, Y: 10})

	rect18 := eraser.Rectangle(pos.Absolute{X: 11, Y: 11}, pos.Absolute{X: 12, Y: 11})
	rect19 := eraser.Rectangle(pos.Absolute{X: 31, Y: 11}, pos.Absolute{X: 31, Y: 11})

	rect20 := eraser.Rectangle(pos.Absolute{X: 11, Y: 12}, pos.Absolute{X: 11, Y: 12})
	rect21 := eraser.Rectangle(pos.Absolute{X: 31, Y: 12}, pos.Absolute{X: 31, Y: 12})

	rect22 := eraser.Rectangle(pos.Absolute{X: 11, Y: 13}, pos.Absolute{X: 11, Y: 13})
	rect23 := eraser.Rectangle(pos.Absolute{X: 32, Y: 13}, pos.Absolute{X: 31, Y: 13})

	rect24 := eraser.Rectangle(pos.Absolute{X: 14,Y: 15}, pos.Absolute{X: 31, Y: 15})

	integrated := eraser.Integrated(rect4, rect5, rect6, rect7, rect8, rect9, rect10,
			rect11, rect12, rect13, rect14, rect15, rect16, rect17, rect18, rect19, rect20,
			rect21, rect22, rect23, rect24)
	base := model.NewSeatBase(11, 4, "A석")
	blockInput := group.BlockInput{
		Criteria:      model.NewSeat(base, 8, 1),
		XSize:         21,
		YSize:         12,
		EmptyChecker:  integrated,
		NameFormatter: nameFormatter,
	}

	rect1 := eraser.Rectangle(pos.Absolute{X: 39,Y: 10}, pos.Absolute{X: 39,Y: 11})
	rect2 := eraser.Rectangle(pos.Absolute{X: 38,Y: 12}, pos.Absolute{X: 39,Y: 13})
	rect3 := eraser.Rectangle(pos.Absolute{X: 37,Y: 14}, pos.Absolute{X: 39,Y: 15})
	rect25 := eraser.Rectangle(pos.Absolute{X: 34,Y: 4},pos.Absolute{X: 39,Y: 4})
	integrated2 := eraser.Integrated(rect1, rect2, rect3, rect25)
	base = model.NewSeatBase(33, 4, "A석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 29, 1),
		XSize:         7,
		YSize:         12,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 6, "A석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 3),
		XSize:         7,
		YSize:         4,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(4, 10, "A석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 7),
		XSize:         6,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(5, 12, "A석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 3, 9),
		XSize:         5,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(6, 14, "A석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 4, 11),
		XSize:         4,
		YSize:         2,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(4, 5, "A석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 2, 2),
		XSize:         6,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(5, 4, "A석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 3, 1),
		XSize:         5,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(29, 15, "A석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 26, 12),
		XSize:         3,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(7, 16, "A석")
	blockInput10 := group.BlockInput{
		Criteria:      model.NewSeat(base, 5, 13),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(35, 16, "A석")
	blockInput11 := group.BlockInput{
		Criteria:      model.NewSeat(base, 31, 13),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	integrated3 := eraser.Integrated()
	base = model.NewSeatBase(38, 4, "A석")
	blockInput12 := group.BlockInput{
		Criteria:      model.NewSeat(base, 34, 1),
		XSize:         1,
		YSize:         1,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	rect26 := eraser.Rectangle(pos.Absolute{X: 9,Y: 16}, pos.Absolute{X: 33,Y: 16})
	integrated4 := eraser.Integrated(rect26)
	base = model.NewSeatBase(8, 16, "휠체어석")
	blockInput13 := group.BlockInput{
		Criteria:      model.NewSeat(base, 35, 1),
		XSize:         27,
		YSize:         1,
		EmptyChecker:  integrated4,
		NameFormatter: nameformatter.Prefix('W'),
	}

	base = model.NewSeatBase(35, 4, "휠체어석")
	blockInput14 := group.BlockInput{
		Criteria:      model.NewSeat(base, 35, 3),
		XSize:         2,
		YSize:         1,
		EmptyChecker:  integrated4,
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
	block10 := group.HorizontalBlock(blockInput10)
	block11 := group.HorizontalBlock(blockInput11)
	block12 := group.HorizontalBlock(blockInput12)
	block13 := group.HorizontalBlock(blockInput13)
	block14 := group.HorizontalBlock(blockInput14)
	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8, block9,
					block10, block11, block12, block13, block14)
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
		XSize:           37,
		YSize:           13,
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