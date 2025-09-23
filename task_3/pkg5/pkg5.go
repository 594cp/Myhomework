package pkg5

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // 一定不能忘记导入数据库驱动
	"gorm.io/gorm"
)

func init() {
	fmt.Println("pkg4 init method invoked")
}

type User struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Post        []Post `gorm:"foreignKey:UserID"`
	PostsNumber int    `db:"posts_number"`
}

type Post struct {
	ID             int       `db:"id"`
	Title          string    `db:"title"`
	Text           string    `db:"text"`
	UserID         int       `db:"user_id"`
	CommentsNumber int       `db:"comments_number"`
	Comments       []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	ID      int    `db:"id"`
	Context string `db:"context"`
	PostID  int    `db:"post_id"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("comment before create invoked")
	// 获取当前帖子对应的用户
	var user1 User
	if err := tx.First(&user1, p.UserID).Error; err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	// 更新用户的 ArticleCount 字段
	user1.PostsNumber++
	if err := tx.Save(&user1).Error; err != nil {
		return fmt.Errorf("failed to update user article count: %w", err)
	}

	return nil
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("comment before create invoked")
	// 获取当前帖子对应的用户
	var post1 Post
	if err := tx.First(&post1, c.PostID).Error; err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	// 更新用户的 ArticleCount 字段
	post1.CommentsNumber++
	if err := tx.Save(&post1).Error; err != nil {
		return fmt.Errorf("failed to update user article count: %w", err)
	}

	return nil
}

// func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
// 	fmt.Println("comment before create invoked")
// 	// 获取当前帖子对应的用户
// 	var post1 Post
// 	fmt.Println("comment before delete invoked ----  ", c.ID, c.Context, c.PostID)
// 	if err := tx.Preload("Comments").First(&post1, c.PostID).Error; err != nil {
// 		return fmt.Errorf("failed to find user: %w", err)
// 	}

// 	// 更新用户的 ArticleCount 字段
// 	post1.CommentsNumber--
// 	if err := tx.Save(&post1).Error; err != nil {
// 		return fmt.Errorf("failed to update user article count: %w", err)
// 	}
// 	return nil
// }

func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	// 获取当前评论对应的帖子
	var post Post
	if err := tx.Preload("Comments").First(&post, c.PostID).Error; err != nil {
		return fmt.Errorf("failed to find post: %w", err)
	}
	post.CommentsNumber--
	// 检查帖子的评论数量

	if err := tx.Save(&post).Error; err != nil {
		return fmt.Errorf("failed to update post comment status: %w", err)
	}

	return nil
}

func RunForInitOrder(db *gorm.DB) {
	fmt.Println("pkg5.PkgNameVar has been initialized")
	db.AutoMigrate(&User{}, &Post{}, &Comment{}) // 全局更新
	//输入用户列表
	// user := []User{
	// 	{
	// 		ID:   1,
	// 		Name: "user1",
	// 		Post: []Post{},
	// 	},
	// 	{
	// 		ID:   2,
	// 		Name: "user2",
	// 		Post: []Post{},
	// 	},
	// }
	// db.Create(&user)
	// /////
	// Post1 := Post{
	// 	ID:       1,
	// 	Title:    "post1",
	// 	Text:     "text1",
	// 	UserID:   1,
	// 	Comments: []Comment{},
	// }
	// db.Create(&Post1)
	// Post2 := Post{
	// 	ID:       2,
	// 	Title:    "post2",
	// 	Text:     "text2",
	// 	UserID:   2,
	// 	Comments: []Comment{},
	// }
	// db.Create(&Post2)
	// Post3 := Post{
	// 	ID:       3,
	// 	Title:    "post3",
	// 	Text:     "text3",
	// 	UserID:   1,
	// 	Comments: []Comment{},
	// }
	// db.Create(&Post3)

	// Comment1 := Comment{
	// 	ID:      1,
	// 	Context: "comment1",
	// 	PostID:  1,
	// }
	// db.Create(&Comment1)
	// Comment2 := Comment{
	// 	ID:      2,
	// 	Context: "comment2",
	// 	PostID:  1,
	// }
	// db.Create(&Comment2)
	// Comment3 := Comment{
	// 	ID:      3,
	// 	Context: "comment3",
	// 	PostID:  2,
	// }
	// db.Create(&Comment3)
	// Comment4 := Comment{
	// 	ID:      4,
	// 	Context: "comment4",
	// 	PostID:  2,
	// }
	// db.Create(&Comment4)
	// Comment5 := Comment{
	// 	ID:      5,
	// 	Context: "comment5",
	// 	PostID:  3,
	// }
	// db.Create(&Comment5)

	// Post3 := Post{
	// 	ID:       4,
	// 	Title:    "post3",
	// 	Text:     "text3",
	// 	UserID:   2,
	// 	Comments: []Comment{},
	// }
	// db.Create(&Post3)

	// Comment6 := Comment{
	// 	ID:      7,
	// 	Context: "comment6",
	// 	PostID:  5,
	// }
	// err := db.Create(&Comment6).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//查看用户2的文章列表
	users := []User{}
	err := db.Preload("Post.Comments").Where("ID = ?", 2).Find(&users).Error
	if err != nil {
		fmt.Println("1d__", err)
	} else {
		for _, user := range users {
			fmt.Println(user.ID, user.Name)
			for _, post := range user.Post {
				fmt.Println("___", post.ID, post.Title)
				for _, comment := range post.Comments {
					fmt.Println("_______", comment.ID, comment.Context)
				}
			}
		}
	}

	//posts1 := []Post{}
	//db.Model(&Post{}).Preload("Comments").Find(&posts1)
	//.Model(&Post{}).Select("COUNT() as COUNTS").Preload("Comments").Where("id = ?", 1).Find(&posts1)
	// for _, post := range posts1 {
	// 	fmt.Println(post.ID, post.Title, post.Comments[0].Context)
	// }
	//fmt.Println(users[0].ID, users[0].Name, users[0].Post[0].Title, users[0].Post[0].Comments[0].Context)

	// Comment6 := Comment{
	// 	ID:      7,
	// 	Context: "comment6",
	// 	PostID:  2,
	// }
	// db.Create(&Comment6)
	// fmt.Println("pkg5.RunForInitOrder has been invoked")

	// Post3 := Post{
	// 	//ID:       4,
	// 	Title:    "post3333",
	// 	Text:     "text3333",
	// 	UserID:   2,
	// 	Comments: []Comment{},
	// }

	//db.Delete(&User{}, 10)
	Pos := []Post{}
	err = db.Find(&Pos).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Pos)
	}

	Com := []Comment{}
	err = db.Find(&Com).Error
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(Com)
	}
	db.Delete(&Comment{}, "3")
}

// fmt.Println(com)
// var users []User
// err = db.Model(&User{}).Preload("Posts").Find(&users).Error
