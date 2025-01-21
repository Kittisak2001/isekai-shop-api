package main

import (
	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database).ConnectionGetting()
	tx := db.Begin()
	defer tx.Rollback()

	itemAdding(tx)

	if err := tx.Commit().Error; err != nil {
		panic(err)
	}
}

func itemAdding(tx *gorm.DB) {
	items := []entities.Item{
		{
			Name:        "Sword",
			Description: "A sword that can be used to fight enemies.",
			Price:       100,
			Picture:     "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSGTvapLJNt7l9Dp-pAsLOoFy8A-Dw2XfBI8w&s",
		},
		{
			Name:        "Shield",
			Description: "A shield that can be used to block enemy attacks.",
			Price:       50,
			Picture:     "https://i.pinimg.com/736x/9a/6c/12/9a6c121be8cc5b037526f9922b956db1.jpg",
		},
		{
			Name:        "Potion",
			Description: "A potion that can be used to heal wounds.",
			Price:       30,
			Picture:     "https://images.rawpixel.com/image_png_social_square/cHJpdmF0ZS9sci9pbWFnZXMvd2Vic2l0ZS8yMDIzLTAxL2ZycG90aW9uX2JvdHRsZV9tYWdpY19wb3Rpb24taW1hZ2Utam9iMTUyNy5wbmc.png",
		},
		{
			Name:        "Bow",
			Description: "A bow that can be used to shoot enemies from afar.",
			Price:       80,
			Picture:     "https://i.pinimg.com/736x/48/c7/02/48c70227e94581e880ea1a063b669ba6.jpg",
		},
		{
			Name:        "Arrow",
			Description: "An arrow that can be used with a bow to shoot enemies from afar.",
			Price:       10,
			Picture:     "https://e7.pngegg.com/pngimages/805/184/png-clipart-the-elder-scrolls-v-skyrim-bow-and-arrow-weapon-archery-arrow-bow-game-branch.png",
		},
	}

	tx.CreateInBatches(items, len(items))
}
