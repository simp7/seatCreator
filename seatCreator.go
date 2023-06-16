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
	hall := eraser.HorizontalHallway(23, 28)

	Arect1 := eraser.Rectangle(pos.Absolute{X: 4, Y: 8}, pos.Absolute{X: 6,Y: 8})
	Arect2 := eraser.Rectangle(pos.Absolute{X: 4, Y: 9}, pos.Absolute{X: 5,Y: 10})
	Arect3 := eraser.Rectangle(pos.Absolute{X: 4, Y: 11}, pos.Absolute{X: 4,Y: 12})
	Arect4 := eraser.Rectangle(pos.Absolute{X: 4, Y: 19}, pos.Absolute{X: 4,Y: 21})
	Arect5 := eraser.Rectangle(pos.Absolute{X: 4, Y: 22}, pos.Absolute{X: 5,Y: 22})
	Arect6 := eraser.Rectangle(pos.Absolute{X: 4, Y: 24}, pos.Absolute{X: 8,Y: 24})
	Arect7 := eraser.Rectangle(pos.Absolute{X: 4, Y: 25}, pos.Absolute{X: 9,Y: 25})
	Arect8 := eraser.Rectangle(pos.Absolute{X: 4, Y: 26}, pos.Absolute{X: 13,Y: 26})
	Arect9 := eraser.Rectangle(pos.Absolute{X: 4, Y: 27}, pos.Absolute{X: 14, Y: 27})
	Arect10 := eraser.Rectangle(pos.Absolute{X: 2, Y: 8}, pos.Absolute{X: 3, Y: 27})

	Arect11 := eraser.Rectangle(pos.Absolute{X: 1, Y: 29}, pos.Absolute{X: 4, Y: 29})
	Arect12 := eraser.Rectangle(pos.Absolute{X: 1, Y: 30}, pos.Absolute{X: 3, Y: 30})
	Arect13 := eraser.Rectangle(pos.Absolute{X: 1, Y: 31}, pos.Absolute{X: 2, Y: 31})
	Arect14 := eraser.Rectangle(pos.Absolute{X: 1, Y: 33}, pos.Absolute{X: 3, Y: 33})

	integrated1 := eraser.Integrated(hall,Arect1, Arect2, Arect3, Arect4, Arect5,
									Arect6,Arect7,Arect8, Arect9, Arect10,
									Arect11, Arect12, Arect13, Arect14)
	base := model.NewSeatBase(2, 8, "A석")
	blockInput1 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         13,
		YSize:         26,
		EmptyChecker:  integrated1,
		NameFormatter: nameFormatter,
	}

	Brect1 := eraser.Rectangle(pos.Absolute{X: 16, Y: 8}, pos.Absolute{X: 16, Y: 27})
	Brect2 := eraser.Rectangle(pos.Absolute{X: 33, Y: 8}, pos.Absolute{X: 33, Y: 27})
	Brect3 := eraser.Rectangle(pos.Absolute{X: 16, Y: 27}, pos.Absolute{X: 18, Y: 27})
	Brect4 := eraser.Rectangle(pos.Absolute{X: 31, Y: 27}, pos.Absolute{X: 33, Y: 27})
	Brect5 := eraser.Rectangle(pos.Absolute{X: 33, Y: 29}, pos.Absolute{X: 33, Y: 29})
	Brect6 := eraser.Rectangle(pos.Absolute{X: 21, Y: 31}, pos.Absolute{X: 28, Y: 33})
	Brect7 := eraser.Rectangle(pos.Absolute{X: 16, Y: 33}, pos.Absolute{X: 18, Y: 33})
	Brect8 := eraser.Rectangle(pos.Absolute{X: 31, Y: 33}, pos.Absolute{X: 33, Y: 33})

	integrated2 := eraser.Integrated(hall, Brect1, Brect2, Brect3, Brect4, Brect5,Brect6, Brect7,Brect8)
	base = model.NewSeatBase(16, 8, "B석")
	blockInput2 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         18,
		YSize:         26,
		EmptyChecker:  integrated2,
		NameFormatter: nameFormatter,
	}

	Crect1 := eraser.Rectangle(pos.Absolute{X: 43, Y: 8}, pos.Absolute{X: 47,Y: 8})
	Crect2 := eraser.Rectangle(pos.Absolute{X: 44, Y: 9}, pos.Absolute{X: 47,Y: 10})
	Crect3 := eraser.Rectangle(pos.Absolute{X: 45, Y: 11}, pos.Absolute{X: 47,Y: 12})
	Crect4 := eraser.Rectangle(pos.Absolute{X: 45, Y: 19}, pos.Absolute{X: 47,Y: 21})
	Crect5 := eraser.Rectangle(pos.Absolute{X: 44, Y: 22}, pos.Absolute{X: 47,Y: 22})
	Crect6 := eraser.Rectangle(pos.Absolute{X: 41, Y: 24}, pos.Absolute{X: 47,Y: 24})
	Crect7 := eraser.Rectangle(pos.Absolute{X: 40, Y: 25}, pos.Absolute{X: 47,Y: 25})
	Crect8 := eraser.Rectangle(pos.Absolute{X: 36, Y: 26}, pos.Absolute{X: 47,Y: 26})
	Crect9 := eraser.Rectangle(pos.Absolute{X: 35, Y: 27}, pos.Absolute{X: 47, Y: 27})

	Crect11 := eraser.Rectangle(pos.Absolute{X: 45, Y: 29}, pos.Absolute{X: 47, Y: 29})
	Crect12 := eraser.Rectangle(pos.Absolute{X: 46, Y: 30}, pos.Absolute{X: 47, Y: 30})
	Crect13 := eraser.Rectangle(pos.Absolute{X: 47, Y: 31}, pos.Absolute{X: 47, Y: 31})
	Crect14 := eraser.Rectangle(pos.Absolute{X: 46, Y: 33}, pos.Absolute{X: 47, Y: 33})
	
	Crect15 := eraser.Rectangle(pos.Absolute{X: 46, Y: 13}, pos.Absolute{X: 48, Y: 18})

	integrated3 := eraser.Integrated(hall,Crect1, Crect2, Crect3, Crect4, Crect5,
		Crect6,Crect7,Crect8, Crect9, 
		Crect11, Crect12, Crect13, Crect14, Crect15)
	base = model.NewSeatBase(35, 8, "C석")
	blockInput3 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         13,
		YSize:         26,
		EmptyChecker:  integrated3,
		NameFormatter: nameFormatter,
	}

	hallO := eraser.VerticalHallway(16, 33)
	Orect1 := eraser.Rectangle(pos.Absolute{X: 8, Y: 3}, pos.Absolute{X: 12, Y: 3})
	Orect2 := eraser.Rectangle(pos.Absolute{X: 37, Y: 3}, pos.Absolute{X: 41, Y: 3})
	Orect3 := eraser.Rectangle(pos.Absolute{X: 8, Y: 4}, pos.Absolute{X: 9, Y: 4})
	Orect4 := eraser.Rectangle(pos.Absolute{X: 40, Y: 4}, pos.Absolute{X: 41, Y: 4})
	Orect5 := eraser.Rectangle(pos.Absolute{X: 8, Y: 5}, pos.Absolute{X: 8, Y: 5})
	Orect6 := eraser.Rectangle(pos.Absolute{X: 41, Y: 5}, pos.Absolute{X: 41, Y: 5})
	integrated4 := eraser.Integrated(hallO, Orect1, Orect2, Orect3, Orect4, Orect5, Orect6)
	base = model.NewSeatBase(8, 3, "O석")
	blockInput4 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         34,
		YSize:         4,
		EmptyChecker:  integrated4,
		NameFormatter: nameFormatter,
	}

	integrated5 := eraser.Integrated()
	base = model.NewSeatBase(2, 3, "L1석")
	blockInput5 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         9,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(1, 10, "L1석")
	blockInput11 := group.BlockInput{
		Criteria:      model.NewSeat(base, 10, 1),
		XSize:         1,
		YSize:         2,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(2, 13, "L2석")
	blockInput6 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(2, 18, "L3석")
	blockInput7 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(1, 18, "L3석")
	blockInput12 := group.BlockInput{
		Criteria:      model.NewSeat(base, 7, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(47, 3, "R1석")
	blockInput8 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         9,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(48, 10, "R1석")
	blockInput13 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         2,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}

	base = model.NewSeatBase(47, 13, "R2석")
	blockInput9 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         4,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(47, 18, "R3석")
	blockInput10 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}
	base = model.NewSeatBase(48, 18, "R3석")
	blockInput14 := group.BlockInput{
		Criteria:      model.NewSeat(base, 1, 1),
		XSize:         1,
		YSize:         6,
		EmptyChecker:  integrated5,
		NameFormatter: nameFormatter,
	}
	block1 := group.HorizontalBlock(blockInput1)
	block2 := group.HorizontalBlock(blockInput2)
	block3 := group.HorizontalBlock(blockInput3)
	block4 := group.HorizontalBlock(blockInput4)
	block5 := group.VerticalBlock(blockInput5)
	block6 := group.VerticalBlock(blockInput6)
	block7 := group.VerticalBlock(blockInput7)
	block8 := group.VerticalBlock(blockInput8)
	block9 := group.VerticalBlock(blockInput9)
	block10 := group.VerticalBlock(blockInput10)
	block11 := group.VerticalBlock(blockInput11)
	block12 := group.VerticalBlock(blockInput12)
	block13 := group.VerticalBlock(blockInput13)
	block14 := group.VerticalBlock(blockInput14)

	return group.Mixed(block1, block2, block3, block4, block5, block6, block7, block8, block9,block10, block11, block12, block13, block14)
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
		YSize:           31,
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
