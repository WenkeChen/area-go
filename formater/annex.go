package formater

import "AreaGo/model"

type AnnexItem struct {
	ID   uint   `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Src  string `json:"src"`
}

//文章meta
func BuildPostAnnexItem(annex model.Annex) AnnexItem {
	return AnnexItem{
		ID:   annex.ID,
		Type: annex.Type,
		Name: annex.Name,
		Src:  annex.Src,
	}
}

//文章所有的meta
func BuildPostAnnexList(annexes []model.Annex) map[string][]AnnexItem {
	annexList := make(map[string][]AnnexItem)
	for _, t := range annexes {
		annexList[t.Type] = append(annexList[t.Type], BuildPostAnnexItem(t))
	}
	return annexList
}
