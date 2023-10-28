package render

func getItemSprite(item int) string {
	switch item {
	case 69:
		return "$"
	default:
		return " "
	}
}
