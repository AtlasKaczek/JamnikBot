package aplikacja

//stracture
type Images struct {
	Data Data `json:"data"`
}

type Data struct {
	Children []Children `json:"children"`
}

type Children struct {
	DataV2 DataV2 `json:"data"`
}

type DataV2 struct {
	Subreddit string `json:"subreddit_name_prefixed"`
	Title     string `json:"title"`
	// Media_metadata Media_metadata `json:"media_metadata"`
	// Gallery_data   Gallery_data   `json:"gallery_data"`
	Url_overriden_by_dest string `json:"url_overridden_by_dest"`
}

// type Media_metadata struct {
// 	ID ID `json:""`
// }

// type Gallery_data struct {
// 	Items []Item `json:"items"`
// }

// type Item struct {
// 	Media_id string `json:"media_id"`
// }

//Getters

func (t Images) GetImageURL(i int) string {
	return t.Data.Children[i].DataV2.Url_overriden_by_dest
}

func (t Images) GetChildrenLen() int {
	return len(t.Data.Children)
}

// func (t Images) GetMediaId(i,j int) string {
// 	return t.Data.Children[i].DataV2.Gallery_data.Items[j].Media_id
// }
