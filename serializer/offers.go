package serializer

import "server/model"

// Offer 招聘序列化器
type Offer struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

// BuildOffers 序列化招聘
func BuildOffer(item model.Offers) Offer {
	return Offer{
		ID:        item.ID,
		Title:     item.Title,
		Content:   item.Content,
		CreatedAt: item.CreatedAt.Unix(),
	}
}

// BuildOffers 序列化招聘列表
func BuildOffers(items []model.Offers) []Offer {
	var offers []Offer
	for _, item := range items {
		offerV := BuildOffer(item)
		offers = append(offers, offerV)
	}
	return offers
}
