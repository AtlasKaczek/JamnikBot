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
	Subreddit             string `json:"subreddit_name_prefixed"`
	Title                 string `json:"title"`
	PostHint              string `json:"post_hint"`
	Url_overriden_by_dest string `json:"url_overridden_by_dest"`
}

//Getters

func (t Images) GetImageURL(i int) string {
	return t.Data.Children[i].DataV2.Url_overriden_by_dest
}

func (t Images) GetChildrenLen() int {
	return len(t.Data.Children)
}

func (t Images) GetImagesIndexList() []int {
	var i_list []int

	for i := 0; i < t.GetChildrenLen(); i++ {
		if t.Data.Children[i].DataV2.PostHint == "image" {
			i_list = append(i_list, i)
		}
	}
	return i_list
}

// func (t Images) GetMediaId(i,j int) string {
// 	return t.Data.Children[i].DataV2.Gallery_data.Items[j].Media_id
// }
