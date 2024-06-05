package honeybee

type Nvidia []struct {
	DeviceAttribute DeviceAttribute `json:"device_attribute"`
	Performance     Performance     `json:"performance"`
}
