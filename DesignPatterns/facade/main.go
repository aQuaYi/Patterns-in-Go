package main

import "fmt"

type facade struct {
	music
	video
	count
}

func (f *facade) getRecommendVideos() error {
	f.getVideos()
	f.getCountByID(111)
	return nil
}

type music struct {
}

func (m *music) getMusic() error {
	fmt.Println("get music material")
	// logic code here
	return nil
}

type video struct {
	vid int64
}

func (v *video) getVideos() error {
	fmt.Println("get videos1")
	return nil
}

type count struct {
	PraiseCnt  int64 //点赞数
	CommentCnt int64 //评论数
	CollectCnt int64 //收藏数
}

func (c *count) getCountByID(id int64) (*count, error) {
	fmt.Println("get video counts")
	return c, nil
}

func main() {
	f := &facade{}
	f.getRecommendVideos()
}
