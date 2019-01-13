package gcutil

type GCIndex struct {
	Title       string     `json:"title"`
	Description string     `json:"desc"`
	Canonical   string     `json:"canonical"`
	GoToUrl     string     `json:"go_to_url"`
	Paging      Pagination `json:"paging"`
	Courses     []Course   `json:"courses"`
}

type GCSingle struct {
	Title       string   `json:"title"`
	Description string   `json:"desc"`
	Canonical   string   `json:"canonical"`
	GoToUrl     string   `json:"go_to_url"`
	Course      Course   `json:"courses"`
	Related     []Course `json:"related"`
}

type GCGoTo struct {
	Course Course `json:"courses"`
}

type Courses struct {
	Course Course `json:"course"`
}

type Pagination struct {
	TotalPages int64 `json:"total_pages"`
	PpbtPage   int64 `json:"ppbt_page"`
	PbtPage    int64 `json:"pbt_page"`
	ActivePage int   `json:"active_page"`
	PrevPage   int   `json:"prev_page"`
	NextPage   int   `json:"next_page"`
}

type Course struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Description   string   `json:"description"`
	Category      string   `json:"category"`
	CatSlug       string   `json:"cat_slug"`
	RemoteUrl     string   `json:"remote_url"`
	Slug          string   `json:"slug"`
	PostIMG       string   `json:"post_img"`
	FrontIMG      string   `json:"front_img"`
	Rating        string   `json:"rating"`
	Enrolled      string   `json:"enrolled"`
	Language      string   `json:"language"`
	CurrentPrice  string   `json:"current_price"`
	OriginalPrice string   `json:"original_price"`
	DiscountRate  string   `json:"discount_rate"`
	CouponCode    string   `json:"coupon_code"`
	ValidCoupon   bool     `json:"valid_coupon"`
	Timestamp     int64    `json:"timestamp"`
	PrettyURL     string   `json:"pretty_url"`
	WillLearn     []string `json:"will_learn"`
}
