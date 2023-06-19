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

	Arect1 := eraser.Rectangle(pos.Absolute{X: 3,Y: 25}, pos.Absolute{X: 3, Y: 27})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 16,Y: 25}, pos.Absolute{X: 16, Y: 26})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 4,Y: 25}, pos.Absolute{X: 4, Y: 25})
	Arect4 := eraser.Rectangle(pos.Absolute{X: 9,Y: 29}, pos.Absolute{X: 10, Y: 29})
	integrated1 := eraser.Integrated(Arect1,Arect2,Arect3, Arect4)
	base := model.NewSeatBase(3, 25, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 18, Y: 25}, pos.Absolute{X: 18, Y: 27})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 36, Y: 25}, pos.Absolute{X: 36, Y: 26})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 19, Y: 25}, pos.Absolute{X: 19, Y: 25})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 35, Y: 29}, pos.Absolute{X: 36, Y: 29})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 18, Y: 29}, pos.Absolute{X: 19, Y: 29})
	Bspecific := eraser.Position(pos.Absolute{X: 26, Y: 29})
	integrated2 := eraser.Integrated(Brect1,Brect2,Brect3, Brect4,Brect5, Bspecific)
	base = model.NewSeatBase(18, 25, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         19,
		YSize:         5,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X: 38 ,Y: 25}, pos.Absolute{X: 38, Y: 27})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 51 ,Y: 25}, pos.Absolute{X: 51, Y: 26})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 39 ,Y: 25}, pos.Absolute{X: 39, Y: 25})
	Crect4 := eraser.Rectangle(pos.Absolute{X: 44 ,Y: 29}, pos.Absolute{X: 45, Y: 29})
	integrated3 := eraser.Integrated(Crect1, Crect2,Crect3,Crect4)
	base = model.NewSeatBase(38, 25, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         14,
		YSize:         5,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(4, 2, "L4석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(4, 8, "L5석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(4, 13, "L6석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         11,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(3, 16, "L6석")
	blockInput11 := group.BlockInput{
		Criteria:      model.NewSeat(base, 12, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(50, 2, "R4석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         5,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(50, 8, "R5석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(50, 13, "R6석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         11,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	base = model.NewSeatBase(51, 16, "R6석")
	blockInput10 := group.BlockInput{
		Criteria:      model.NewSeat(base, 12, 1),
		XSize:         1,
		YSize:         8,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}


	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.VerticalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.VerticalBlock(blockInput6)
	block7 := group.VerticalBlock(blockInput7)
	block8 := group.VerticalBlock(blockInput8)
	block9 := group.VerticalBlock(blockInput9)
	block10 := group.VerticalBlock(blockInput10)
	block11 := group.VerticalBlock(blockInput11)

	return group.Mixed(block1, block2, block3, block4, block5,block6,block7,block8,block9,block10, block11)
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
		XSize:           44,
		YSize:           29,
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
