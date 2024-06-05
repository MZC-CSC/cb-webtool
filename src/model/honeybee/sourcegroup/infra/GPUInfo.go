package honeybee

type Gpu struct {
	Drm    Drm    `json:"drm"`
	Nvidia Nvidia `json:"nvidia"`
}
